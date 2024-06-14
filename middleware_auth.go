package main

import (
	"net/http"

	"github.com/ItzTas/bitcoinAPI/internal/database"
)

type handlerSecure = func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler handlerSecure) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
