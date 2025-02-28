-- name: ProductCreate :one
insert into shop.Products(name, price)
values ($1, $2) returning id;

-- name: ProductUpdate :exec
update shop.products
set price = $1
where name = $2;


-- name: ProductDelete :exec
delete from shop.products 
where name = $1;


-- name: Products :many
select * from shop.products p ;

-- name: ProductGetRangePrice :many
select * from shop.products p 
where price between $1 and $2
order by price ;

-- name: ProductGetByName :one
select * from shop.products p 
where p.name = $1;

-- name: ProductGetById :one
select * from shop.products p 
where p.id = $1;

-- name: ProductID :one
select id from shop.products p
where p.name = $1;
