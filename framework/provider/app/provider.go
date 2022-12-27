package app

import (
	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

// HtdAppProvider 提供App的具体实现方法
type HtdAppProvider struct {
	BaseFolder string
}

// Register 注册HtdApp方法
func (h *HtdAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHtdApp
}

// Boot 启动调用
func (h *HtdAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *HtdAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *HtdAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *HtdAppProvider) Name() string {
	return contract.AppKey
}
