package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/client"
	"github.com/ItzTas/bitcoinAPI/internal/database"
)

var (
	ErrorNonExisting = errors.New("coin doesn't exist")
)

func (cfg *apiConfig) cryptoSaver(limit *int, interval time.Duration, addicionalCoins ...client.Coin) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		coins, err := cfg.client.GetCoinList(limit)
		if err != nil {
			fmt.Printf("\nCould not get coins list:\n %v\n\n", err)
			return
		}

		if addicionalCoins != nil {
			coins = append(coins, addicionalCoins...)
		}

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(coins []client.Coin) {
			defer wg.Done()
			for _, coin := range coins {
				err := cfg.updateCryptoData(coin)
				if err != nil {
					if errors.Is(err, ErrorNonExisting) {
						continue
					}
					fmt.Printf("\nerror updating crypto:\n %v\n\n", err)
					continue
				}
			}
		}(coins)
		wg.Wait()
		fmt.Println("Cryptos Saved!")
	}
}

func (cfg *apiConfig) updateCryptoData(coin client.Coin) error {
	coinData, err := cfg.client.GetCoinData(coin.ID)
	if coinData.ID == "" {
		return ErrorNonExisting
	}
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	if _, err := cfg.DB.GetCryptoByID(ctx, coin.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cfg.createCrypto(coinData)
		}
		return err
	}

	ucParams := database.UpdateCryptoParams{
		CurrentPriceUsd: priceToString(coinData.MarketData.CurrentPrice.Usd),
		CurrentPriceEur: priceToString(coinData.MarketData.CurrentPrice.Eur),
		DescriptionEn:   coinData.Description.En,
		UpdatedAt:       time.Now().UTC(),
		ID:              coinData.ID,
	}

	_, err = cfg.DB.UpdateCrypto(ctx, ucParams)
	if err != nil {
		return fmt.Errorf("could not update crypto: \n%v \n%v", ucParams, err)
	}
	return nil
}

func (cfg *apiConfig) createCrypto(coinData client.CoinData) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeOut)
	defer cancel()

	ccParams := database.CreateCryptoParams{
		CurrentPriceUsd: priceToString(coinData.MarketData.CurrentPrice.Usd),
		CurrentPriceEur: priceToString(coinData.MarketData.CurrentPrice.Eur),
		DescriptionEn:   coinData.Description.En,
		UpdatedAt:       time.Now().UTC(),
		ID:              coinData.ID,
		Symbol:          coinData.Symbol,
		Name:            coinData.Name,
	}
	_, err := cfg.DB.CreateCrypto(ctx, ccParams)
	if err != nil {
		errorParams := database.CreateCryptoParams{
			CurrentPriceUsd: priceToString(coinData.MarketData.CurrentPrice.Usd),
			CurrentPriceEur: priceToString(coinData.MarketData.CurrentPrice.Eur),
			UpdatedAt:       time.Now().UTC(),
			ID:              coinData.ID,
			Symbol:          coinData.Symbol,
			Name:            coinData.Name,
		}
		return fmt.Errorf("could not create crypto: \n%v \n%v", errorParams, err)
	}
	return nil
}

func priceToString(price float64) string {
	return strconv.FormatFloat(price, 'f', -1, 64)
}
