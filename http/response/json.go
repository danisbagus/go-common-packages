package response

import (
	"encoding/json"
	"net/http"
)

const JsonContentType = "application/json"

type ErrorData struct {
	ErrorID int    `json:"error_id"`
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
	w.Header().Add("Content-Type", JsonContentType)
	w.WriteHeader(code)

	errorData := ErrorData{ErrorID: code, Message: message}
	Write(w, code, errorData)
}
