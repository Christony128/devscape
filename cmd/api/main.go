package main

import (
	"fmt"
	"net/http"
	"server/internal/router"
)

func main() {

	r := router.Setup()


	fmt.Println("🚀 Server running on http://localhost:8080")
	

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}