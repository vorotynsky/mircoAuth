// Copyright (c) 2020 Vorotynsky Maxim

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/zpatrick/go-config"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func signalInterruptHandle(handler func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		handler()
	}()
}

func main() {
	iniFile := config.NewINIFile("config.ini")
	c := config.NewConfig([]config.Provider{iniFile})

	port, _ := c.Int("server.port")
	host := fmt.Sprintf("localhost:%d", port)

	handler := msg("hello from go")

	httpSrv := &http.Server{Addr: host, Handler: handler}

	httpWG := &sync.WaitGroup{}
	httpWG.Add(1)

	go func(server *http.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}(httpSrv, httpWG)

	fmt.Printf("Server is listening on http://%s\n", host)
	defer fmt.Println("Server has stopped listening.")

	signalInterruptHandle(func() { httpSrv.Close() })
	httpWG.Wait()
}
