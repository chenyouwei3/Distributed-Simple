package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewDecoder(buf)
	err := enc.Decode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceURL, "application/json", buf)
	if err != nil {
		return err
	}
	//判断业务是否请求成功
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("")
	}
	return nil
}
