package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")

	id, err := uuid.Parse(userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	dbuser, err := cfg.DB.GetUserByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithJSON(w, http.StatusBadRequest, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get user")
		return
	}

	user := databaseUserToUser(dbuser)

	type returnVals struct {
		ID        uuid.UUID `json:"id"`
		UserName  string    `json:"user_name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Currency  string    `json:"currency"`
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Currency:  user.Currency,
	})
}
