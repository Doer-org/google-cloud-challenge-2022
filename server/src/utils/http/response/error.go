package response

import (
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/utils/logger"
)

type ErrResponse struct {
	Message string `json:"message"`
}

// TODO: 成功時のステータスを引数に受け取り、
// ConvertToJsonResponseAndErrCheckかこちらで統一した処理にしたほうがよさそう
func NewErrResponse(w http.ResponseWriter, message error) {
	b, err := json.Marshal(
		ErrResponse{
			Message: message.Error(),
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(b); err != nil {
		logger.Println(err.Error())
	}
}
