package app

import (
	"github.com/codegangsta/cli"
	"common"
	"core"
)

func initAppFlags(){

	//debug开关
	common.AddFlagBool(cli.BoolFlag{
		Name:  "debug",
		Usage: "if open the debug mode",
	})

	//设置并发数
	common.AddFlagInt(cli.IntFlag{
		Name:  "c",
		Value: 10,
		Usage: "number of concurrent connections to use",
	})

	//测试持续时间
	common.AddFlagInt(cli.IntFlag{
		Name:  "d",
		Value: 10,
		Usage: "duration of test in seconds",
	})

	//调用http超时时间
	common.AddFlagInt(cli.IntFlag{
		Name:  "t",
		Value: 1000,
		Usage: "socket/request timeout in (ms)",
	})

	//http 方法 GET/POST
	common.AddFlagString(cli.StringFlag{
		Name:  "m",
		Value: "GET",
		Usage: "http method, use 'GET' or 'POST' ",
	})

	//设置header
	common.AddFlagString(cli.StringFlag{
		Name:  "H",
		Value: "",
		Usage: "the http headers sent to the target url, e.g:'keyA:ValueA;KeyB:ValueB' ",
	})

	//是否开启 keep-alived
	common.AddFlagBool(cli.BoolFlag{
		Name:  "keepalive",
		Usage: "if keep-alives",
	})

	//是否压缩
	common.AddFlagBool(cli.BoolFlag{
		Name:  "compress",
		Usage: "if prevents sending the \"Accept-Encoding: gzip\" header",
	})

	//POST数据文件
	common.AddFlagString(cli.StringFlag{
		Name:  "dataFile",
		Value: "",
		Usage: "load the par which store in the file",
	})
}

func createPayload(c * cli.Context, strUrl string) error{
	conn_num := c.Int("c")
	duration := c.Int("d")
	timeout := c.Int("t")
	method := c.String("m")
	header := c.String("H")
	keepalive := c.Bool("keepalive")
	compress := c.Bool("compress")
	dataFile := c.String("dataFile")
	debug := c.Bool("debug")

	res, err := core.CreatePayload(strUrl, conn_num, duration, timeout, method, header, keepalive, compress, dataFile)
	if err != nil {
		return err
	}

	if !debug {
		common.FKLogPrintf(res)
	}
	return nil
}