-- name: CreateDeposit :one 
INSERT INTO deposits (id, wallet_id, amount, executed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDepositsByUserID :many
SELECT * FROM deposits
WHERE wallet_id IN (
    SELECT id
    FROM wallets
    WHERE owner_id = $1
);

-- name: GetDepositsByUserIDWithLimit :many
SELECT * FROM deposits
WHERE wallet_id IN (
    SELECT id
    FROM wallets
    WHERE owner_id = $1
)
LIMIT $2;