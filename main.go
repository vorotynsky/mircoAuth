// Copyright (c) 2020 Vorotynsky Maxim

package main

import (
	"fmt"
	"net/http"

	"github.com/zpatrick/go-config"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	iniFile := config.NewINIFile("config.ini")
	c := config.NewConfig([]config.Provider{iniFile})

	port, _ := c.Int("server.port")
	host := fmt.Sprintf("localhost:%d", port)

	handler := msg("hello from go")
	fmt.Printf("Server is listening on http://%s\n", host)
	http.ListenAndServe(host, handler)
}
