-- name: CreateDeposit :one 
INSERT INTO deposits (id, wallet_id, amount, executed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;