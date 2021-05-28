-- +goose Up
CREATE TABLE IF NOT EXISTS t_operation (
    operation_type_id INT NOT NULL,
    description VARCHAR(30) NOT NULL,
    movement_type VARCHAR(2) NOT NULL,
    PRIMARY KEY(operation_type_id)
);

-- +goose Down
DROP TABLE t_operation;