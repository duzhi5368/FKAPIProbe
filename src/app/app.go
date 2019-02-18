package app

import (
	"github.com/codegangsta/cli"
	"const"
	"common"
	"os"
)

//应用启动Action
func fkAppAction(c *cli.Context) (err error) {
	strUrl := os.Args[len(os.Args)-1]
	return createPayload(c, strUrl)
}

// 程序入口
func StartUp(){
	// 初始化APP Flag
	initAppFlags()
	// 创建控制台App
	cliApp := cli.NewApp()
	cliApp.Name = "FKAPIProbe"
	cliApp.Usage = "这是FreeKnight自用的简易Web API压测和性能分析工具"
	cliApp.Version = _const.APP_VERSION
	cliApp.Flags = common.GetAppFlags()
	cliApp.Action = common.ActionWrapper(fkAppAction)
	cliApp.Run(os.Args)
}
