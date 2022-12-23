package kernel

import (
	"github.com/htd/framework/gin"
	"net/http"
)

type HtdKernelService struct {
	engine *gin.Engine
}

func NewHtdKernelService(params ...any) (any, error) {
	httpEngine := params[0].(*gin.Engine)
	return &HtdKernelService{engine: httpEngine}, nil
}

func (s *HtdKernelService) HttpEngine() http.Handler {
	return s.engine
}
