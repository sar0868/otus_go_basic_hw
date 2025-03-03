-- +goose Up
-- +goose StatementBegin
create table if not exists shop.Products(
	id serial primary key,
	name text not null unique,
	price numeric not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shop.Products;
-- +goose StatementEnd
