package common

import (
	"errors"
	"os"
	"time"
	"strings"
	"net/url"
	"net/http"
)

//检查指定路径的文件是否存在
func IsDataFileExist(filePath string) error {
	if filePath == "" {
		return errors.New("数据文件路径为空。")
	}

	if _, err := os.Stat(filePath); err != nil {
		return errors.New("打开数据文件失败: " + filePath + " 错误: " + err.Error())
	}
	return nil
}

func MaxDuration(d1 time.Duration, d2 time.Duration) time.Duration {
	if d1 > d2 {
		return d1
	} else {
		return d2
	}
}

func MinDuration(d1 time.Duration, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d1
	} else {
		return d2
	}
}

// 计算HTTP头字节长度
func EstimateHttpHeadersSize(headers http.Header) (result int64) {
	result = 0

	for k, v := range headers {
		result += int64(len(k) + len(": \r\n"))
		for _, s := range v {
			result += int64(len(s))
		}
	}

	result += int64(len("\r\n"))

	return result
}

// 格式化Duration输出
func FmtDuration(d time.Duration) string {
	return d.String()
	//ms := float64(d) / float64(time.Millisecond)
	//return fmt.Sprintf("%.02f ms",  ms)
}

func EscapeUrlStr(in string) string {
	qm := strings.Index(in, "?")
	if qm != -1 {
		qry := in[qm+1:]
		qrys := strings.Split(qry, "&")
		var query string = ""
		var qEscaped string = ""
		var first bool = true
		for _, q := range qrys {
			qSplit := strings.Split(q, "=")
			if len(qSplit) == 2 {
				qEscaped = qSplit[0] + "=" + url.QueryEscape(qSplit[1])
			} else {
				qEscaped = qSplit[0]
			}
			if first {
				first = false
			} else {
				query += "&"
			}
			query += qEscaped

		}
		return in[:qm] + "?" + query
	} else {
		return in
	}
}