package http

import (
	"github.com/htd/app/http/module/demo"
	"github.com/htd/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
