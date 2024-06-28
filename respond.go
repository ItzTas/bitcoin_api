package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJSON[T any](w http.ResponseWriter, status int, payload T) {
	w.Header().Set("Content-type", "application/json")
	resp, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	if status >= 500 {
		fmt.Printf("\nResponding with 5XX status code: %v\n", status)
	}
	type errorMessage struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, errorMessage{
		Error: message,
	})
}

func respondWithNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
