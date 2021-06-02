-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_transactions (
    transaction_id SERIAL,
    account_id INT NOT NULL,
    operation_type_id INT NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    balance NUMERIC(15,2) NOT NULL,
    event_date TIMESTAMP NOT NULL,
    PRIMARY KEY(transaction_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_transactions;
-- +goose StatementEnd
