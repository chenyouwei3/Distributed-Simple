package registry

// 服务注册用的结构体参数的格式

type Registration struct {
	ServiceName      ServiceName   //服务的name
	ServiceURL       string        //服务的url
	RequiredServices []ServiceName //当前服务依赖的其他服务
	ServiceUpdateURL string        //依赖更新请求地址
	HeartBeatURL     string        //心跳协议地址
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
