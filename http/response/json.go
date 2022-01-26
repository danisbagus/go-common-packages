package response

import (
	"encoding/json"
	"net/http"
)

const JsonContentType = "application/json"

func Write(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", JsonContentType)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
