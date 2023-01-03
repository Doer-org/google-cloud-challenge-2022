package response

type Response struct {
	Message string `json:"message"`
}

func NewResponse(message string) *Response {
	return &Response{
		Message: message,
	}
}
