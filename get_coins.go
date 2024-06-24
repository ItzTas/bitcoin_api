package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/ItzTas/coinerAPI/internal/database"
	"github.com/thoas/go-funk"
)

type Cryptocurrency struct {
	ID              string    `json:"id"`
	Symbol          string    `json:"symbol"`
	Name            string    `json:"name"`
	CurrentPriceUsd string    `json:"current_price_usd"`
	CurrentPriceEur string    `json:"current_price_eur"`
	DescriptionEn   string    `json:"description_en"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func databaseCryptocurrencyToCryptocurrency(coin database.Cryptocurrency) Cryptocurrency {
	return Cryptocurrency(coin)
}

func (cfg *apiConfig) handlerRetriveCoins(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	dbcoins, err := cfg.DB.GetCryptocurrencies(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get coins")
		return
	}

	coins := funk.Map(dbcoins, func(dbcoin database.Cryptocurrency) Cryptocurrency {
		return databaseCryptocurrencyToCryptocurrency(dbcoin)
	}).([]Cryptocurrency)

	respondWithJSON(w, http.StatusOK, coins)
}

func (cfg *apiConfig) handlerRetriveCoinByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	coinID := r.PathValue("coin_id")
	coin, err := cfg.DB.GetCryptocurrencyByID(ctx, coinID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Coin does not exist")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not get coin")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseCryptocurrencyToCryptocurrency(coin))
}
