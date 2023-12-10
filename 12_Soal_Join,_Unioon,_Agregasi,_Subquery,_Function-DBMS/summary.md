# Join - Union - Agregasi - Subquery - Function (DBMS)

Database : Schema & Data Defition Language

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

Sebagai seorang ahli informatika, pemahaman yang mendalam tentang Schema dan Data Definition Language (DDL) sangat penting dalam merancang, mengelola, dan memahami basis data. Berikut adalah beberapa rincian diluar hal umum yang seorang ahli informatika sebaiknya ketahui:

1. Database Independence:

Memahami konsep kemandirian database. Beberapa Sistem Manajemen Basis Data (DBMS) menyediakan fitur untuk mengisolasi kode aplikasi dari detail implementasi spesifik database, sehingga dapat dengan mudah berpindah dari satu DBMS ke DBMS lainnya tanpa perubahan besar dalam kode aplikasi.
2. Normalization:

Pemahaman yang mendalam tentang normalisasi dan keuntungannya. Normalisasi adalah proses desain basis data untuk mengurangi redundansi dan meningkatkan integritas data.
3. Denormalization:

Mengetahui kapan dan mengapa harus menggunakan denormalisasi. Terkadang, untuk meningkatkan kinerja query, dilakukan denormalisasi dengan mengorbankan sebagian kecil dari normalisasi.
4. Distributed Databases:

Memahami konsep basis data terdistribusi, di mana data tersebar di beberapa lokasi fisik atau server. Mengetahui tantangan, manfaat, dan strategi manajemen data terdistribusi.
5. Database Design Patterns:

Mengetahui pola desain (design patterns) dalam basis data. Ini mencakup pola untuk hierarki data, pola partisi, dan pola lainnya yang dapat meningkatkan kinerja dan keamanan.
6. Version Control for Database Schema:

Mengelola versi skema database. Pemahaman tentang bagaimana mengelola perubahan skema database secara terkontrol, mempertahankan konsistensi dan integritas data.
7. Database Refactoring:

Memahami konsep refactoring di tingkat basis data. Menerapkan perubahan pada skema dan struktur basis data dengan aman tanpa mengorbankan integritas data.
8. Schema Evolution:

Mengetahui cara mengelola evolusi skema seiring waktu. Dengan pertumbuhan aplikasi, skema database dapat mengalami perubahan yang perlu dikelola secara efisien.
9. Metadata Management:

Memahami arti metadata dalam basis data. Metadata adalah informasi tentang struktur dan properti data. Mengelola metadata penting untuk pengoptimalan dan manajemen sistem.
10. Data Warehousing Concepts:

Pemahaman tentang data warehousing dan data mart. Konsep ini melibatkan penyimpanan dan analisis data besar-besaran untuk mendukung pengambilan keputusan bisnis.
11. Data Archiving and Purging:

Mengetahui strategi dan konsep data archiving dan data purging. Ini melibatkan mengelola data lama, tidak aktif, atau tidak diperlukan agar tidak membebani performa sistem.
12. Database Security:

Memahami keamanan basis data dan implementasi kontrol akses. Mengetahui cara melindungi data sensitif dan menerapkan audit trail.
13. Query Optimization:

Mengetahui teknik dan strategi pengoptimalan kueri. Memahami eksekusi rencana kueri, indeks, dan statistik pengoptimalan kueri.
14. Sharding:

Pemahaman tentang sharding, yaitu membagi data menjadi bagian-bagian yang lebih kecil dan mendistribusikannya di beberapa server untuk meningkatkan kinerja dan skalabilitas.
15. In-memory Databases:

Memahami basis data in-memory dan keuntungan penggunaannya dalam situasi tertentu.
Pemahaman mendalam tentang konsep-konsep ini membantu seorang ahli informatika merancang dan mengelola basis data dengan efisien, mengatasi tantangan yang muncul, dan menjaga konsistensi dan keamanan data.