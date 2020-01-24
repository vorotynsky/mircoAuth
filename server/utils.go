// Copyright (c) 2020 Vorotynsky Maxim

package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Error appError `json:"data"`
	}
)

func DisplayError(w http.ResponseWriter, handlerError error, message string, code int) {
	err := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Println("[Error]:", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if data, err := json.Marshal(errorResource{Error: err}); err == nil {
		w.Write(data)
	}
}
