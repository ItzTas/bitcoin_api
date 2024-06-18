-- +goose Up
ALTER TABLE transactions
ADD COLUMN is_between_users BOOLEAN NOT NULL DEFAULT TRUE;

-- +goose Down 
ALTER TABLE transactions 
DROP COLUMN is_between_users;