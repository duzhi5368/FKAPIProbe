package core

import (
	"sync/atomic"
	"time"
	"net/http"
	"net/url"
	"crypto/tls"
	"strings"
	"common"
	"bytes"
	"const"
	"io/ioutil"
)

type Worker struct{
	Url 		string
	Duration 	int
	ConnNum 	int
	Timeout 	int
	Method 		string
	Header 		string
	IsKeepAlive bool
	IsCompress 	bool
	StatsChan  	chan *Stats
	BodyData	[]byte
	Interrupted int32
}

func NewWorker(strUrl string, connNum int, duration int, timeout int,
	header string, method string, statsChan chan *Stats, keepalive bool,
	compress bool, bodyData []byte	)(worker *Worker){
	worker = &Worker{strUrl, duration, connNum, timeout, method,
	header, keepalive,  compress, statsChan, bodyData, 0}
	return
}

func (w *Worker) Run(){
	stats := &Stats{MinRequestTime: time.Minute}

	start := time.Now()
	httpClient := &http.Client{}

	pUrl, _ := url.Parse(w.Url)
	var tlsConfig *tls.Config
	if pUrl.Scheme == "https"{
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}else{
		tlsConfig = nil
	}

	httpClient.Transport = &http.Transport{
		DisableKeepAlives: !w.IsKeepAlive,
		DisableCompression: !w.IsCompress,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(w.Timeout),
		TLSClientConfig: tlsConfig,
	}
	httpClient.Timeout = time.Second * time.Duration(w.Duration)

	// HEADER
	sets := strings.Split(w.Header, ";")
	headerMap := make(map[string]string)
	for i := range sets{
		split := strings.SplitN(sets[i], ":", 2)
		if len(split) == 2{
			headerMap[split[0]] = split[1]
		}
	}

	// 持续间隔
	for time.Since(start).Seconds() <= float64(w.Duration) && atomic.LoadInt32(&w.Interrupted) == 0{
		responseSize, num2x, num5x, reqDuration := DoRequest(httpClient, headerMap, w.Method, w.Url, w.BodyData)
		if responseSize > 0{
			stats.RespSize += int64(responseSize)
			stats.Duration += reqDuration
			stats.MaxRequestTime = common.MaxDuration(reqDuration, stats.MaxRequestTime)
			stats.MinRequestTime = common.MinDuration(reqDuration, stats.MinRequestTime)
			stats.NumRequests++
			stats.Num2X +=num2x
		}else{
			stats.Num5X += num5x
			stats.NumErrs++
		}
	}
	w.StatsChan <- stats
}

func (w *Worker) Stop(){
	atomic.StoreInt32(&w.Interrupted, 1)
}

// http请求
func DoRequest(httpClient *http.Client, header map[string]string, method string, strUrl string, bodyData []byte)(responseSize int, num2x int, num5x int, duration time.Duration){
	responseSize = -1
	duration = -1
	num2x = 0
	num5x = 0
	strUrl = common.EscapeUrlStr(strUrl)

	req, err := http.NewRequest(method, strUrl, bytes.NewBuffer(bodyData))
	if err != nil{
		common.FKLogPrintln("Create http request error:", err)
		return
	}
	req.Header.Add("User-Agent", _const.USER_AGENT)
	for k, v := range header{
		req.Header.Set(k, v)
	}

	start := time.Now()
	response, err := httpClient.Do(req)
	if err != nil{
		rr, ok := err.(*url.Error)
		if !ok{
			common.FKLogPrintln("Http request error:", err, rr)
			return
		}else{
			common.FKLogPrintln("Http request error:", err)
			return
		}
	}
	if response == nil{
		common.FKLogPrintln("Empty http response.")
		return
	}
	defer func(){
		if response != nil && response.Body != nil{
			response.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		common.FKLogPrintln("Read http response error:", err)
	}

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusMethodNotAllowed{
		duration = time.Since(start)
		responseSize = len(body) + int(common.EstimateHttpHeadersSize(response.Header))
		num2x += 1
	} else if response.StatusCode == http.StatusMovedPermanently || response.StatusCode == http.StatusTemporaryRedirect {
		duration = time.Since(start)
		responseSize = int(response.ContentLength) + int(common.EstimateHttpHeadersSize(response.Header))
	} else if response.StatusCode >= 500 {
		num5x += 1
	} else {
		common.FKLogPrintln("received unknown status code", response.StatusCode, "from", response.Header, "content", string(body), req)
	}

	return
}