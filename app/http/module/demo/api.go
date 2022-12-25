package demo

import (
	demoService "github.com/htd/app/provider/demo"
	"github.com/htd/framework/contract"
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
	logger := c.MustMakeLog()
	logger.Info(c, "demo success", nil)

	zaplog := c.MustMakeZapLog()
	zaplog.Info("zap log ok")

	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	c.JSON(200, password)
	c.ISetOkStatus().IJson(s)
}
