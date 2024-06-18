-- name: CreateWallet :one
INSERT INTO wallets (id, owner_id, crypto_type_id, balance_usd, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserTypeWallet :one
SELECT * FROM wallets
WHERE owner_id = $1 AND crypto_type_id = $2;

-- name: UpdateWallet :one
UPDATE wallets
SET balance_usd = $1, updated_at = $2
WHERE id = $3
RETURNING *;