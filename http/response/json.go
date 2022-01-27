package response

import (
	"encoding/json"
	"net/http"
)

const JsonContentType = "application/json"

type ErrorData struct {
	Message string `json:"message"`
}

func Write(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", JsonContentType)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func Error(w http.ResponseWriter, code int, message string) {
	errorData := ErrorData{Message: message}
	Write(w, code, errorData)
}
