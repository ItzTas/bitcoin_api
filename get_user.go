package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
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
			respondWithJSON(w, http.StatusNotFound, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToReturnValsUSer(dbuser))
}

func (cfg *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	limitstr := r.URL.Query().Get("limit")
	if limitstr == "" {
		limitstr = "20"
	}

	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid limit query")
		return
	}

	dbusers, err := cfg.DB.GetUsers(context.TODO(), int64(limit))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get users")
		return
	}

	users := funk.Map(dbusers, func(dbu database.User) returnValsUser {
		return databaseUserToReturnValsUSer(dbu)
	}).([]returnValsUser)

	respondWithJSON(w, http.StatusOK, users)
}
