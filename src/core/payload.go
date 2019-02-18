package core

import (
	"net/url"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"common"
	"io/ioutil"
	"time"
)

func CreatePayload(strUrl string,
	connNum int,
	duration int,
	timeout int,
	method string,
	header string,
	keepalive bool,
	compress bool,
	dataFile string) (string, error) {

	// 检查URL的合法性
	pUrl, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	} else if pUrl.Host == "" {
		return "", errors.New("查询URL不应为空")
	}

	statsChann := make(chan *Stats, connNum)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT /*, syscall.SIGSTOP*/)

	common.FKLogPrintf("将使用 %v 个协程同时调用 %v 持续 %v 秒.\n", connNum, strUrl, duration)

	// 读取Body
	var data []byte
	if dataFile != "" && common.IsDataFileExist(dataFile) == nil {
		f, _ := os.Open(dataFile)
		defer f.Close()
		data, _ = ioutil.ReadAll(f)
	}

	// 创建WORKER
	worker := NewWorker(strUrl, connNum, duration, timeout, header, method, statsChann, keepalive, compress, data )
	for i := 0; i < connNum; i++{
		go worker.Run()
	}

	respondIndexs := 0
	aggStats := Stats{MinRequestTime:time.Minute}
	for respondIndexs < connNum{
		select {
		case <-sigChan:
			worker.Stop()
			common.FKLogPrintln("协程压测结束.")
		case stats := <-statsChann:
			aggStats.NumErrs += stats.NumErrs
			aggStats.NumRequests += stats.NumRequests
			aggStats.RespSize += stats.RespSize
			aggStats.Duration += stats.Duration
			aggStats.MaxRequestTime = common.MaxDuration(stats.MaxRequestTime, aggStats.MaxRequestTime)
			aggStats.MinRequestTime = common.MinDuration(stats.MinRequestTime, aggStats.MinRequestTime)
			aggStats.Num2X += stats.Num2X
			aggStats.Num5X += stats.Num5X
			respondIndexs++
		}
	}

	common.FKLogPrintf("压测完成. 收到 %v 个协程测试结果.\n", respondIndexs)
	if aggStats.NumRequests <= 0{
		return "", errors.New("压测未收到任何有效响应.")
	}

	return aggStats.String(respondIndexs), nil
}
