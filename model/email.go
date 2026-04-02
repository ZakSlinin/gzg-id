package model

import "time"

type EmailVerification struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}
