package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

// 为日志文件做准备
var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(date []byte) (int, error) {
	//没有文件并创建
	file, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(date)
}

// Run 生成日志文件
func Run(destination string) {
	//地址,前缀,日期时间
	log = stlog.New(fileLog(destination), "[go]-", stlog.Flags())
}

//日志服务的逻辑

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
