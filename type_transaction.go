package main

import (
	"time"

	"github.com/ItzTas/coinerAPI/internal/database"
	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Amount     string    `json:"amount"`
	ExecutedAt time.Time `json:"executed_at"`
}

func databaseTransactionToTransaction(transaction database.Transaction) Transaction {
	return Transaction(transaction)
}
