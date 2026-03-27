package config

import(
	"log"
	"database/sql"
)

func CreateTable(db *sql.DB){
	_,err:=db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
	    id SERIAL PRIMARY KEY,
		task TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE,
		createdat TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
		userID UUID REFERENCES users(id)
	);`)

	if err!=nil{
		log.Fatal("Failed to create tables:",err)
	}

	_,err=db.Exec(`CREATE TABLE IF NOT EXISTS users (
	    id TEXT PRIMARY KEY,
		username TEXT,
		password TEXT
	);`)

	if err!=nil{
		log.Fatal("Failed to create tables:",err)
	}
}