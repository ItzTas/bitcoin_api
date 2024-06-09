package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	const readTimeout = 5 * time.Minute
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/error", errorTest)

	server := http.Server{
		Handler:     mux,
		ReadTimeout: readTimeout,
		Addr:        ":" + port,
	}

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
