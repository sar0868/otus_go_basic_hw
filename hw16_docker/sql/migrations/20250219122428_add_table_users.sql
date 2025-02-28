-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shop.Users(
	id serial primary key,
	name varchar(120) not null unique,
	email varchar(120),
	password varchar(120) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shop.Users;
-- +goose StatementEnd
