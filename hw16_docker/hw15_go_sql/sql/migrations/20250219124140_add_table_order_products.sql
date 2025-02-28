-- +goose Up
-- +goose StatementBegin
create table if not exists shop.OrderProducts(
	order_id integer references shop.Orders(id) on delete cascade,
	product_id integer references shop.Products(id) on delete cascade,
	quantity numeric default 0,
	primary key(order_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shop.OrderProducts;
-- +goose StatementEnd
