package demo

const Key = "htd:demo"

// Demo服务的接口
type Service interface {
	GetHello() Hi
}

// Demo服务接口定义的一个数据结构
type Hi struct {
	Name string
}
