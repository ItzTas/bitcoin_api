-- name: CreateWallet :one
INSERT INTO wallets (id, owner_id, crypto_type_id, balance_usd, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;