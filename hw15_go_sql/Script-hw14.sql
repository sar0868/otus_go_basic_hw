create schema if not exists shop;

--=======================================
--create tables

create table if not exists shop.Users(
	id serial primary key,
	name varchar(120) not null unique,
	email varchar(120),
	password varchar(120) not null
);

create table if not exists shop.Orders(
	id serial primary key,
	user_id integer references shop.Users(id) on delete cascade,
	order_date timestamptz not null default now(),
	total_amount numeric not null default 0
);

create table if not exists shop.Products(
	id serial primary key,
	name text not null unique,
	price numeric not null default 0
);

create table if not exists shop.OrderProducts(
	order_id integer references shop.Orders(id) on delete cascade,
	product_id integer references shop.Products(id) on delete cascade,
	quantity numeric not null default 0,
	primary key(order_id, product_id)
);

--alter table shop.OrderProducts
--add quantity numeric not null default 0; 

--alter table shop.orders 
--alter column total_amount set default 0;

--================================
-- insert data

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


insert into shop.Orders (user_id)
values 
((select id from shop.Users where name like 'qwe')
);

insert into shop.Orderproducts 
(order_id, product_id, quantity)
values 
(
	(select id from shop.orders 
	where user_id = (select id from shop.Users where name like 'qwe')),
	(select id from shop.products
where name like 'milk'), 2),
(
	(select id from shop.orders 
	where user_id = (select id from shop.Users where name like 'qwe')),
	(select id from shop.products
	where name like 'bread'), 2);

update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = 3
	group by op.order_id)
where ord.id = 3;


insert into shop.orderproducts 
(order_id, product_id, quantity)
values
(2, 2, 2),
(2, 1, 0.3);

insert into shop.Orders (user_id)
values 
((select id from shop.Users where name like 'admi%')),
((select id from shop.Users where name like 'admi%'));

insert into shop.Orderproducts 
(order_id, product_id, quantity)
values 
(4, (select id from shop.products where name like 'potat%'), 4.76),
(4, (select id from shop.products where name like 'butter'), 0.8);

insert into shop.Orderproducts 
(order_id, product_id, quantity)
values 
(5, (select id from shop.products where name like 'milk'), 1);

update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = 4
	group by op.order_id)
where ord.id = 4;

update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = 5
	group by op.order_id)
where ord.id = 5;

--========================
--update data (users, products)

update shop.users 
set name = 'user'
where "name" like 'as%';

update shop.products
set price = 102.2
where name like 'mil%';

--========================================
--delete data (users, products, orders)

DELETE from shop.Users
where name='qwe';

delete from shop.products 
where name like 'milk';

update shop.orders ord
set total_amount=(
select sum(p.price * op.quantity) from 
shop.orderproducts op
inner join shop.products p on op.product_id=p.id
where op.order_id = 1
group by op.order_id);

delete from shop.orders 
where id=2;

--========================================
--Напишите запрос на выборку пользователей и выборку товаров
--Напишите запрос на выборку заказов по пользователю
--Напишите запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)

select * from shop.users u ;
select * from shop.products p ;
select * from shop.orderproducts op ;
select  * from shop.orders o ;

select * from shop.users u 
where name like 'qwe';

select name, price from shop.products p 
where price between 50 and 120
order by price ;

select u.name as "user", o.id as order_id, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id 
where u.name like 'admi%';

select u.name as "user", sum(o.total_amount) as total_orders, avg(p.price) as "avr price" 
from shop.orders o  
inner join shop.orderproducts op on o.id = op.order_id 
inner join shop.products p on op.product_id = p.id 
right join shop.Users u on o.user_id = u.id 
group by u.name;

--=================================================
--create index

create index if not exists ix_users_id on shop.Users (id);
create index if not exists ix_users_name on shop.Users (name);
create index if not exists ix_orders_id on shop.Orders (id);
create index if not exists ix_orders_order_date on shop.Orders (order_date);
create index if not exists ix_product_id on shop.Products (id);
create index if not exists ix_product_name on shop.Products (name);
create index if not exists ix_orders_product on shop.OrderProducts(order_id, product_id);
