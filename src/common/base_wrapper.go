package common

import (
	"log"
	"fmt"
)

func FKLogPrintln(v ...interface{}){
	log.Println(v)
}

func FKLogPrintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func FKPanic(v interface{}){
	panic(v)
}