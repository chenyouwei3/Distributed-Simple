package registry

// 注册舒服用的结构体参数的格式
type registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
)
