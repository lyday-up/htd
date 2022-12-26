package test

import (
	"github.com/htd/framework"
	"github.com/htd/framework/provider/app"
)

const (
	BasePath = "/Users/ly/htd"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewHtdContainer()
	// 绑定App服务提供者
	container.Bind(&app.HtdAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	return container
}
