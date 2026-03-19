package router

import (
	"net/http"
	"server/internal/handlers"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("POST /users", handlers.CreateUser)

	return mux
}