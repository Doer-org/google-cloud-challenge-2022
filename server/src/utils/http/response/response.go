package response

import (
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/utils/logger"
)

func WriteJsonResponse(w http.ResponseWriter, j any, status int) {
	b, err := json.Marshal(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		logger.Println(err.Error())
	}
}
