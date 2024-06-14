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
	DB        *database.Queries
	jwtSecret string
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
	jwtSecret := os.Getenv("JWT_SECRET")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		fmt.Println("Could not stablish connection to the database")
		return
	}

	dbQueries := database.New(db)

	cfg := apiConfig{
		DB:        dbQueries,
		jwtSecret: jwtSecret,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/error", errorTest)

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users/{user_id}", cfg.handlerGetUserByID)
	mux.HandleFunc("GET /v1/users", cfg.handlerGetUsers) // supports limit query (defaults to 20)
	// mux.HandleFunc("PUT /v1/users", )

	mux.HandleFunc("POST /v1/login", cfg.handlerLogin)

	server := http.Server{
		Handler:     mux,
		ReadTimeout: readTimeout,
		Addr:        ":" + port,
	}

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
