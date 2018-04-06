package auth

import (
	"seniors50plus/internal/user"
)

type Error struct {
	Message string `json:"error"`
}

type TemplateInfo struct {
	Firstname string
	URL       string
}

type AuthResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
}

type RegistrationRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
