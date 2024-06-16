-- +goose Up
CREATE TABLE cryptocurrencies (
    id TEXT PRIMARY KEY,
    symbol TEXT UNIQUE,
    name TEXT UNIQUE,
    current_price_usd NUMERIC,
    current_price_eur NUMERIC,
    description_en TEXT
);

-- +goose Down
DROP TABLE cryptocurrencies;