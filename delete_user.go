package main

import (
	"context"
	"net/http"

	"github.com/ItzTas/bitcoinAPI/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.PathValue("user_id")
	if userIDstr == "" {
		respondWithError(w, http.StatusBadRequest, "Empty user id")
		return
	}

	if err := auth.AuthenticateDeleteKey(r.Header, cfg.deleteCodeSecret); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid")
		return
	}

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	_, err = cfg.DB.DeleteUser(ctx, userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not delete")
		return
	}

	respondWithNoContent(w)
}
