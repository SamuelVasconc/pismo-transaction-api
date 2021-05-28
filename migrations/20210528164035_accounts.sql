-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_accounts (
    account_id SERIAL,
    document_number VARCHAR(30) NOT NULL,
    PRIMARY KEY(account_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_accounts;
-- +goose StatementEnd
