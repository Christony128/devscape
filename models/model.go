package models

import(
	"time"
)

type User struct {
  ID int64 `json:"user_id"`
  Username string `json:"username"`
  Password string `json:"password"`
}

type Todo struct {
	ID   int64   `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}