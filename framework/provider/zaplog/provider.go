package zaplog

import (
	"fmt"

	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

// HtdZapServiceProvider 服务提供者
type HtdZapServiceProvider struct {
	Encoding   string // json | text
	Level      string // 日志级别
	LogPath    string // 日志路径
	LogNmae    string // LogName: 日志文件的名称
	MaxSize    int    // MaxSize：在进行切割之前，日志文件的最大大小（以 MB 为单位）
	MaxBackups int    // MaxBackups：保留旧文件的最大个数
	MaxAge     int    // MaxAges：保留旧文件的最大天数
	Compress   bool
}

// Register 注册一个服务实例
func (l *HtdZapServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewHtdZapLog
}

// Boot 启动的时候注入
func (l *HtdZapServiceProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (l *HtdZapServiceProvider) IsDefer() bool {
	return false
}

// Params 定义要传递给实例化方法的参数
func (l *HtdZapServiceProvider) Params(c framework.Container) []interface{} {
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	fmt.Println("conf", configService.Get("app.path"))
	level := configService.Get("log.level")
	fmt.Println("level", level)
	encoder := configService.Get("log.encoder")
	fmt.Println("coder", encoder)

	logFile := configService.Get("log.log_name")
	fmt.Println("level", logFile)

	maxSize := configService.GetInt("log.max_size")
	maxBackups := configService.GetInt("log.max_backups")
	maxAge := configService.GetInt("log.max_age")
	compress := configService.GetBool("log.compress")
	logPath := configService.GetString("log.log_path")
	return []interface{}{c, encoder, level, logFile, maxSize, maxBackups, maxAge, compress, logPath}
}

// Name 定义对应的服务字符串凭证
func (l *HtdZapServiceProvider) Name() string {
	return contract.ZapLogKey
}
