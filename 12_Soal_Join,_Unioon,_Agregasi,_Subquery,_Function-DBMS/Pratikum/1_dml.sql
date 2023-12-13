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


--     i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)


--     j. Insert 3 product di masing-masing transaksi.


-- 2. Select
--     a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.


--     b. Tampilkan product dengan id = 3.


--     c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.


--     d. Hitung jumlah user / pelanggan dengan status gender Perempuan.


--     e. Tampilkan data pelanggan dengan urutan sesuai nama abjad


--     f. Tampilkan 5 data pada data product


-- 3. Update
--     a. Ubah data product id 1 dengan nama ‘product dummy’.


--     b. Update qty = 3 pada transaction detail dengan product id = 1.


-- 4. Delete
--     a. Delete data pada tabel product dengan id = 1.


--     b. Delete pada pada tabel product dengan product type id 1.
