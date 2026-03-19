package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	dsn := "postgres://postgres:root@localhost:5433/test?sslmode=disable"
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Database Config Error:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}
	fmt.Println("Database Connection: ONLINE")
	return db
}