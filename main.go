package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/client"
	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB        *database.Queries
	client    *client.Client
	jwtSecret string
}

var (
	limit = 10
)

const (
	DefaultContextTimeOut = 1 * time.Minute
	DefaultClientTImeOut  = 5 * time.Minute
	DefautSaverInterval   = 120 * time.Minute
)

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
	geckoKey := os.Getenv("COIN_GECKO_KEY")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		fmt.Println("Could not stablish connection to the database")
		return
	}

	dbQueries := database.New(db)

	cfg := apiConfig{
		DB:        dbQueries,
		jwtSecret: jwtSecret,
		client:    client.NewClient(DefaultClientTImeOut, geckoKey),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/error", errorTest)
	mux.HandleFunc("GET /v1/secure_endpoint", cfg.middlewareAuth(secureEndpoint))

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users/{user_id}", cfg.handlerGetUserByID)
	mux.HandleFunc("GET /v1/users", cfg.handlerGetUsers) // supports limit query (defaults to 20)
	mux.HandleFunc("PUT /v1/users", cfg.middlewareAuth(cfg.handlerUpdateUser))

	mux.HandleFunc("POST /v1/login", cfg.handlerLogin)

	mux.HandleFunc("POST /v1/wallets", cfg.middlewareAuth(cfg.handlerCreateWallet))

	server := http.Server{
		Handler:     mux,
		ReadTimeout: readTimeout,
		Addr:        ":" + port,
	}

	go cfg.cryptoSaver(&limit, DefautSaverInterval,
		client.Coin{
			ID:     "bitcoin",
			Symbol: "btc",
			Name:   "Bitcoin",
		},
		client.Coin{
			ID:     "ethereum",
			Symbol: "eth",
			Name:   "Ethereum",
		},
	)

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
