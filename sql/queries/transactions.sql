-- name: CreateTransaction :one
INSERT INTO transactions (id, sender_id, receiver_id, amount, executed_at, is_between_users)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;