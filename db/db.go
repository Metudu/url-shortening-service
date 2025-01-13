package db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	database *sql.DB
	err error
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "secret123!"
	dbname = "url-shortening-service"
)

func InitDB() {
	log.Println("Connecting to the database...")
	database, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("An error occured when connecting to the database: %v", err.Error())
	}

	if err := database.Ping(); err != nil {
		log.Fatalf("An error occured when pinging the database: %v", err.Error())
	}
	log.Println("Connected!")
}

func GetDB() *sql.DB {
	return database
}