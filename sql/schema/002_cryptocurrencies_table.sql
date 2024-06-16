-- +goose Up
CREATE TABLE cryptocurrencies (
    id TEXT PRIMARY KEY,
    symbol TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    current_price_usd NUMERIC NOT NULL,
    current_price_eur NUMERIC NOT NULL,
    description_en TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE cryptocurrencies;