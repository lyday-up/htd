package services

import (
	"os"

	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

// HtdConsoleLog 代表控制台输出
type HtdConsoleLog struct {
	HtdLog
}

// NewHtdConsoleLog 实例化HtdConsoleLog
func NewHtdConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &HtdConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
