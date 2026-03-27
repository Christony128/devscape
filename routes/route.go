package routes

import (
    "devscape/handlers"
    "devscape/middleware" //MIDDLEWARE
    "net/http" //MIDDLEWARE
    "github.com/gorilla/mux"
)

// type Router struct {
// 	NotFoundHandler http.Handler
// 	MethodNotAllowedHandler http.Handler
// 	KeepContext bool
// }

func SetupRoutes(h *handlers.Handler) *mux.Router {
	//NewRouter this returns an instance of a router
    
    r := mux.NewRouter()
    r.HandleFunc("/register", h.Register).Methods("POST")//auth
    r.HandleFunc("/login", h.Login).Methods("POST")//auth

    
    r.Handle("/todos", middleware.AuthMiddleware(http.HandlerFunc(h.PostTask))).Methods("POST") //middleware
	r.Handle("/todos", middleware.AuthMiddleware(http.HandlerFunc(h.GetTasks))).Methods("GET")  //middleware
	r.Handle("/todos/{id}/done", middleware.AuthMiddleware(http.HandlerFunc(h.MarkDone))).Methods("PATCH") //middleware
	r.Handle("/todos/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.DeleteTask))).Methods("DELETE")    //middleware
    return r
}
