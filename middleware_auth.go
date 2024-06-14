package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ItzTas/bitcoinAPI/internal/auth"
	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

type handlerSecure = func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler handlerSecure) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dbUser, err := cfg.getUserByToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		handler(w, r, dbUser)
	}
}

func (cfg *apiConfig) getUserByToken(header http.Header) (database.User, error) {
	token, err := auth.GetBearerToken(header)
	if err != nil {
		return database.User{}, err
	}
	idstr, err := auth.GetIDByToken(token, cfg.jwtSecret)
	if err != nil {
		return database.User{}, errors.New("could not get id")
	}
	id, err := uuid.Parse(idstr)
	if err != nil {
		return database.User{}, errors.New("could not parse id")
	}
	dbuser, err := cfg.DB.GetUserByID(context.TODO(), id)
	if err != nil {
		return database.User{}, errors.New("could not get user")
	}
	return dbuser, nil
}

func secureEndpoint(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type returnVal struct {
		Hello string `json:"hello"`
	}
	respondWithJSON(w, http.StatusOK, returnVal{
		Hello: fmt.Sprintf("Hello there! %v", dbUser.UserName),
	})
}
