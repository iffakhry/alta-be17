-- DDL 
-- membuat database
CREATE DATABASE db_be17;

-- memilih database yang akan dimanage
USE db_be17;

-- membuat table
CREATE TABLE users(
id int primary key auto_increment,
name varchar(100) not null,
email varchar(100) not null unique
);

CREATE TABLE products(
id int primary key auto_increment,
user_id int,
product_name varchar(200) not null,
stock int,
price decimal,
jenis varchar(50),
constraint fk_user_product FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transactions(
id int primary key auto_increment,
user_id int,
total_price decimal,
created_at datetime default current_timestamp,
constraint fk_user_transaction FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transaction_details(
id int primary key auto_increment,
transaction_id int,
product_id int,
quantity int,
total_price decimal,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
deleted_at datetime,
constraint fk_transaction_detail FOREIGN KEY (transaction_id) REFERENCES transactions(id),
constraint fk_product_detail FOREIGN KEY (product_id) REFERENCES products(id)
);
-- menampilkan seluruh table yg ada di db
show tables;

-- menampilkan seluruh colom di sebuah table
show columns from products;

CREATE TABLE dummy(
id int primary key,
name varchar(50)
);

-- memodifikasi table
ALTER TABLE dummy
ADD COLUMN description varchar(100);

ALTER TABLE dummy
MODIFY COLUMN description varchar(255);

-- menghapus table
DROP TABLE dummy;

-- Upload file SQL ke gdrive dengan format BE17 - Nama
-- invite ke email fakhry@alterra.id
