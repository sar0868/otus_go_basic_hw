-- name: UserAdd :execresult
insert into shop.Users(name, email, "password") 
values ($1, $2, $3);

-- name: ProductCreate :execresult
insert into shop.Products(name, price)
values ($1, $2);


-- name: OrdersCreate :execresult
insert into shop.Orders(user_id)
values ($1);

-- name: OrderProductsCreate :execresult
insert into shop.orderproducts 
(order_id, product_id, quantity)
values ($1, $2, $3);

-- name: OrderCreate :exec
insert into shop.orderproducts 
(order_id, product_id, quantity)
values ($1, $2, $3);

-- name: OrderUpdate :exec
update shop.orders ord
set total_amount=(
select sum(p.price * op.quantity) from 
shop.orderproducts op
inner join shop.products p on op.product_id=p.id
where op.order_id = $1
group by op.order_id)
where ord.id  = $1;

-- name: OrderCreateByUser :exec
insert into shop.Orders (user_id)
values ((select id from shop.Users where name = $1));

-- name: OrderProductCreateByUserAndProductQuanti :exec
insert into shop.Orderproducts 
(order_id, product_id, quantity)
values 
(
    (select id from shop.orders 
    where user_id = (select id from shop.Users u where u.name = $1)),
    (select id from shop.products p where p.name = $2),
    $3
);

-- name: OrderUpdateTotalAmountByOrderID :exec
update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = $1
	group by op.order_id)
where ord.id = $1;

-- name: UserUpdate :exec
update shop.users 
set name = $1
where "name" = $2;

-- name: ProductUpdate :exec
update shop.products
set price = $1
where name = $2;

-- name: UserDelete :exec
DELETE from shop.Users
where name=$1;

-- name: ProductDelete :exec
delete from shop.products 
where name = $1;

-- name: OrderDelete :exec
delete from shop.orders 
where id=$1;

-- name: Users :many
select * from shop.users u ;

-- name: Products :many
select * from shop.products p ;

-- name: OrdersProducts :many
select * from shop.orderproducts op ;

-- name: Orders :many
select  * from shop.orders o ;

-- name: UserGetByName :one
select * from shop.users u 
where name = $1;

-- name: ProductGetRangePrice :many
select * from shop.products p 
where price between $1 and $2
order by price ;

-- name: GetUserOrdersByName :many
select u.name as "user", o.id as order_id, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id 
where u.name like $1;

-- name: UsersSumTotalOrdersAvrPrice :many
select u.name as "user", sum(o.total_amount) as total_orders, avg(p.price) as "avr price" 
from shop.orders o  
inner join shop.orderproducts op on o.id = op.order_id 
inner join shop.products p on op.product_id = p.id 
right join shop.Users u on o.user_id = u.id 
group by u.name;
