package contract

const (
	envProduction  = "production"
	EnvTesting     = "Testing"
	EnvDevelopment = "development"
	EnvKey         = "htd:env"
)

type Env interface {
	// AppEnv 获取当前环境
	AppEnv() string
	IsExist(string) bool
	// Get 获取环境变量，没有返回空
	Get(string) string
	// All 获取所有环境变量，最终结果
	All() map[string]string
}
