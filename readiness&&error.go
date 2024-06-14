package main

import "net/http"

func readiness(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, response{
		Status: "ok",
	})
}

func errorTest(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "error")
}
