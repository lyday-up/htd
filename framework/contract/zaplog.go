package contract

// import "github.com/htd/framework/zap/zapcore"

const LogZapKey = "htd:zap"

// Log 定义了日志服务协议
type ZapLog interface {
	// // Panic 表示会导致整个程序出现崩溃的日志信息
	// Panic(...any)
	// // Fatal 表示会导致当前这个请求出现提前终止的错误信息
	// Fatal(...any)
	// // Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	// Error(...any)
	// // Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
	// Warn(...any)
	// // Info 表示正常的日志信息输出
	// Info(...any)
	// // Debug 表示在调试状态下打印出来的日志信息
	// Debug(...any)
	// // Trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
	// Info(msg string, fields ...zapcore.Field)
	Info(args ...any)
}
