package core

import (
	"time"
	"fmt"
	"common"
)

type Stats struct {
	Url            string
	RespSize       int64
	Duration       time.Duration
	MinRequestTime time.Duration
	MaxRequestTime time.Duration
	NumRequests    int
	NumErrs        int
	Num5X          int
	Num2X          int
}


// 输出统计信息
func (aggStats *Stats) String(responders int) (res string) {
	avgThreadDur := aggStats.Duration / time.Duration(responders)
	reqRate := float64(aggStats.NumRequests) / avgThreadDur.Seconds()
	qpsRate := float64(aggStats.NumRequests) / (float64(responders) * avgThreadDur.Seconds())
	avgReqTime := aggStats.Duration / time.Duration(aggStats.NumRequests)
	bytesRate := float64(aggStats.RespSize) / avgThreadDur.Seconds()
	successedRate := float64(aggStats.NumRequests) / float64(aggStats.Num2X) * 100.0
	res += "--------------------------------------\n"
	res += fmt.Sprintf("%v 秒内共发送 %v 个请求\n", avgThreadDur, aggStats.NumRequests)
	res += fmt.Sprintf("每秒平均请求次数:\t%.2f 次\n", reqRate)
	res += fmt.Sprintf("单协程每秒平均请求次数:\t%.2f 次\n", qpsRate)
	res += fmt.Sprintf("收到总字节数 %v\n", common.ByteSize{float64(aggStats.RespSize)})
	res += fmt.Sprintf("平均每秒收到数据:\t%v\n", common.ByteSize{bytesRate})
	res += fmt.Sprintf("平均响应时间:\t%s\n", common.FmtDuration(avgReqTime))
	res += fmt.Sprintf("最快响应时间:\t%s\n", common.FmtDuration(aggStats.MinRequestTime))
	res += fmt.Sprintf("最慢响应时间:\t%s\n", common.FmtDuration(aggStats.MaxRequestTime))
	res += fmt.Sprintf("请求异常数量:\t%v 个\n", aggStats.NumErrs)
	res += fmt.Sprintf("成功请求数量:\t%v 个，占比%.1f%%\n", aggStats.Num2X, successedRate)
	res += fmt.Sprintf("请求错误数量:\t%v 个（服务器内部错误）\n", aggStats.Num5X)
	res += "--------------------------------------\n"
	return
}