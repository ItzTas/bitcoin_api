package main

import (
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

type Transaction struct {
	ID             uuid.UUID `json:"id"`
	SenderID       uuid.UUID `json:"sender_id"`
	ReceiverID     uuid.UUID `json:"receiver_id"`
	Amount         string    `json:"amount"`
	ExecutedAt     time.Time `json:"executed_at"`
	IsBetweenUsers bool      `json:"is_between_users"`
}

func databaseTransactionToTransaction(transaction database.Transaction) Transaction {
	return Transaction(transaction)
}
