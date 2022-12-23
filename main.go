package main

import (
	"github.com/htd/app/console"
	"github.com/htd/app/http"
	"github.com/htd/framework"
	"github.com/htd/framework/provider/app"
	"github.com/htd/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHtdContainer()
	// 绑定App服务提供者
	container.Bind(&app.HtdAppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.HtdKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)

}
