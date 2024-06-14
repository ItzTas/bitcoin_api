package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/auth"
	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

func isErrorCode(err error, errcode pq.ErrorCode) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == errcode
	}
	return false
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode params")
		return
	}

	hashsedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not hash password")
		return
	}

	id := uuid.New()
	cuParams := database.CreateUserParams{
		ID:        id,
		UserName:  params.UserName,
		Email:     params.Email,
		Password:  hashsedPassword,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Currency:  "0",
	}

	dbuser, err := cfg.DB.CreateUser(context.TODO(), cuParams)
	if err != nil {
		if isErrorCode(err, UniqueViolationErr) {
			respondWithError(w, http.StatusBadRequest, "User already exists")
			return
		}
		respondWithError(w, http.StatusBadRequest, "Could not create user")
		return
	}

	user := databaseUserToUser(dbuser)

	respondWithJSON(w, http.StatusCreated, user)
}
