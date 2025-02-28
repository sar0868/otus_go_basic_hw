-- name: UserAdd :one
insert into shop.Users(name, email, "password") 
values ($1, $2, $3) returning id;


-- name: UserUpdate :exec
update shop.users 
set name = $2
where "name" = $1;

-- name: UserDelete :exec
DELETE from shop.Users
where name=$1;


-- name: Users :many
select * from shop.users u ;

-- name: UserGetByName :one
select * from shop.users u 
where name = $1;

-- name: GetUserOrdersByName :many
select u.name as "user", o.id as order_id, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id 
where u.name like $1;

-- name: UsersStatistic :many
select u.name as "user", sum(o.total_amount) as total_orders, avg(p.price) as "avr price" 
from shop.orders o  
inner join shop.orderproducts op on o.id = op.order_id 
inner join shop.products p on op.product_id = p.id 
right join shop.Users u on o.user_id = u.id 
group by u.name;

-- name: UserByID :one
select * from shop.Users u
where u.id = $1;
