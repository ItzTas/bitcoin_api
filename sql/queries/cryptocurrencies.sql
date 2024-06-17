-- name: CreateCrypto :one
INSERT INTO cryptocurrencies (id, symbol, name, current_price_usd, current_price_eur, description_en, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCrypto :one
UPDATE cryptocurrencies
SET current_price_usd = $1, current_price_eur = $2, description_en = $3, updated_at = $4
WHERE id = $5
RETURNING *;

-- name: GetCryptoByID :one
SELECT * FROM cryptocurrencies
WHERE id = $1;

-- name: GetCryptocurrencies :many
SELECT * FROM cryptocurrencies;

-- name: GetCryptocurrencyByID :one
SELECT * FROM cryptocurrencies
WHERE id = $1;