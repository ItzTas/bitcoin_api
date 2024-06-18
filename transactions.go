package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) handlerSendToAccount(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type paramaters struct {
		SendQuantity string `json:"send_quantity"`
	}

	params := paramaters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode parameters")
		return
	}

	receiverIDstr := r.PathValue("receiver_id")
	if receiverIDstr == "" {
		respondWithError(w, http.StatusBadRequest, "no receiver id")
		return
	}

	currency, err := decimal.NewFromString(dbUser.Currency)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not convert currency to decimal")
		return
	}

	sendQuant, err := decimal.NewFromString(params.SendQuantity)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not convert send quantity to decimal")
		return
	}

	if currency.Compare(sendQuant) < 0 {
		respondWithError(w, http.StatusBadRequest, "Send quantity bigger than currency")
		return
	}

	if sendQuant.Compare(decimal.Zero) < 0 {
		respondWithError(w, http.StatusBadRequest, "The send quantity cannot be less than 0")
		return
	}

	receiverID, err := uuid.Parse(receiverIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse receiverID")
		return
	}

	transaction, err := cfg.makeUserTransaction(dbUser.ID, receiverID, sendQuant)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not make transaction: \n%v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseTransactionToTransaction(transaction))
}

func (cfg *apiConfig) makeUserTransaction(senderID, receiverID uuid.UUID, quantity decimal.Decimal) (database.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	sender, err := cfg.DB.GetUserByID(ctx, senderID)
	if err != nil {
		return database.Transaction{}, fmt.Errorf("could not get sender id: \n%v", err)
	}

	receiver, err := cfg.DB.GetUserByID(ctx, receiverID)
	if err != nil {
		return database.Transaction{}, fmt.Errorf("could not get receiver id: \n%v", err)
	}

	receiverCurrency, err := decimal.NewFromString(receiver.Currency)
	if err != nil {
		return database.Transaction{}, fmt.Errorf("could not get receiver currency id: \n%v", err)
	}

	senderCurrency, err := decimal.NewFromString(sender.Currency)
	if err != nil {
		return database.Transaction{}, fmt.Errorf("could not get sender currency id: \n%v", err)
	}

	uucSenderParams := database.UpdateUserCurrencyParams{
		Currency:  senderCurrency.Sub(quantity).String(),
		UpdatedAt: time.Now().UTC(),
		ID:        senderID,
	}
	_, err = cfg.DB.UpdateUserCurrency(ctx, uucSenderParams)
	if err != nil {
		return database.Transaction{}, fmt.Errorf("could not update sender currency: %v", err)
	}

	uucReceiverParams := database.UpdateUserCurrencyParams{
		Currency:  receiverCurrency.Add(quantity).String(),
		UpdatedAt: time.Now().UTC(),
		ID:        receiverID,
	}

	_, err = cfg.DB.UpdateUserCurrency(ctx, uucReceiverParams)
	if err != nil {
		if err = cfg.rollbackCurrencyUpdate(ctx, sender); err != nil {
			return database.Transaction{}, fmt.Errorf("cannot rollback: \n%v", err)
		}
		return database.Transaction{}, fmt.Errorf("could not update receiver currency: %v", err)
	}

	id := uuid.New()
	transactionParams := database.CreateTransactionParams{
		ID:         id,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     quantity.String(),
		ExecutedAt: time.Now().UTC(),
	}

	transaction, err := cfg.DB.CreateTransaction(ctx, transactionParams)
	if err != nil {
		if err = cfg.rollbackCurrencyUpdate(ctx, sender); err != nil {
			return database.Transaction{}, fmt.Errorf("could not create transaction currency and could not row back sender data: %v", err)
		}
		if err = cfg.rollbackCurrencyUpdate(ctx, receiver); err != nil {
			return database.Transaction{}, fmt.Errorf("could not create transaction currency and could not row back receiver data: %v", err)
		}
		return database.Transaction{}, fmt.Errorf("could not create transaction currency: %v", err)
	}
	return transaction, nil
}

func (cfg *apiConfig) rollbackCurrencyUpdate(ctx context.Context, user database.User) error {
	uucParams := database.UpdateUserCurrencyParams{
		Currency:  user.Currency,
		UpdatedAt: user.UpdatedAt,
		ID:        user.ID,
	}
	_, err := cfg.DB.UpdateUserCurrency(ctx, uucParams)
	if err != nil {
		return fmt.Errorf("cannot rollback: \n %v", err)
	}
	return nil
}
