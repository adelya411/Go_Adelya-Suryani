-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM `transactions` WHERE id_user IN(1, 2);

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) FROM transactions WHERE id_user = 1;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(t.id_transaction) AS total_transaction_product_type_2 FROM transactions t LEFT JOIN transaction_details td ON td.id_transaction=t.id_transaction LEFT JOIN products p ON td.id_product=p.id_product WHERE p.id_product_type=2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT p.*, pt.name AS product_type FROM products p LEFT JOIN product_types pt ON p.id_product_type=pt.id_product_type; 

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT t.*, p.name AS product_name, u.name AS user_name FROM transactions t LEFT JOIN users u ON t.id_user=u.id_user LEFT JOIN transaction_details td ON t.id_transaction=td.id_transaction LEFT JOIN products p ON td.id_product=p.id_product;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$
CREATE TRIGGER delete_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN                                     
DECLARE v_id_transaction INT;             
SET v_id_transaction=OLD.id_transaction;              
DELETE FROM transaction_details WHERE id_transaction=v_id_transaction;
END$$
DELIMITER ;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER $$
CREATE TRIGGER update_total_qty
BEFORE DELETE ON transaction_details FOR EACH ROW   
BEGIN                                               
UPDATE transactions SET total_qty=total_qty-(SELECT qty FROM transaction_details WHERE id_transaction=OLD.id_transaction) WHERE id_transaction=OLD.id_transaction;                    
END $$
DELIMITER ;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * FROM products WHERE id_product NOT IN (SELECT id_product FROM transaction_details);
