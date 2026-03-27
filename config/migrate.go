package config

import(
	"log"
	"database/sql"
)

func CreateTable(db *sql.DB){
	_,err:=db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
	    ID SERIAL PRIMARY KEY,
		Task TEXT NOT NULL,
		Done BOOLEAN DEFAULT FALSE,
		CreatedAt TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
	);`)

	if err!=nil{
		log.Fatal("Failed to create tables:",err)
	}

	_,err=db.Exec(`CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT
	);`)

	if err!=nil{
		log.Fatal("Failed to create tables:",err)
	}
}