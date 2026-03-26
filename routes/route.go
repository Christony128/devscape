package routes

import (
    "devscape/handlers"
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
    r.HandleFunc("/todos", h.PostTask).Methods("POST")
    r.HandleFunc("/todos", h.GetTasks).Methods("GET")
	r.HandleFunc("/todos/{id}/done",h.MarkDone).Methods("PATCH")
	r.HandleFunc("/todos/{id}",h.DeleteTask).Methods("DELETE")
    return r
}
