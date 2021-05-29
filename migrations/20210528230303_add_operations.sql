-- +goose Up
-- +goose StatementBegin
INSERT INTO t_operation (operation_type_id, description, movement_type) 
VALUES 
(1, 'COMPRA A VISTA', 'S'),
(2, 'COMPRA PARCELADA', 'S'),
(3, 'SAQUE', 'S'),
(4, 'PAGAMENTO', 'E')

-- +goose StatementEnd
