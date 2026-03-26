package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"database/sql"
	"devscape/models"
	"github.com/gorilla/mux"
)

// Create Todo
//query -- runs a SQL statement that returns multiple rows 
//queryRow -- runs a SQL statement that returns one rows

type Handler struct {
	DB *sql.DB
}

func(h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Insert into DB
	err = h.DB.QueryRow(
		"INSERT INTO tasks (task) VALUES ($1) RETURNING id,task,createdat",
		todo.Task,
	).Scan(&todo.ID,&todo.Task,&todo.CreatedAt)
 	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// Get all Todos
func(h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		rows.Scan(&t.ID, &t.Task, &t.Done,&t.CreatedAt)
		todos = append(todos, t)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Mark Task as Done
func(h *Handler) MarkDone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := h.DB.Exec(
		"UPDATE tasks SET done = TRUE WHERE id=$1",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated"))
}

// Delete Task
func(h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := h.DB.Exec(
		"DELETE FROM tasks WHERE id=$1",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Deleted"))
}