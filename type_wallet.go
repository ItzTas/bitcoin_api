package main

import (
	"time"

	"github.com/ItzTas/coinerAPI/internal/database"
	"github.com/google/uuid"
)

type Wallet struct {
	ID           uuid.UUID `json:"id"`
	OwnerID      uuid.UUID `json:"owner_id"`
	CryptoTypeID string    `json:"crypto_type_id"`
	BalanceUsd   string    `json:"balance_usd"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func databaseWalletToWallet(wallet database.Wallet) Wallet {
	return Wallet(wallet)
}
