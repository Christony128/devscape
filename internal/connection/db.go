package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Database Config Error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database Connection Failed. Is Postgres running?:", err)
	}

	fmt.Println("Database Connection: ONLINE")
	return db
}