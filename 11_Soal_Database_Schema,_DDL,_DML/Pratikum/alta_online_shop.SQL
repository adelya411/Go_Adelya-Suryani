-- 1. Create database alta_online_shop.
CREATE DATABASE alta_online_shop;
USE alta_online_shop; 

-- 2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
-- a. Create table user.
CREATE TABLE users (
    id_user INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    birth_date DATE,
    status_user VARCHAR(50),
    gender VARCHAR(10),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- b. Create table product, product type, operators, product description, payment_method.
CREATE TABLE products (
    id_product INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    code VARCHAR(10),
    price DOUBLE,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_types (
    id_product_type INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE operators (
    id_operator INT UNIQUE AUTO_INCREMENT  PRIMARY KEY,
    name VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_descriptions (
    id_product_descripstion INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payment_methods (
    id_payment_method INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- c. Create table transaction, transaction detail.
CREATE TABLE transactions (
    id_transaction INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(100),
    total_price DOUBLE,
    total_qty INT 
);

CREATE TABLE transaction_details (
    id_transaction_detail INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    price DOUBLE,
    qty INT(11) 
);

-- 3. Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE kurirs (
    id_kurir INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 4. Tambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurirs 
ADD COLUMN ongkos_dasar DECIMAL(10, 2)
AFTER update_at;

-- 5. Rename tabel kurir menjadi shipping.
ALTER TABLE kurirs RENAME TO shippings;

-- 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shippings;

-- 7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
-- a. 1-to-1: payment method description.
CREATE TABLE payment_method_descriptions (
    id_payment_method_description INT AUTO_INCREMENT PRIMARY KEY,
    id_payment_method INT,
    description TEXT,
    FOREIGN KEY (id_payment_method) REFERENCES payment_methods(id_payment_method)
);

-- b. 1-to-many: user dengan alamat.
CREATE TABLE address (
    id_address INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255),
    id_user INT,
    FOREIGN KEY (id_user) REFERENCES users(id_user)
);

-- c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_details (
    id_user_payment_method_detail INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT,
    id_payment_method INT,
    FOREIGN KEY (id_user) REFERENCES users(id_user),
    FOREIGN KEY (id_payment_method) REFERENCES payment_methods(id_payment_method)
);

