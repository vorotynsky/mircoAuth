// Copyright (c) 2020 Vorotynsky Maxim

package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func StartServer(server *http.Server, wg *sync.WaitGroup) (err error) {
	signalInterruptHandle(func() { server.Close() })

	go func() {
		defer wg.Done()
		if err = server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		} else {
			err = nil
		}
	}()
	return
}

func signalInterruptHandle(handler func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		handler()
	}()
}
