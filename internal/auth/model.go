package auth

import (
	"seniors50plus/internal/user"
)

type Error struct {
	message string `json:"message"`
}

type AuthResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
	Error `json:"error"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
