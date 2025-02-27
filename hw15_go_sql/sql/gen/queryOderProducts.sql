-- name: OrderProductsCreate :exec
insert into shop.orderproducts 
(order_id, product_id, quantity)
values ($1, $2, $3);

-- name: ProductAddOrder :exec
insert into shop.Orderproducts 
(order_id, product_id, quantity)
values 
( 
    $1, 
    (select id from shop.products p where p.name = $2),
    $3
);

-- name: OrdersProducts :many
select * from shop.orderproducts op ;
