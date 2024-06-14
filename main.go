package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	const readTimeout = 5 * time.Minute
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading the port")
		return
	}

	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		fmt.Println("Could not stablish connection to the database")
		return
	}

	dbQueries := database.New(db)

	cfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/error", errorTest)

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users/{user_id}", cfg.handlerGetUserByID)

	server := http.Server{
		Handler:     mux,
		ReadTimeout: readTimeout,
		Addr:        ":" + port,
	}

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
