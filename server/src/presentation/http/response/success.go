package response

import "net/http"

type SuccessJson struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewSuccessJson(code int, status string, message string) *SuccessJson {
	return &SuccessJson{
		Code:    code,
		Status:  status,
		Message: message,
	}
}

func New200SuccessJson(message string) *SuccessJson {
	return &SuccessJson{
		Code:    http.StatusOK,
		Status:  "StatusOK",
		Message: message,
	}
}
