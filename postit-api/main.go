package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")

	cfg := mysql.Config{
		User:   DB_USER,
		Passwd: DB_PASS,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", DB_HOST, DB_PORT),
		DBName: DB_NAME,
	}

	sqlStorage := NewMySqlStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStorage(db)

	apiServer := NewApiServer(":3000", store)
	err = apiServer.Start()
	if err != nil {
		log.Fatal(err)
	}
}
