-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_operation (
    operation_type_id SERIAL,
    description VARCHAR(30) NOT NULL,
    movement_type VARCHAR(2) NOT NULL,
    PRIMARY KEY(operation_type_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_operation;
-- +goose StatementEnd
