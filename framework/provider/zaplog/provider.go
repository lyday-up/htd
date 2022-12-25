package zaplog

import (
	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

// HtdZapServiceProvider 服务提供者
type HtdZapServiceProvider struct {
	Encoding string // json | text
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

	return []interface{}{c}
}

// Name 定义对应的服务字符串凭证
func (l *HtdZapServiceProvider) Name() string {
	return contract.LogZapKey
}
