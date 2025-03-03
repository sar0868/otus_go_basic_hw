-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS shop;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA if EXISTS shop

-- +goose StatementEnd
