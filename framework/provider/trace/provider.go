package trace

import (
	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

type HtdTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *HtdTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewHtdTraceService
}

// Boot will called when the service instantiate
func (provider *HtdTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HtdTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HtdTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

// / Name define the name for this service
func (provider *HtdTraceProvider) Name() string {
	return contract.TraceKey
}
