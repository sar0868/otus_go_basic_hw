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

-- name: OrdersProductsFull :many
select op.order_id, p.name, op.quantity from shop.orderproducts op 
inner join shop.products p on op.product_id = p.id;

-- name: OrderProductUpdate :exec
update shop.orderproducts op
set quantity = $3
where op.order_id = $1 and op.product_id = $2;
