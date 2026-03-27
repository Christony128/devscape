package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"database/sql"
	"devscape/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt" //auth
	"github.com/golang-jwt/jwt/v5" //auth
	"time" //auth
	"os" //auth
)
type Handler struct {
	DB *sql.DB
}

// Register a new user and return a token immediately
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    if user.Username == "" || user.Password == "" {
        http.Error(w, "Username and password required", http.StatusBadRequest)
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

    // 1. Use QueryRow with 'RETURNING id' to grab the newly created database ID
    var newUserID int64
    err := h.DB.QueryRow(
        "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
        user.Username, string(hashed),
    ).Scan(&newUserID)
    
    if err != nil {
        http.Error(w, "Username might already exist", http.StatusConflict)
        return
    }

    // 2. Generate the token using that new ID
    token := GenerateJWT(newUserID, user.Username)

    // 3. Send the token back as JSON instead of the plain text message
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}

// Login authenticates a user and returns a token
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    // Decode the incoming JSON (username and password)
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // 1. Fetch the user's ID and Hashed Password from the database
    var dbUser models.User
    err := h.DB.QueryRow(
        "SELECT id, password FROM users WHERE username=$1",
        user.Username,
    ).Scan(&dbUser.ID, &dbUser.Password)

    // If the query fails, the username doesn't exist in our DB
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // 2. Compare the hashed password from the DB with the plain-text one they just sent
    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
    if err != nil {
        // Password doesn't match
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // 3. Everything matches! Generate the JWT using the ID we just scanned
    token := GenerateJWT(dbUser.ID, user.Username)

    // 4. Send the token back as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}

// GenerateJWT is the helper function your Register and Login handlers are looking for
func GenerateJWT(userID int64, username string) string {
    // Create the claims (data inside the token)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    // Get your secret key from the .env file
    secretKey := os.Getenv("JWT_SECRET")
    if secretKey == "" {
        // If you forgot to set this in .env, the app will crash here with a clear message
        panic("JWT_SECRET environment variable is not set!")
    }

    // Sign the token and return the string
    tokenStr, _ := token.SignedString([]byte(secretKey))
    return tokenStr
}










func(h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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