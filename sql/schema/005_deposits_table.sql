-- +goose Up
CREATE TABLE deposits (
    id UUID PRIMARY KEY,
    wallet_id UUID NOT NULL,
    amount NUMERIC NOT NULL,
    executed_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_wallet_id
        FOREIGN KEY (wallet_id)
        REFERENCES wallets(id)
);

-- +goose Down 
DROP TABLE deposits;