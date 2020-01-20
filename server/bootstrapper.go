// Copyright (c) 2020 Vorotynsky Maxim

package server

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Host, ConnString string
}

var Configuration config

func SetUp() (err error) {
	if err = initConfig(); err != nil {
		log.Fatalln("[configuration]:", err)
	}
	return
}

func initConfig() (err error) {
	file, err := os.Open("server/config.json")
	defer file.Close()
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	Configuration = config{}
	err = decoder.Decode(&Configuration)
	return
}
