package response

type ErrResponse struct {
	Message string `json:"message"`
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Message: err.Error(),
	}
}
