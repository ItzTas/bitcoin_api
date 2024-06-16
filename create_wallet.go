package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateWallet(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type paramethers struct {
		CryptoTypeID string `json:"crypto_type_id"`
	}

	params := paramethers{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode params")
		return
	}

	id := uuid.New()
	cwParams := database.CreateWalletParams{
		ID:           id,
		OwnerID:      dbUser.ID,
		CryptoTypeID: params.CryptoTypeID,
		BalanceUsd:   "0",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	wallet, err := cfg.DB.CreateWallet(ctx, cwParams)
	if err != nil {
		if isErrorCode(err, UniqueViolationErr) {
			respondWithError(w, http.StatusBadRequest, "Wallet already exists")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create wallet")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseWalletToWallet(wallet))
}
