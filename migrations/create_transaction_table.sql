-- +goose Up
CREATE TABLE IF NOT EXISTS t_transactions (
    transaction_id INT NOT NULL,
    account_id INT NOT NULL,
    operation_type_id INT NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    event_date TIMESTAMP NOT NULL,
    PRIMARY KEY(transaction_id)
);

-- +goose Down
DROP TABLE t_transactions;