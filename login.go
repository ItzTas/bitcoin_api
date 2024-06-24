package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ItzTas/coinerAPI/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultExpiration = 8 * time.Hour
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type returnVals struct {
		returnValsUser
		Token string `json:"token"`
	}

	params := paramethers{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode params")
		return
	}

	dbuser, err := cfg.DB.GetUserByEmail(context.TODO(), params.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "User does not exist")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "could not get user")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(params.Password)); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Passwords don't match")
		return
	}

	token, err := auth.NewJWT(dbuser, cfg.jwtSecret, DefaultExpiration)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not create token")
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		returnValsUser: databaseUserToReturnValsUser(dbuser),
		Token:          token,
	})
}
