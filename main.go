package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	const readTimeout = 5 * time.Minute
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading the port")
		return
	}

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
