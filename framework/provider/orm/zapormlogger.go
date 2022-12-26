package orm

import (
	"context"
	"gorm.io/gorm/logger"
	"time"

	"github.com/htd/framework/contract"
	// "github.com/htd/framework/provider/zaplog"
)

// OrmLogger orm的日志实现类, 实现了gorm.Logger.Interface
type OrmZapLogger struct {
	logger contract.ZapLoger // zaplog.HtdZapLogService // 有一个logger对象存放htd的log服务
}

// NewOrmLogger 初始化一个ormLogger,
func NewOrmZapLogger(logger contract.ZapLoger) *OrmZapLogger {
	return &OrmZapLogger{logger: logger}
}

// Trace 对接htd的Trace输出
func (z *OrmZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	// s := fmt.Sprintf("orm trace sql: %v ,rows: %v, spent time: %v, err: %v", sql,rows,elapsed, err)
	s := "orm trace sql: %v ,rows: %v, spent time: %v, err: %v"
	z.logger.Tracef(s, sql, rows, elapsed, err)
}

func (z *OrmZapLogger) Panic(ctx context.Context, str string, args ...any) {
	z.logger.Panicf(str, args...)
}

func (z *OrmZapLogger) Fatal(ctx context.Context, str string, args ...any) {
	z.logger.Fatalf(str, args...)

}
func (z *OrmZapLogger) Debug(ctx context.Context, str string, args ...any) {
	z.logger.Debugf(str, args...)
}
func (z *OrmZapLogger) Info(ctx context.Context, str string, args ...any) {
	z.logger.Infof(str, args...)
}
func (z *OrmZapLogger) Warn(ctx context.Context, str string, args ...any) {
	z.logger.Warnf(str, args...)
}
func (z *OrmZapLogger) Error(ctx context.Context, str string, args ...any) {
	z.logger.Errorf(str, args...)
}

func (z *OrmZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	return z
}
