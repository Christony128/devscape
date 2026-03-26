package main 

import (
	"net/http"
	"devscape/config"
	"devscape/handlers"
	"devscape/routes"
	"fmt"
)

func main(){
	db:=config.ConnectDB()
	config.CreateTable(db)
	h:=&handlers.Handler{DB:db}
	r:=routes.SetupRoutes(h)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",r)
}