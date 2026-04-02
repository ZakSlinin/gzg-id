package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Username     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	FirstName    string    `json:"first_name"`
	Surname      string    `json:"surname"`
	IsVerified   bool      `json:"is_verified"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email     string
	Password  string
	Username  string
	Avatar    string
	FirstName string
	Surname   string
}

type EmailVerificationToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}
