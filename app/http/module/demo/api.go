package demo

import (
	"fmt"

	demoService "github.com/htd/app/provider/demo"
	"github.com/htd/framework/gin"
)

func Register(r *gin.Engine) error {

	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo2", Demo)
	return nil
}

func Demo(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	s := demoProvider.GetHello()
	c.ISetOkStatus().IJson(s)
}
