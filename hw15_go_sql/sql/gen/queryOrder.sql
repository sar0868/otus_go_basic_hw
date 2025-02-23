-- name: OrdersCreate :execresult
insert into shop.Orders(user_id)
values ($1);

-- name: OrderUpdate :exec
update shop.orders ord
set total_amount=(
select sum(p.price * op.quantity) from 
shop.orderproducts op
inner join shop.products p on op.product_id=p.id
where op.order_id = $1
group by op.order_id)
where ord.id  = $1;

-- name: OrderDelete :exec
delete from shop.orders 
where id=$1;

-- name: OrderCreateByUser :exec
insert into shop.Orders (user_id)
values ((select id from shop.Users where name = $1));

-- name: OrderUpdateTotalAmountByOrderID :exec
update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = $1
	group by op.order_id)
where ord.id = $1;

-- name: Orders :many
select  * from shop.orders o ;
