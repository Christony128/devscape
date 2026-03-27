package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"database/sql"
	"devscape/models"
	"github.com/gorilla/mux"
	"github.com/golang-jwt/jwt/v5"
)

// Create Todo
//query -- runs a SQL statement that returns multiple rows 
//queryRow -- runs a SQL statement that returns one rows

type Handler struct {
	DB *sql.DB
}

//this actually register the user 
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.password), 10)
	_, err := h.DB.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.sername, string(hashed),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("User created"))
}

func GenerateJWT(userID int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte("secret"))
	return tokenStr
}

//for login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	var dbUser models.User
	err := h.DB.QueryRow(
		"SELECT id, password FROM users WHERE username=$1",
		user.Username,
	).Scan(&dbUser.ID, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid user", 401)
		return
	}
	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Wrong password", 401)
		return
	}
	token := GenerateJWT(dbUser.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var todo models.Task
	json.NewDecoder(r.Body).Decode(&todo)

	// get user from middleware
	userID := r.Context().Value("userID").(int64)

	err := h.DB.QueryRow(
		"INSERT INTO tasks (task, user_id) VALUES ($1, $2) RETURNING id, task, createdat",
		todo.Task, userID,
	).Scan(&todo.ID, &todo.Task, &todo.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

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