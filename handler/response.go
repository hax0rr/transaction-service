package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"errors,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

func GenerateResponse(w http.ResponseWriter, data interface{}, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var resp Response
	if err != nil {
		resp.Error = &Error{Message: err.Error()}
	} else {
		resp.Data = data
	}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
