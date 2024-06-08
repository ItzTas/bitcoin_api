package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func respondWithJSON[T any](w http.ResponseWriter, status int, payload T) {
// 	resp, err := json.Marshal(payload)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(status)
// 	_, err = w.Write(resp)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// func respondWithError(w http.ResponseWriter)
