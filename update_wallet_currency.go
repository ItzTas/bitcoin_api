package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) handlerUpdateWalletCurrency(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type paramaters struct {
		Value string `json:"value"`
	}
	cryptoType := r.PathValue("coin_id")
	if cryptoType == "" {
		respondWithError(w, http.StatusBadRequest, "no crypto id provided")
		return
	}

	params := paramaters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode params")
		return
	}

	v, err := decimal.NewFromString(params.Value)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse value")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	gubtwParams := database.GetUserTypeWalletParams{
		OwnerID:      dbUser.ID,
		CryptoTypeID: cryptoType,
	}
	wallet, err := cfg.DB.GetUserTypeWallet(ctx, gubtwParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Wallet with given coin id does not exist")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "could not get wallet")
		return
	}

	walletVal, err := decimal.NewFromString(wallet.BalanceUsd)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse wallet balance")
		return
	}

	if v.LessThan(decimal.Zero) && v.Add(walletVal).LessThan(decimal.Zero) {
		respondWithError(w, http.StatusBadRequest, "The value cannot be greater than the wallet currency")
		return
	}

	userCurrency, err := decimal.NewFromString(dbUser.Currency)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse user currency")
		return
	}

	if v.GreaterThan(userCurrency) {
		fmt.Println(v)
		respondWithError(w, http.StatusBadRequest, "value cannot be grater than the user currency")
		return
	}

	deposit, err := cfg.makeDeposit(dbUser, wallet, v)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create transaction: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseDepositToDeposit(deposit))
}

func (cfg *apiConfig) makeDeposit(user database.User, wallet database.Wallet, quantity decimal.Decimal) (database.Deposit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()
	id := uuid.New()

	userCurrency, err := decimal.NewFromString(user.Currency)
	if err != nil {
		return database.Deposit{}, err
	}

	walletCurrency, err := decimal.NewFromString(wallet.BalanceUsd)
	if err != nil {
		return database.Deposit{}, err
	}

	uucParams := database.UpdateUserCurrencyParams{
		Currency:  userCurrency.Sub(quantity).String(),
		UpdatedAt: time.Now().UTC(),
		ID:        user.ID,
	}
	_, err = cfg.DB.UpdateUserCurrency(ctx, uucParams)
	if err != nil {
		return database.Deposit{}, err
	}

	uwParams := database.UpdateWalletParams{
		BalanceUsd: walletCurrency.Add(quantity).String(),
		UpdatedAt:  time.Now().UTC(),
		ID:         wallet.ID,
	}

	_, err = cfg.DB.UpdateWallet(ctx, uwParams)

	if err != nil {
		err := cfg.rollbackCurrencyUpdate(ctx, user)
		if err != nil {
			return database.Deposit{}, fmt.Errorf("could not roll user back: \n%v", err)
		}
		return database.Deposit{}, err
	}

	cdParams := database.CreateDepositParams{
		ID:         id,
		WalletID:   wallet.ID,
		Amount:     quantity.String(),
		ExecutedAt: time.Now().UTC(),
	}

	deposit, err := cfg.DB.CreateDeposit(ctx, cdParams)
	if err != nil {
		if err := cfg.rollbackCurrencyUpdate(ctx, user); err != nil {
			return database.Deposit{}, fmt.Errorf("could not roll user back: \n%v", err)
		}

		if err := cfg.rollbackCurrencyWalletupdate(ctx, wallet); err != nil {
			return database.Deposit{}, fmt.Errorf("could not roll wallet back: \n%v", err)
		}
	}
	return deposit, err
}

func (cfg *apiConfig) rollbackCurrencyWalletupdate(ctx context.Context, wallet database.Wallet) error {
	uwParams := database.UpdateWalletParams{
		BalanceUsd: wallet.BalanceUsd,
		UpdatedAt:  wallet.UpdatedAt,
		ID:         wallet.ID,
	}
	_, err := cfg.DB.UpdateWallet(ctx, uwParams)
	if err != nil {
		return fmt.Errorf("cannot rollback: \n %v", err)
	}
	return nil
}
