package config

import(
	"log"
	"os"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() *sql.DB{
	// connectionStr :="postgres://postgres:postgres@localhost:5432/todoapp"
	// connectionStr :="postgres://postgres:root@localhost:5433/test?sslmode=disable"
	connectionStr := os.Getenv("DB_URL")

	//this DB returned here is a pool of zero or more underlying connections and its safe for concurrent use by multiple go routines
	db,err:=sql.Open("pgx",connectionStr)
	if err!=nil{
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the postgrSQL server ")

	return db
}