-- name: CreateTransaction :one
INSERT INTO transactions (id, sender_id, receiver_id, amount, executed_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;