-- +goose Up
CREATE TABLE IF NOT EXISTS t_accounts (
    account_id INT NOT NULL,
    document_number VARCHAR(30) NOT NULL,
    PRIMARY KEY(account_id)
);

-- +goose Down
DROP TABLE t_accounts;