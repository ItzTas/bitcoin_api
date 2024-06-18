-- name: CreateTransaction :one
INSERT INTO transactions (id, sender_id, receiver_id, amount, executed_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserTransactions :many
SELECT *, 
    CASE 
        WHEN sender_id = $1 THEN 'sender'
        ELSE 'receiver'
    END AS user_role
FROM transactions
WHERE sender_id = $1 OR receiver_id = $1;

-- name: GetUserTransactionsWithLimit :many
SELECT *, 
    CASE 
        WHEN sender_id = $1 THEN 'sender'
        ELSE 'receiver'
    END AS user_role
FROM transactions
WHERE sender_id = $1 OR receiver_id = $1
LIMIT $2;