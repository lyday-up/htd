package zaplog

import (
	"os"
	"path/filepath"

	"github.com/htd/framework/util"
	"github.com/htd/framework/zap"
	"github.com/htd/framework/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type HtdZapLogService struct {
	*zap.Logger
}

func NewHtdZapLog(params ...any) (any, error) {
	// c := params[0].(framework.Container)
	logEncoder := params[1].(string)
	logLevel := params[2].(string)
	logFile := params[3].(string)
	maxSize := params[4].(int)
	maxBackups := params[5].(int)
	maxAge := params[6].(int)
	compress := params[7].(bool)
	logPath := params[8].(string)

	// 如果folder不存在，则创建
	if !util.Exists(logPath) {
		os.MkdirAll(logPath, os.ModePerm)
	}

	fileName := (filepath.Join(logPath, logFile))

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	if logEncoder == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)
	var core zapcore.Core

	switch logLevel {
	case "debug":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	case "info":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	case "warn":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.WarnLevel)
	case "error":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)
	case "fatal":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.FatalLevel)
	case "panic":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.PanicLevel)
	default:
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	}

	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	return sugarLogger, nil
}

// func (z *HtdZapLogService) Panic(str string, args ...any) {
// 	z.Sugar().Panicf(str, args...)
// }

// func (z *HtdZapLogService) Fatal(str string, args ...any) {
// 	z.Sugar().Fatalf(str, args...)

// }
// func (z *HtdZapLogService) Debug(str string, args ...any) {
// 	z.Sugar().Debugf(str, args...)
// }
// func (z *HtdZapLogService) Info(str string, args ...any) {
// 	z.Sugar().Infof(str, args...)
// }
// func (z *HtdZapLogService) Warn(str string, args ...any) {
// 	z.Sugar().Warnf(str, args...)
// }
// func (z *HtdZapLogService) Error(str string, args ...any) {
// 	z.Sugar().Errorf(str, args...)
// }
// func (z *HtdZapLogService) Trace(str string, args ...any) {
// 	z.Sugar().Tracef(str, args...)
// }

// func (z *HtdZapLogService) SetModle(contract.LogLevel) contract.ZapLoger{

// 	return
// }
