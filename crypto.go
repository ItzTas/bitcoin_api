package main

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/client"
	"github.com/ItzTas/bitcoinAPI/internal/database"
)

func (cfg *apiConfig) updateCryptoData(coin client.Coin) error {
	coinData, err := cfg.client.GetCoinData(coin.ID)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	if _, err := cfg.DB.GetCryptoByID(ctx, coin.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cfg.createCrypto(coinData)
		} else {
			return err
		}
	}

	currenPriceUsd := strconv.FormatFloat(coinData.MarketData.CurrentPrice.Usd, 'f', -1, 64)
	currenPriceEur := strconv.FormatFloat(coinData.MarketData.CurrentPrice.Eur, 'f', -1, 64)

	ucParams := database.UpdateCryptoParams{
		CurrentPriceUsd: currenPriceUsd,
		CurrentPriceEur: currenPriceEur,
		DescriptionEn:   coinData.Description.En,
		UpdatedAt:       time.Now().UTC(),
		ID:              coinData.ID,
	}

	_, err = cfg.DB.UpdateCrypto(ctx, ucParams)
	return err
}

func (cfg *apiConfig) createCrypto(coinData *client.CoinData) error {
	if coinData.ID == "" {
		return errors.New("coin doesn't exist")
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	currenPriceUsd := strconv.FormatFloat(coinData.MarketData.CurrentPrice.Usd, 'f', -1, 64)
	currenPriceEur := strconv.FormatFloat(coinData.MarketData.CurrentPrice.Eur, 'f', -1, 64)
	ccParams := database.CreateCryptoParams{
		CurrentPriceUsd: currenPriceUsd,
		CurrentPriceEur: currenPriceEur,
		DescriptionEn:   coinData.Description.En,
		UpdatedAt:       time.Now().UTC(),
		ID:              coinData.ID,
		Symbol:          coinData.Symbol,
		Name:            coinData.Name,
	}
	_, err := cfg.DB.CreateCrypto(ctx, ccParams)
	return err
}
