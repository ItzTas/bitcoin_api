package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ItzTas/coinerAPI/internal/auth"
	"github.com/ItzTas/coinerAPI/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type paramaters struct {
		NewPassword *string `json:"new_password"`
		NewEmail    *string `json:"new_email"`
		NewUserName *string `json:"new_user_name"`
	}

	params := paramaters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode params")
		return
	}

	email := dbUser.Email
	if params.NewEmail != nil {
		email = *params.NewEmail
	}

	userName := dbUser.UserName
	if params.NewUserName != nil {
		userName = *params.NewUserName
	}

	var hashedPassword string
	var err error
	if params.NewPassword != nil {
		hashedPassword, err = auth.HashPassword(*params.NewPassword)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not hash password")
			return
		}
	} else {
		hashedPassword = dbUser.Password
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	if ctx.Err() != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("context error: %v", ctx.Err().Error()))
		return
	}

	upUparams := database.UpdateUserParams{
		Password:  hashedPassword,
		Email:     email,
		UserName:  userName,
		UpdatedAt: time.Now().UTC(),
		ID:        dbUser.ID,
	}

	_, err = cfg.DB.UpdateUser(ctx, upUparams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not update user:\n %v", err))
		return
	}

	respondWithNoContent(w)
}
