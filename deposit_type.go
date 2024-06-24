package main

import (
	"time"

	"github.com/ItzTas/coinerAPI/internal/database"
	"github.com/google/uuid"
)

type Deposit struct {
	ID         uuid.UUID `json:"id"`
	WalletID   uuid.UUID `json:"wallet_id"`
	Amount     string    `json:"amount"`
	ExecutedAt time.Time `json:"executed_at"`
}

func databaseDepositToDeposit(deposit database.Deposit) Deposit {
	return Deposit(deposit)
}
