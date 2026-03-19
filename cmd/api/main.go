package main

import (
	"log"
	"net/http"
	"os"
	//db -> changes to main, db.go, env and router.go
	"github.com/joho/godotenv"
	"server/internal/connection"
	"server/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL not found in .env")
	}

	db := database.InitDB(dsn)
	defer db.Close()
	r := router.Setup(db)
	port := os.Getenv("PORT")
	
	log.Printf("🚀 Server starting on port %s", port)
	http.ListenAndServe(":"+port, r)
}