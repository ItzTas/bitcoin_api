package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Passwords don't match")
		return
	}

}
