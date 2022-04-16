package views

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}
