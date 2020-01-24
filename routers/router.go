// Copyright (c) 2020 Vorotynsky Maxim

package routers

import (
	"github.com/gorilla/mux"
	"microAuth/controllers"
	"microAuth/server"
)

var userController controllers.UserController

func initControllers() {
	userController = controllers.UserController{server.UserRepository}
}

func InitRouters() *mux.Router {
	initControllers()
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRouter(router, userController)
	return router
}
