package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteJsonHeader(w http.ResponseWriter, status int) http.ResponseWriter {
	w.WriteHeader(status)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	return w
}

func MakeResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

func JsonErrorResponse(w http.ResponseWriter, status int, message string) {
	w = WriteJsonHeader(w, status)
	json.NewEncoder(w).Encode(MakeResponse(status, message))
}
