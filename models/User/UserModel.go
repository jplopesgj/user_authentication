package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserFields struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserFields struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedAt string `json:"updated_at"`
}
