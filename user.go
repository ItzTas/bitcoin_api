package main

import (
	"time"

	"github.com/ItzTas/coinerAPI/internal/database"
	"github.com/google/uuid"
)

type returnValsUser struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Currency  string    `json:"currency"`
}

func databaseUserToReturnValsUser(user database.User) returnValsUser {
	return returnValsUser{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Currency:  user.Currency,
	}
}
