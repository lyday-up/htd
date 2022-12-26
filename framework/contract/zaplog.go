package contract

// import "github.com/htd/framework/zap/zapcore"

const ZapLogKey = "htd:zap"

// Log 定义了日志服务协议
type ZapLoger interface {
	// Panic 表示会导致整个程序出现崩溃的日志信息
	Panicf(string, ...interface{})
	// Fatal 表示会导致当前这个请求出现提前终止的错误信息
	Fatalf(string, ...interface{})
	// Debug 表示在调试状态下打印出来的日志信息
	Debugf(string, ...interface{})
	// Info 表示正常的日志信息输出
	Infof(string, ...interface{})
	// Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
	Warnf(string, ...interface{})
	// Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	Errorf(string, ...interface{})

	Tracef(string, ...interface{})

	Panic(...interface{})
	// Fatal 表示会导致当前这个请求出现提前终止的错误信息
	Fatal(...interface{})
	// Debug 表示在调试状态下打印出来的日志信息
	Debug(...interface{})
	// Info 表示正常的日志信息输出
	Info(...interface{})
	// Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
	Warn(...interface{})
	// Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	Error(...interface{})

	Trace(...interface{})

	SSync()
}
