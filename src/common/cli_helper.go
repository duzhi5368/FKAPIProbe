package common

import (
	"github.com/codegangsta/cli"
	"fmt"
)


var appFlags = map[string]cli.Flag{}

func ActionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := action(context); err != nil {
			FKLogPrintln(err.Error())
		}
	}
}

func GetAppFlags() (afs []cli.Flag) {
	for _, f := range appFlags {
		afs = append(afs, f)
	}
	return
}

func AddFlagString(sf cli.StringFlag) cli.StringFlag {
	if _, ok := appFlags[sf.Name]; ok {
		FKPanic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}

func AddFlagBool(sf cli.BoolFlag) cli.BoolFlag {
	if _, ok := appFlags[sf.Name]; ok {
		FKPanic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}

func AddFlagInt(sf cli.IntFlag) cli.IntFlag {
	if _, ok := appFlags[sf.Name]; ok {
		FKPanic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}
