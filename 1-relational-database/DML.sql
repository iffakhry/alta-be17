-- DML
-- INSERT

-- insert ke table users
INSERT INTO users(name, email)
VALUES ("Adi", "adi@mail.com");

INSERT INTO users(name, email)
VALUES ("Budi", "budi@mail.com");

INSERT INTO users(name, email)
VALUES ("Budi2", "budi2@mail.com");

INSERT INTO users(id, name, email)
VALUES (1, "Budi", "budi2@mail.com");

INSERT INTO users(name, email)
VALUES ("Cindi", "cindi@mail.com");

-- insert ke table products
INSERT INTO products(user_id, product_name, stock, price, jenis)
VALUES (1, "Wafer Kong Guan", 10, 50000, "makanan");

INSERT INTO products(user_id, product_name, stock, price, jenis)
VALUES (1, "Wafer Nissin", 10, 40000, "makanan");

INSERT INTO products(user_id, product_name, stock, price, jenis)
VALUES (2, "Wafer Kong Guan Cokelat", 10, 55000, "makanan");

INSERT INTO products(user_id, product_name, stock, price, jenis)
VALUES (2, "Pop Ice", 10, 5000, "minuman");

-- insert ke table transaction
INSERT INTO transactions(user_id, total_price)
VALUES (2, 100000);

INSERT INTO transactions(user_id, total_price)
VALUES (5, 105000);

-- insert ke table transaction detail
INSERT INTO transaction_details(transaction_id, product_id, quantity, total_price)
VALUES (1, 1, 2, 100000);

INSERT INTO transaction_details(transaction_id, product_id, quantity, total_price)
VALUES (3, 3, 1, 55000);

INSERT INTO transaction_details(transaction_id, product_id, quantity, total_price)
VALUES (2, 1, 2, 100000);
INSERT INTO transaction_details(transaction_id, product_id, quantity, total_price)
VALUES (2, 5, 1, 5000);

-- SELECT
SELECT * FROM users;
SELECT * FROM products;
SELECT * FROM transactions;
SELECT * FROM transaction_details;

SELECT id, email FROM users;

SELECT * FROM users where id = 2;
SELECT * FROM users where name = "budi";

-- UPDATE

UPDATE users SET
name = "Budi21"
WHERE id = 4;

UPDATE users SET
name = "Budi21",
email = "budi21@mail.com"
WHERE id = 4;

-- DELETE
DELETE FROM users
WHERE id=4;


-- LIKE / BETWEEN
select * from products;
select * from products where jenis="makanan";
select * from products where product_name like "wafer%";

select * from products where product_name like "%o%";

select * from products where id between 1 and 3;

-- AND OR
select * from products where id between 1 and 3 OR jenis="minuman";
select * from products where id between 1 and 3 AND jenis="minuman";
select * from products where id between 1 and 3 AND jenis="minuman" OR stock=10;
select * from products where id between 1 and 3 AND jenis="makanan" AND price=50000;

-- ORDER BY
-- ASC : kecil ke besar
-- DESC: besar ke kecil
select * from products ORDER BY price DESC;
select * from products ORDER BY price ASC;
select * from products ORDER BY product_name DESC;

-- LIMIT
select * from products ORDER BY price DESC LIMIT 2;

-- JOIN
select * from products;
select * from users;

SELECT products.id, products.product_name, products.stock, products.price, products.user_id, users.name, users.email
FROM products
INNER JOIN users ON products.user_id = users.id;

SELECT p.id, p.product_name, p.stock, p.price, p.user_id, u.name as user_name, u.email
FROM products p
INNER JOIN users u ON p.user_id = u.id;

-- alias
SELECT p.id, p.product_name, p.stock, p.price, p.user_id, u.name as user_name, u.email
FROM products p
LEFT JOIN users u ON p.user_id = u.id;

SELECT p.id, p.product_name, p.stock, p.price, p.user_id, u.name as user_name, u.email
FROM products p
RIGHT JOIN users u ON p.user_id = u.id;

SELECT u.id, u.name as user_name, u.email, p.id, p.product_name, p.stock, p.price, p.user_id
FROM users u
LEFT JOIN products p ON p.user_id = u.id;


SELECT * FROM users;
SELECT * FROM products;
SELECT * FROM transactions;
SELECT * FROM transaction_details;

-- menampilkan data di table transaction beserta nama usernya (user yang melakukan transaksi)
select t.id, t.user_id, u.name as username, t.total_price, t.created_at
from transactions as t
inner join users as u on t.user_id=u.id;

select td.id, td.transaction_id, td.product_id, p.product_name, u.name, td.quantity
from transaction_details as td
inner join products as p on p.id = td.product_id
inner join users as u on u.id = p.user_id;

-- tampilkan product yang pernah dibeli oleh user dengan id 2

select p.product_name, t.user_id, t.created_at from products as p
inner join transaction_details as td on p.id = td.product_id
inner join transactions as t on td.transaction_id = t.id
where t.user_id = 2;

