package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

type GetUserTransactionsRow struct {
	ID         uuid.UUID `json:"id"`
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Amount     string    `json:"amount"`
	ExecutedAt time.Time `json:"executed_at"`
	UserRole   string    `json:"user_role"`
}

func (cfg *apiConfig) handlerGetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.PathValue("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	limitstr := r.URL.Query().Get("limit")
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

	transactions := funk.Map(dbtransactions, func(transaction database.GetUserTransactionsRow) GetUserTransactionsRow {
		return GetUserTransactionsRow(transaction)
	}).([]GetUserTransactionsRow)

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
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusBadRequest, "user does not have any transaction")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get user transactions")
		return
	}

	transactions := funk.Map(dbtransactions, func(transaction database.GetUserTransactionsWithLimitRow) GetUserTransactionsRow {
		return GetUserTransactionsRow(transaction)
	}).([]GetUserTransactionsRow)

	respondWithJSON(w, http.StatusOK, transactions)
}
func (cfg *apiConfig) handlerGetUserDeposits(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.PathValue("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	limitstr := r.URL.Query().Get("limit")
	if limitstr != "" {
		cfg.handlerGetUserDepositsWithLimit(w, limitstr, userID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	dbdeposits, err := cfg.DB.GetDepositsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusBadRequest, "user does not have any deposits")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get user deposits")
		return
	}

	deposits := funk.Map(dbdeposits, func(dbdeposit database.Deposit) Deposit {
		return databaseDepositToDeposit(dbdeposit)
	}).([]Deposit)

	respondWithJSON(w, http.StatusOK, deposits)
}

func (cfg *apiConfig) handlerGetUserDepositsWithLimit(w http.ResponseWriter, limitstr string, userID uuid.UUID) {
	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid limit")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	gudbiwlParams := database.GetDepositsByUserIDWithLimitParams{
		OwnerID: userID,
		Limit:   int64(limit),
	}
	dbdepodits, err := cfg.DB.GetDepositsByUserIDWithLimit(ctx, gudbiwlParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusBadRequest, "user does not have any deposits")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get user deposits")
		return
	}

	deposits := funk.Map(dbdepodits, func(dbdeposit database.Deposit) Deposit {
		return databaseDepositToDeposit(dbdeposit)
	}).([]Deposit)

	respondWithJSON(w, http.StatusOK, deposits)
}
