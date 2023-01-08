package response

type ErrResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Errors  string `json:"errors"`
}

func NewErrResponse(code int,status string,err error) *ErrResponse {
	return &ErrResponse{
		Code: code,
		Status: status,
		Errors: err.Error(),
	}
}
