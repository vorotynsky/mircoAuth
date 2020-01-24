// Copyright (c) 2020 Vorotynsky Maxim

package server

import (
	"database/sql"
	"log"
	"microAuth/data"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB
var UserRepository data.UserRepository

func initDatabase() (err error) {
	database, err = sql.Open("mysql", Configuration.ConnString)
	UserRepository = data.SqlUserRepository{GetDatabaseConnection}
	return
}

func GetDatabaseConnection() *sql.DB {
	if database == nil {
		if err := initDatabase(); err != nil {
			log.Fatalln("[GetDatabaseConnection]:", err)
		}
	}
	return database
}

func CloseDatabaseConnection() {
	if database != nil {
		database.Close()
		database = nil
	}
}
