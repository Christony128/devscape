package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Fetched users",
		})
}
	
func CreateUser(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User received",
			"name":    body["name"],
		})
}