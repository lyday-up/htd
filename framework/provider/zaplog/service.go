package zaplog

import (
	"os"
	"path/filepath"

	"github.com/htd/framework"
	"github.com/htd/framework/contract"
	"github.com/htd/framework/util"
	"github.com/htd/framework/zap"
	"github.com/htd/framework/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type HtdZapLogService struct {
	*zap.Logger
}

func NewHtdZapLog(params ...any) (any, error) {

	c := params[0].(framework.Container)

	appService := c.MustMake(contract.AppKey).(contract.App)
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	// 从配置文件中获取folder信息，否则使用默认的LogFolder文件夹
	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	// 如果folder不存在，则创建
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	// 从配置文件中获取file信息，否则使用默认的hade.log
	file := "htd.log"
	if configService.IsExist("log.file") {
		file = configService.GetString("log.file")
	}

	writeSyncer := getLogWriter(filepath.Join(folder, file))
	encoder := getTextEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	return sugarLogger, nil
}

func getLogWriter(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,  // Filename: 日志文件的位置
		MaxSize:    10,    // MaxSize：在进行切割之前，日志文件的最大大小（以 MB 为单位）
		MaxBackups: 5,     // MaxBackups：保留旧文件的最大个数
		MaxAge:     30,    // MaxAges：保留旧文件的最大天数
		Compress:   false, // Compress：是否压缩 / 归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getTextEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
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
