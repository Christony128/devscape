package models

import(
	"time"
	"github.com/google/uuid"
)

type User struct {
  ID string `json:”user_id”`
  username string `json:”username”`
  password string `json:”password”`
}

type Todo struct {
	ID   int64   `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UserId  uuid.UUID `json:"userID"`
}