package controllers

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	// dbHost := os.Getenv("DB_HOST")
	// fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp_tools?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
