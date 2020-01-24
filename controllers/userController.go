// Copyright (c) 2020 Vorotynsky Maxim

package controllers

import (
	"encoding/json"
	"errors"
	"microAuth/data"
	"microAuth/model"
	"microAuth/server"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	Repository data.UserRepository
}

type (
	userResource struct {
		User model.User `json:"data"`
	}
	userDataResource struct {
		UserData model.UserData `json:"data"`
	}
	userWithPasswordResource struct {
		User model.UserWithPurePassword `json:"data"`
	}
)

func (controller *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, errFlag := controller.parseId(vars, w)
	if errFlag {
		return
	}

	if user, err := controller.Repository.GetUserById(userId); err == nil {
		writeUser(user.UserData, w)
		return
	}
	http.NotFound(w, r)
}

func (controller *UserController) GetUserByName(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["name"]
	if user, err := controller.Repository.GetUserByName(userName); err == nil {
		writeUser(user.UserData, w)
		return
	}
	http.NotFound(w, r)
}

func (controller *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, errFlag := controller.parseId(vars, w)
	if errFlag {
		return
	}

	err := controller.Repository.DeleteById(userId)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) DeleteUserByUserName(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["name"]

	err := controller.Repository.DeleteByName(userName)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dataResource userWithPasswordResource
	if err := json.NewDecoder(r.Body).Decode(&dataResource); err != nil {
		server.DisplayError(w, err, "Invalid user data", http.StatusBadRequest)
		return
	}
	user := dataResource.User.HashPassword()
	if user == nil {
		server.DisplayError(w, errors.New("User password isn't hashed."),
			"Oops... Registration has been failed.", http.StatusInternalServerError)
		return
	}
	if err := controller.Repository.Create(user); err != nil {
		server.DisplayError(w, err, "Oops... Registration has been failed.", http.StatusInternalServerError)
		return
	}
	writeUser(dataResource.User.UserData, w)
}

func (controller *UserController) parseId(vars map[string]string, w http.ResponseWriter) (int32, bool) {
	userId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		server.DisplayError(w, err, "Invalid user id type.", http.StatusBadRequest)
		return 0, true
	}
	return int32(userId), false
}

func writeUser(user model.UserData, w http.ResponseWriter) error {
	data, err := json.Marshal(userDataResource{user})
	if err != nil {
		server.DisplayError(w, err, "Something goes wrong...", http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}
