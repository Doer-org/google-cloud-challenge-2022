package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, j any, status int) {
	if err := json.NewEncoder(w).Encode(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(status)
}
