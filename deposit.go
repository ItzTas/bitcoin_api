package main

import (
	"encoding/json"
	"net/http"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) handlerSendToAccount(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type paramaters struct {
		SendQuantity string `json:"send_quantity"`
	}

	recID := r.PathValue("receiver_id")
	if recID == "" {
		respondWithError(w, http.StatusBadRequest, "no receiver id")
		return
	}

	params := paramaters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode paramethers")
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

	if compVal := currency.Cmp(sendQuant); compVal < 0 {
		respondWithError(w, http.StatusBadRequest, "Send quantity bigger than currency")
		return
	}

}
