package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "127.0.0.1"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "REST"
)

func InitDB() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
	if err != nil {
		return nil, err 
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	_, err := InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database table initialized successfully")
}