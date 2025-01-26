create schema if not exists shop;

create table shop.Users(
	id serial primary key,
	name varchar(120) not null,
	email varchar(120),
	password varchar(120) not null
);

create table shop.Orders(
	id serial primary key,
	user_id integer references shop.Users(id) on delete cascade,
	order_date timestamptz not null default now(),
	total_amount numeric not null
);

create table shop.Products(
	id serial primary key,
	name text not null,
	price numeric not null default 0
);

create table shop.OrderProducts(
	order_id integer references shop.Orders(id) on delete cascade,
	product_id integer references shop.Products(id) on delete cascade,
	primary key(order_id, product_id)
);

alter table shop.OrderProducts
add quantity numeric not null default 0; 

alter table shop.orders 
alter column total_amount set default 0;

alter table shop.orders 
alter column 


================================

insert into shop.Users(name, email, "password") 
values
('admin', 'admin@example.com', 'admin'),
('qwe', '', 'qwe1'),
('asd','asd@mail.com', 'asd1');

insert into shop.Products(name, price) values
('milk', 92.5),
('bread', 35.5),
('butter', 234.90),
('potatoes', 54.60);



insert into shop.Orders(user_id)
values
(1),
(2);

insert into shop.orderproducts 
(order_id, product_id, quantity)
values
(1, 1, 1),
(1, 2, 1),
(1, 4, 0.75);

insert into shop.orderproducts 
(order_id, product_id, quantity)
values
(2, 2, 2),
(2, 1, 0.3);

update shop.orders ord
set total_amount=(
select sum(p.price * op.quantity) from 
shop.orderproducts op
inner join shop.products p on op.product_id=p.id
where op.order_id = 1
group by op.order_id)
where ord.id  = 1;

update shop.orders ord
set total_amount=(
select sum(p.price * op.quantity) from 
shop.orderproducts op
inner join shop.products p on op.product_id=p.id
where op.order_id = 2
group by op.order_id)
where ord.id  = 2;


select * from shop.users u ;
select * from shop.products p ;
select * from shop.orderproducts op ;
select  * from shop.orders o ;

