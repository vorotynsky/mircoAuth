// Copyright (c) 2020 Vorotynsky Maxim

package main

import (
	"log"
	"microAuth/routers"
	"microAuth/server"
	"net/http"
	"sync"
)

func main() {
	if err := server.SetUp(); err != nil {
		panic("Server set up error.")
	}
	defer server.CloseDatabaseConnection()

	handler := routers.InitRouters()
	httpSrv := &http.Server{Addr: server.Configuration.Host, Handler: handler}

	// Run the server

	httpWG := &sync.WaitGroup{}
	httpWG.Add(1)

	server.StartServer(httpSrv, httpWG)

	log.Printf("Server is listening on http://%s\n", server.Configuration.Host)
	defer log.Println("Server has stopped listening.")

	httpWG.Wait()
}
