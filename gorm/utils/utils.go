package utils

func ResponseJson(data, Error any, code int) any {
	return struct {
		Data  any `json:"data,omitempty"`
		Code  int `json:"code,omitempty"`
		Error any `json:"error,omitempty"`
	}{
		Data:  data,
		Code:  code,
		Error: Error,
	}
}
