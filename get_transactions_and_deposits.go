package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

func (cfg *apiConfig) handlerGetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.PathValue("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	limitstr := r.URL.Query().Get("limitstr")
	if limitstr != "" {
		cfg.handlerGetUserTransactionsWithLimit(w, limitstr, userID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	dbtransactions, err := cfg.DB.GetUserTransactions(ctx, userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get user transactions")
		return
	}

	transactions := funk.Map(dbtransactions, func(transaction database.Transaction) Transaction {
		return databaseTransactionToTransaction(transaction)
	}).([]Transaction)

	respondWithJSON(w, http.StatusOK, transactions)
}

func (cfg *apiConfig) handlerGetUserTransactionsWithLimit(w http.ResponseWriter, limitstr string, userID uuid.UUID) {
	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid limit param")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	gutwlParams := database.GetUserTransactionsWithLimitParams{
		SenderID: userID,
		Limit:    int64(limit),
	}

	dbtransactions, err := cfg.DB.GetUserTransactionsWithLimit(ctx, gutwlParams)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get user transactions")
		return
	}

	transactions := funk.Map(dbtransactions, func(transaction database.Transaction) Transaction {
		return databaseTransactionToTransaction(transaction)
	}).([]Transaction)

	respondWithJSON(w, http.StatusOK, transactions)
}
