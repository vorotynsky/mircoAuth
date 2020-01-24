// Copyright (c) 2020 Vorotynsky Maxim

package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"microAuth/controllers"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle!")
}

func SetUserRouter(router *mux.Router, controller controllers.UserController) *mux.Router {
	router.HandleFunc("/user/id/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/user/name/{name}", controller.GetUserByName).Methods("GET")
	router.HandleFunc("/user/create", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/delete/id/{id}", controller.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/user/delete/name/{name}", controller.DeleteUserByUserName).Methods("DELETE")
	return router
}
