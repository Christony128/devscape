package router

import (
	"database/sql" // import
	"net/http"
	"server/internal/handlers"
)
//add for db
func Setup(db *sql.DB) *http.ServeMux { 
	mux := http.NewServeMux()

	// pass the db connection to handlers
	mux.HandleFunc("GET /users", handlers.GetUsers(db))
	mux.HandleFunc("POST /users", handlers.CreateUser(db))

	return mux
}