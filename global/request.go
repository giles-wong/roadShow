package global

type Request struct {
	Params map[string]interface{}
}

func (r *Request) GetParams() map[string]interface{} {
	return r.Params
}

// MustParam 必须要存在的参数字段
func (r *Request) MustParam() []string {
	var mustParam = [...]string{"app_key", "app_version", "method", "timestamp", "timeout", "noncestr", "sign"}

	return mustParam[:]
}