select p.product_name, t.user_id, u.name as nama_pembeli, t.created_at from products as p
inner join transaction_details as td on p.id = td.product_id
inner join transactions as t on td.transaction_id = t.id
inner join users as u on t.user_id=u.id
where t.user_id = 2;


select p.product_name, td.quantity, p.user_id as id_pemilik, upemilik.name as nama_pemilik, t.user_id as id_pembeli, u.name as nama_pembeli, t.created_at 
from products as p
inner join transaction_details as td on p.id = td.product_id
inner join transactions as t on td.transaction_id = t.id
inner join users as u on t.user_id=u.id
inner join users as upemilik on p.user_id = upemilik.id
where t.user_id = 2;

-- RESTRICT : harus menghapus dari child dulu, baru ke parent
select * from users;
select * from products;
select * from transactions;

insert into users(name, email)
values ("Dedi Baru", "dedibaru@mail.com");

insert into users(name, email)
values ("dummy", "dummy@mail.com");

insert into products(user_id, product_name, stock, price, jenis)
values(8, "nutrisari baru", 15, 1000, "minuman");

insert into products(user_id, product_name, stock, price, jenis)
values(8, "jasjus baru", 10, 1000, "minuman");

insert into products(user_id, product_name, stock, price, jenis)
values(1, "oreo", 30, 10000, "makanan");

DELETE from products where id = 6;
DELETE from users where id = 8;


-- CASCADE : ketika kita menghapus data di parent, maka data di child akan ikut ke hapus.
DELETE from users where id = 8;

-- UNION & UNION ALL
SELECT age
FROM Teachers
UNION
SELECT age
FROM Students;

-- AGGREGATION
-- MIN
select * from products;
select MIN(price) from products;
select MIN(stock) from products;

select MIN(price), min(stock), user_id from products group by user_id;
select MAX(product_name) from products where product_name like "wafer kong%";

-- MAX
select MAX(price) from products;

-- SUM
select SUM(stock) from products where user_id = 1;
select SUM(price) from products where user_id = 1;

-- AVG
select AVG(price) from products;
select AVG(price), user_id from products group by user_id;

-- count
select count(product_name) from products;
select count(id) from products where user_id=2;
-- mendapatkan banyaknya product yang dimiliki oleh masing" user
select count(product_name), user_id from products group by user_id;

select max(p.price), u.id, u.name, u.email from products as p
inner join users as u on p.user_id = u.id
group by p.user_id
order by max(p.price) DESC
limit 1;

-- HAVING
-- menampilkan jumlah stock barang milik users
select u.id, u.name, sum(p.stock) from users as u
inner join products as p on p.user_id=u.id
group by u.id;

-- menampilkan jumlah stock barang milik users yang diatas 15
select u.id, u.name, sum(p.stock) from users as u
inner join products as p on p.user_id=u.id
group by u.id
having sum(p.stock) > 15;

select u.id, u.name, count(p.id) from users as u
inner join products as p on p.user_id=u.id
group by u.id
having count(p.id) > 2;

-- SUBQUERY
select * from users;
select * from products;

-- tampilkan user id yang memiliki product
select distinct user_id from products;


-- tampilkan id, nama dan email dari user yang memiliki product
select id, name, email from users where id IN (select distinct user_id from products);
-- select id, name, email from users where id IN (1,2)

-- tampilkan id, nama dan email dari user yang tidak memiliki product
select id, name, email from users where id NOT IN (select distinct user_id from products);

-- tampilkan id, nama dan email dari user yang memiliki product menggunakan pendekatan join
select distinct u.id, u.name, u.email from users as u
inner join products as p on u.id = p.user_id;


-- FUNCTION
-- fungsi untuk menghitung banyaknya product yang dimiliki oleh user.
DELIMITER $$
CREATE FUNCTION f_count_product(user_id_param int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total int;
select count(id) INTO total from products where user_id=user_id_param;
RETURN total;
END$$
DELIMITER ;

-- menampilkan function yang telah dibuat
SHOW FUNCTION STATUS WHERE db='db_be17';

-- menghapus function
DROP FUNCTION `db_be17`.`f_count_product`;

-- menjalankan/memanggil function
select f_count_product(7);


-- SOFT DELETE
select * from transaction_details;
insert into transaction_details(transaction_id, product_id, quantity, total_price)
values(3, 1, 1, 50000);

-- menghapus data menggunakan pendekatan soft delete
update transaction_details set
deleted_at = now()
where id=5;

-- menampilkan data yang belum terhapus
select id, transaction_id, product_id, quantity, total_price, created_at, updated_at, deleted_at
FROM transaction_details
WHERE deleted_at is null;

-- menampilkan data yang telah dihapus
select id, transaction_id, product_id, quantity, total_price, created_at, updated_at, deleted_at
FROM transaction_details
WHERE deleted_at is not null;



