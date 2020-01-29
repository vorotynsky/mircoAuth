// Copyright (c) 2020 Vorotynsky Maxim

package routers

import (
	"github.com/gorilla/mux"
	"microAuth/controllers"
	"microAuth/model"
	"microAuth/server"
	"time"
)

var userController controllers.UserController
var jwtController controllers.JwtController

func initControllers() {
	userController = controllers.UserController{server.UserRepository}
	jwtController = controllers.NewJwtController(
		time.Duration(server.Configuration.JwtConfig.DefaultDuration),
		time.Duration(server.Configuration.JwtConfig.MaxDuration),
		server.UserRepository,
		model.NewJwtSecret(server.Configuration.JwtConfig.JwtSecret))
}

func InitRouters() *mux.Router {
	initControllers()
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRouter(router, userController)
	router = SetAuthRouter(router, jwtController)
	return router
}
