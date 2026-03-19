package main

import (
	"fmt"
	"net/http"
	
	"server/internal/connection"
	"server/internal/router"
)

func main() {
	db := database.InitDB()
	defer db.Close()
	r := router.Setup()

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}