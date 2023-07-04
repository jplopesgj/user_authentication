package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Users []User

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}
