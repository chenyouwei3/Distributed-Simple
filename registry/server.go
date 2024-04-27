package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServiceURL = "http://localhost" + ServerPort + "/services"

// 辅助用的结构体和函数
type registry struct {
	registrations []registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{
	registrations: make([]registration, 0),
	mutex:         new(sync.Mutex),
}

//逻辑

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("收到注册服务请求")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("已添加服务: %v 地址:%s\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
