# Datatbase - DDL - DML

Database : Schema & Data Definition Language

Database adalah sekumpulan data yang terorganisir.

Database Relationship:
1. *One to One*, 1 user hanya memiliki 1 foto profil.
2. *One to Many*, 1 user bisa memiliki banyak tweets.
3. *Many to Many*, 1 user bisa memiliki banyak follower user, 1 user bisa di follow banyak user.

Software yang menggunakan Relational Database Model sebagai dasarnya, contoh: MySQL.

Jenis Perintah SQL:
- *DDL (Data Definition Language)**, Digunakan untuk membuat, mengubah, dan menghapus objek database seperti tabel, indeks, dan constraint. **Statement*: Create, Use, Drop, Rename, Alter.
- *DML (Data Manipulation Language)**, Perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database. **Statement Operation*: Insert, Select, Update, Delete. DML Statement: Like/Between, And/Or, Order By, Limit.
- *DCL (Data Control Language)*

Tipe Data MySQL:
- Num
- Huruf
- Date

## Join
Sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel.

Join Standar SQL ANSI:
1. *Inner*, akan mengembalikan baris-baris dari dua tabel atau lebih yang memnuhi syarat.
2. *Left*, akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenal kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join.
3. *Right*, akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yangn memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join.

Union, ada hal yang harus diperhatikan dari union adalah jumlah field yang dikeluarkan/dipanggil harus sama.

## Aggregate 
Fungsi agregasi adalah fungsi dimana nilai beberaap baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

Fungsi Agregasi SQL:
1. *Min*, fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.
2. *Max*, Digunakan untuk mendapatkan nilai maximum atau nilai terbesar dari sebuah data record ditabel.
3. *Sum*, Digunakan untuk mnedapatkan jumlah total nilai dari sebuah data atau record ditabel.
4. *Avg*, digunakan untuk mencari nilai rata-rata dari sebuah data atau record ditabel.
5. *Count*, digunakan untuk mencari jumlah dari sebuah data atau record ditabel.
6. *Having*, digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi aggregat.

## SubQuery
Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain.

Sebuah subquery  digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.

## Function
Sebuah kumpulan statetment yang akan mengembalikan sebuah nilai balik pada pemanggilnya.

Database: Schema & Data Definition Language (DDL) - Rincian Diluar Hal Umum:

1. Schema:

Definisi Umum: Schema adalah struktur logis yang mendefinisikan cara data disimpan dan diatur di dalam database. Ini mencakup tabel, kolom, indeks, dan ketergantungan lainnya.
Rincian Diluar Hal Umum:
Public vs. Private Schema: Beberapa database mendukung konsep skema publik dan pribadi. Skema publik dapat diakses oleh banyak pengguna, sementara skema pribadi hanya dapat diakses oleh pemiliknya.
Default Schema: Beberapa database mendukung pengaturan skema default untuk pengguna tertentu.
2. Data Definition Language (DDL):

Definisi Umum: DDL adalah bagian dari SQL yang digunakan untuk mendefinisikan dan mengelola struktur database. Ini mencakup pembuatan, modifikasi, dan penghapusan objek database.
Rincian Diluar Hal Umum:
CREATE Statement:
CREATE TABLE: Membuat tabel baru.
CREATE INDEX: Membuat indeks pada tabel.
CREATE VIEW: Membuat view untuk mengakses data dari satu atau lebih tabel.
CREATE SCHEMA: Membuat skema baru.
CREATE DATABASE: Membuat database baru.
ALTER Statement:
ALTER TABLE: Mengubah struktur tabel, seperti menambah atau menghapus kolom.
ALTER INDEX: Mengubah indeks, seperti menambah atau menghapus kolom indeks.
DROP Statement:
DROP TABLE: Menghapus tabel beserta semua data yang ada di dalamnya.
DROP INDEX: Menghapus indeks dari tabel.
DROP VIEW: Menghapus view.
Constraints:
PRIMARY KEY: Menentukan kolom sebagai kunci utama.
FOREIGN KEY: Menetapkan kunci asing ke tabel lain.
UNIQUE: Menentukan bahwa nilai dalam kolom harus unik.
CHECK: Memeriksa kondisi untuk nilai di kolom tertentu.
Triggers:
DDL juga dapat digunakan untuk membuat dan mengelola trigger, yang merupakan tindakan yang diaktifkan secara otomatis ketika suatu peristiwa terjadi pada tabel.
3. Data Types:

Definisi Umum: Data types menentukan jenis nilai yang dapat disimpan dalam suatu kolom.
Rincian Diluar Hal Umum:
Spatial Data Types: Beberapa database mendukung tipe data spasial untuk memanipulasi data geometri dan geografis.
JSON Data Types: Beberapa database modern mendukung tipe data khusus untuk penyimpanan dan manipulasi data JSON.
4. Views:

Definisi Umum: View adalah hasil dari perintah SELECT yang disimpan dalam basis data dengan nama dan terlihat seperti tabel.
Rincian Diluar Hal Umum:
Indexed Views: Beberapa database mendukung pembuatan indeks pada view untuk meningkatkan performa.
5. Stored Procedures dan Functions:

Definisi Umum: Stored procedures dan functions adalah objek DDL yang dapat menyimpan logika bisnis dan dapat dipanggil oleh aplikasi atau prosedur lain.
Rincian Diluar Hal Umum:
User-Defined Functions (UDF): Beberapa database mendukung pembuatan fungsi pengguna yang dapat digunakan dalam pernyataan SQL.
Triggers: Triggers adalah jenis stored procedure yang secara otomatis dijalankan ketika suatu peristiwa tertentu terjadi pada tabel.
6. Security and Privileges:

Definisi Umum: DDL dapat digunakan untuk mengatur hak akses dan keamanan database.
Rincian Diluar Hal Umum:
Roles: Beberapa database mendukung konsep peran, di mana hak akses diberikan kepada peran dan pengguna ditugaskan ke peran tertentu.
Row-Level Security: Beberapa database mendukung keamanan tingkat baris, di mana hak akses ditetapkan untuk baris data tertentu.
7. Data Partitioning:

Definisi Umum: Beberapa database mendukung data partitioning, di mana tabel besar dibagi menjadi bagian-bagian yang lebih kecil untuk meningkatkan performa dan manajemen data.
Rincian Diluar Hal Umum:
Range, List, Hash Partitioning: Metode pembagian dapat berupa berdasarkan rentang nilai, daftar nilai, atau fungsi hash tertentu.
Pemahaman mendalam tentang DDL dan konsep-konsep terkaitnya memungkinkan pengembang dan administrator database untuk merancang, mengelola, dan mengoptimalkan struktur database dengan lebih baik. Hal-hal ini menjadi kritis ketika berurusan dengan database besar dan kompleks yang digunakan dalam lingkungan bisnis atau proyek skala besar.

