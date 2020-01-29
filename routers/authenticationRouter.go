// Copyright (c) 2020 Vorotynsky Maxim

package routers

import (
	"github.com/gorilla/mux"
	"microAuth/controllers"
)

func SetAuthRouter(router *mux.Router, controller controllers.JwtController) *mux.Router {
	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.HandleFunc("/auth/refresh", controller.Refresh).Methods("POST")
	return router
}
