package response

import "net/http"

type ErrJson struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors string `json:"errors"`
}

// TODO: なぜか200でかえる
func NewErrJson(code int, status string, err error) *ErrJson {
	return &ErrJson{
		Code:   code,
		Status: status,
		Errors: err.Error(),
	}
}

func New404ErrJson(err error) *ErrJson {
	return &ErrJson{
		Code:   http.StatusBadRequest,
		Status: "StatusBadRequest",
		Errors: err.Error(),
	}
}
