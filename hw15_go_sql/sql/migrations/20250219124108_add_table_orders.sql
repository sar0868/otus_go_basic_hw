-- +goose Up
-- +goose StatementBegin
create table if not exists shop.Orders(
	id serial primary key,
	user_id integer references shop.Users(id) on delete cascade,
	order_date timestamptz not null default now(),
	total_amount numeric not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shop.Orders;
-- +goose StatementEnd
