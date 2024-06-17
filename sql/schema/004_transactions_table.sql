-- +goose Up
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    amount NUMERIC NOT NULL,
    executed_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_sender_id
        FOREIGN KEY (sender_id)
        REFERENCES users(id),
    CONSTRAINT fk_receiver_id
        FOREIGN KEY (receiver_id)
        REFERENCES users(id)
);

-- +goose Down
DROP TABLE transactions;