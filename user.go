package main

import (
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

type returnValsUser struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Currency  string    `json:"currency"`
}

func databaseUserToReturnValsUSer(user database.User) returnValsUser {
	return returnValsUser{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Currency:  user.Currency,
	}
}
