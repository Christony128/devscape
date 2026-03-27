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
)

func main(){
	godotenv.Load()
	db:=config.ConnectDB()
	config.CreateTable(db)
	h:=&handlers.Handler{DB:db}
	r:=routes.SetupRoutes(h)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",r)
}