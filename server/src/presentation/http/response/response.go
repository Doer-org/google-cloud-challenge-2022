package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, j interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
