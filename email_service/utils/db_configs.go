package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DBClient *sqlx.DB

func InitDatabase(username string, password string, hostname string, dbName string) {
	dsn:= fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Could not create database manager:", err.Error())
	}

	err = db.Ping()
	if  err != nil{
		log.Fatalln("Could not ping database", err.Error())
	}

	DBClient = db
}