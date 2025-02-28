-- name: OrderDelete :exec
delete from shop.orders 
where id=$1;

-- name: OrderCreateByUser :one
insert into shop.Orders (user_id)
values ((select id from shop.Users where name = $1))
returning id;

-- name: OrderUpdateTotal :exec
update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = $1
	group by op.order_id)
where ord.id = $1;

-- name: Orders :many
select  o.id, u.name, o.order_date, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id
order by o.id, u.name;

-- name: OrderUser :one
select  o.id, u.name, o.order_date, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id
where o.id = $1;

-- name: OrderByID :one
select * from shop.orders o
where o.id = $1;
