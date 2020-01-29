// Copyright (c) 2020 Vorotynsky Maxim

package server

import (
	"encoding/json"
	"log"
	"microAuth/model"
	"os"
)

type (
	jwtConfig struct {
		JwtSecret       string
		DefaultDuration model.Duration
		MaxDuration     model.Duration
	}
	config struct {
		Host, ConnString string
		JwtConfig        jwtConfig
	}
)

var Configuration config

func SetUp() (err error) {
	if err = initConfig(); err != nil {
		log.Fatalln("[configuration]:", err)
	}
	if err = initDatabase(); err != nil {
		log.Fatalln("[database initialization]:", err)
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
