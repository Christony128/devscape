package main 
//model.go - added user id in ToDo struct
//route.go - added routes for register and login
//handler.go - added handler functions for register and login
//DROP TASKS AND USERS TABLES
import (
	"net/http"
	"devscape/config"
	"devscape/handlers"
	"devscape/routes"
	"github.com/joho/godotenv"
	"fmt"
	"github.com/rs/cors"
)

func main(){
	godotenv.Load()
	db:=config.ConnectDB()
	config.CreateTable(db)
	h:=&handlers.Handler{DB:db}
	r:=routes.SetupRoutes(h)

	c := cors.New(cors.Options{ //auth
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},             
		AllowCredentials: true,                                                
		Debug:            true, 
	}) //auth

	// 5. Wrap the router with CORS
	handler := c.Handler(r)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",handler) //PUT HANDLER HERE
}