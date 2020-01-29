// Copyright (c) 2020 Vorotynsky Maxim

package controllers

import (
	"encoding/json"
	"microAuth/data"
	"microAuth/model"
	"microAuth/server"
	"net/http"
	"time"
)

type JwtController struct {
	Duration    time.Duration
	MaxDuration time.Duration
	Repository  data.UserRepository
	secretKey   model.JwtSecret
}

type (
	loginData struct {
		UserName, Password string
	}
	tokenResourse struct {
		Token model.Token `json:"jwt_auth_token"`
	}
)

func NewJwtController(duration time.Duration, maxDuration time.Duration, repository data.UserRepository, secretKey model.JwtSecret) JwtController {
	return JwtController{duration, maxDuration, repository, secretKey}
}

func (controller *JwtController) Login(w http.ResponseWriter, r *http.Request) {
	var data loginData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		server.DisplayError(w, err, "Invalid login data format. ", http.StatusBadRequest)
		return
	}

	var user model.User
	user, err := controller.Repository.GetUserByName(data.UserName)
	if err != nil || user.ConfirmPassword(data.Password) != nil {
		server.DisplayError(w, err, "Invalid login data. ", http.StatusUnauthorized)
		return
	}

	token, err := user.GenerateToken(controller.secretKey, controller.Duration)
	if err != nil {
		server.DisplayError(w, err, "Token generating error. ", http.StatusInternalServerError)
		return
	}
	writeToken(token, w)
}

func (controller *JwtController) Refresh(w http.ResponseWriter, r *http.Request) {
	var token tokenResourse
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		server.DisplayError(w, err, "Invalid token. ", http.StatusUnauthorized)
		return
	}

	if err := token.Token.Refresh(controller.secretKey, controller.Duration); err != nil {
		server.DisplayError(w, err, "Token refresh error. ", http.StatusInternalServerError)
		return
	}
	writeToken(token.Token, w)
}

func writeToken(token model.Token, w http.ResponseWriter) error {
	data, err := json.Marshal(tokenResourse{token})
	if err != nil {
		server.DisplayError(w, err, "Something goes wrong...", http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}
