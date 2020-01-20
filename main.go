// Copyright (c) 2020 Vorotynsky Maxim

package main

import (
	"fmt"
	"log"
	"microAuth/server"
	"net/http"
	"sync"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	server.SetUp()

	handler := msg("hello from go")
	httpSrv := &http.Server{Addr: server.Configuration.Host, Handler: handler}

	// Run the server

	httpWG := &sync.WaitGroup{}
	httpWG.Add(1)

	server.StartServer(httpSrv, httpWG)

	log.Printf("Server is listening on http://%s\n", server.Configuration.Host)
	defer log.Println("Server has stopped listening.")

	httpWG.Wait()
}
