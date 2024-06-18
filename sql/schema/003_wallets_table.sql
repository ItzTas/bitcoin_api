-- +goose Up
CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL,
    crypto_type_id TEXT NOT NULL,
    balance_usd NUMERIC NOT NULL, 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_crypto_type_id
        FOREIGN KEY (crypto_type_id)
        REFERENCES cryptocurrencies(id),
    CONSTRAINT unique_owner_id_crypto_type_id
        UNIQUE(owner_id, crypto_type_id)
);

-- +goose Down
DROP TABLE wallets;