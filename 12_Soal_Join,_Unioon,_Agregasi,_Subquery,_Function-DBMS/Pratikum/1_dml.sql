-- 1. Insert
--     a. Insert 5 operators pada table operators.
INSERT INTO operators (id_operator, name) VALUES ("1","Operator1"),("2","Operator2"),("3","Operator3"),("4","Operator4"),("5","Operator5");

--     b. Insert 3 product type.
INSERT INTO product_types (id_product_type, name) VALUES ("1","Ciki-ciki"),("2","Perment"),("3","Minuman");

--     c. Insert 2 product dengan product type id = 1, dan operators id = 3. 
INSERT INTO products (id_product, id_product_type, id_product_description, id_operator, name, code, price, stok) VALUES (1, 1, 1, 3, "Pilus", "01A", 1000, 5), (2, 1, 2, 3, "Nabati", "02A", 2000, 10);

--     d. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (id_product, id_product_type, id_product_description, id_operator, name, code, price, stok) VALUES (3, 2, 3, 1, "Kis", "01B", 500, 50), (4, 2, 4, 1, "Kopiko", "02B", 200, 80), (5, 2, 5, 1, "Mentos", "03B", 600, 100);

--     e. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO products (id_product, id_product_type, id_product_description, id_operator, name, code, price, stok) VALUES (6, 3, 6, 4, "Coca-cola", "01C", 10000, 60), (7, 3, 7, 4, "Nutribus", "02C", 7000, 90), (8, 3, 8, 4, "Kulakul", "03C", 1500, 200);

--     f. Insert product description pada setiap product.
INSERT INTO product_descriptions (id_product_description, description) VALUES (1, "Berisi bola kecil berasa kerupuk"), (2, "Wafer yang berlapis"), (3, "Perment mint ada 2 rasa"), (4, "Perment rasa kopi"), (5, "Perment mint bersoda"), (6, "Minuman bersoda tinggi"), (7, "Minuman susu bercampur buah"), (8, "Minuman 500-an");

--     g. Insert 3 payment methods.
INSERT INTO payment_methods (id_payment_method, name) VALUES (1, "Transfer bank"), (2, "Dana"), (3, "Gopay");

--     h. Insert 5 user pada tabel user.
INSERT INTO users (id_user, name, address, birth_date, status_user, gender) VALUES (1, "Sugiono", "Bukittinggi", "2000-09-19", "Pelanggan", "Laki-laki"), (2, "Tantri", "Pasaman", "2001-04-24", "Pelanggan", "Perempuan"), (3, "Chelsi", "Padang", "2003-11-04", "Pelanggan", "Perempuan"), (4, "Juki", "Riau", "2003-03-12", "Pelanggan", "Laki-laki"), (5, "Chelsi", "Padang", "2002-09-07", "Pelanggan", "Perempuan");

--     i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (id_transaction, id_user, id_payment_method, code, total_price, total_qty) VALUES (1, 1, 1, "01A", 3000, 3), (2, 1, 2, "02B", 2000, 10), (3, 1, 3, "03C", 3000, 2), (4, 2, 1, "02A", 4000, 2), (5, 2, 2, "01B", 2000, 4), (6, 2, 3, "02C", 14000, 2), (7, 3, 1, "01A", 1000, 1), (8, 3, 2, "03B", 6000, 10), (9, 3, 3, "01C", 10000, 1), (10, 4, 1, "01A", 2000, 2), (11, 4, 2, "03B", 12000, 20), (12, 4, 3, "03C", 6000, 4), (13, 5, 1, "02A", 2000, 1), (14, 5, 2, "02B", 4000, 20), (15, 5, 3, "02C", 7000, 1);

--     j. Insert 3 product di masing-masing transaksi.
INSERT INTO transaction_details (id_transaction_detail, id_transaction, id_product, price, qty) VALUES (1, 1, 1, 1000, 1), (2, 1, 2, 2000, 2), (3, 1, 3, 500, 10), (4, 2, 4, 200, 10), (5, 2, 5, 600, 10), (6, 2, 6, 100000, 1), (7, 3, 7, 7000, 2), (8, 3, 8, 1500, 2), (9, 3, 1, 1000, 4), (10, 4, 2, 2000, 1), (11, 4, 3, 500, 20), (12, 4, 4, 200, 20), (13, 5, 5, 600, 20), (14, 5, 6, 100000, 1), (15, 5, 7, 7000, 1);

-- 2. Select
--     a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT * FROM `users` WHERE gender = "Laki-laki";

--     b. Tampilkan product dengan id = 3.
SELECT * FROM `products` WHERE id_product = 3;

--     c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM `users` WHERE name LIKE '%a%'  AND created_at >= DATE_SUB(CURDATE(), INTERVAL 7 DAY) AND created_at <= CURDATE();

--     d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(id_user) FROM users WHERE gender = 'Perempuan';

--     e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SElECT * FROM users ORDER BY name ASC;

--     f. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;

-- 3. Update
--     a. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products
SET name = 'product dummy'
WHERE id_product = 1;

SELECT * FROM products where id_product = 1;
    
--     b. Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_details
SET qty = 3
WHERE id_product = 1;

SELECT * FROM transaction_details WHERE id_product = 3;

-- 4. Delete
--     a. Delete data pada tabel product dengan id = 1.
DELETE FROM transaction_details WHERE id_product = 3;

SELECT * FROM transaction_details;

--     b. Delete pada pada tabel product dengan product type id 1.
DELETE FROM products WHERE id_product_type = 3;

SELECT * FROM products;
