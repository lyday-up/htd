package env

import (
	"github.com/htd/framework"
	"github.com/htd/framework/contract"
)

type HtdEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *HtdEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewHtdEnv
}

// Boot will called when the service instantiate
func (provider *HtdEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HtdEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HtdEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// / Name define the name for this service
func (provider *HtdEnvProvider) Name() string {
	return contract.EnvKey
}
