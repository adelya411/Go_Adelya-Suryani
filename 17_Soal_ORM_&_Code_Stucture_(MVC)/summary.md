# ORM and Code Structure (MVC)

ORM & Code Structure

## ORM
 ORM adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel dengan menggunakan bahasa pemrograman berorientasi objek.

 Keuntungan:
 1. Lebih sedikit kueri berulang
 2. Mengambil data secara otomatis ke dalam objek siap pakai
 3. Cara sederhana jika ingin menyaring data sebelum menyimpannya dalam database
 4. Beberapa memiliki fitur cache query

 Kelemahan:
 1. Tambahkan lapisan dalam kode dan timbulkan overhead proses
 2. Memuat data hubungan yang tidak perlu
 3. Kueri SQL mentah yang kompleks bisa sulit ditulis dengan ORM (>10 join tabel)
 4. Fungsi SQL tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi tertentu (contoh: MySQL: FORCE INDEX)

 ### Database Migration
 Cara untuk memperbarui versi database agar sesuai dengan perubahan versi aplikasi.

Perubahan dapat berupa pembaruan ke versi terbaru atau pengembalian ke versi sebelumnya.

*Database Relations in GORM:*
1. Belongs to, Cocok untuk hubungan yang saling terhubung. 
2. Has one, Cocok untuk mewakili entitas tunggal yang hanya memiliki satu entitas tepat.
3. Has many, Cocok untuk mewakili entitas tunggal yang memiliki banyak entitas.
4. Many to many, Cocok untuk mewakili banyak entitas yang memiliki banyak entitas. Karakteristik utama dari hubungan ini adalah penggunaan tabel pivot.

*Database Transaction in GORM:*
- Transaksi adalah urutan operasi database. Jika salah satu operasi gagal, maka seluruh transaksi dianggap gagal. 
-  Transaksi berguna untuk memastikan integritas dan konsistensi data. 
- Transaksi didukung dalam GORM.

*Install GORM:* 
- https://gorm.io/docs/
- https://gorm.io/docs/migration.html 

Cara menghubungkan aplikasi Go dengan database :
1. Buat InitDB() untuk Koneksi Database, Fungsi InitDB() untuk membuat koneksi aplikasi ke database.
2. Buat Skema Pengguna & InitialMigration(), Buat skema pengguna untuk database menggunakan Struct (Objek). fungsi InitialMigration() untuk membuat database di MySQL.
3. Panggil InitDB() dan InitialMigration(), Panggil fungsi InitDB() dan InitialMigration() untuk menggunakannya.
4. Buat GetUsersController(), Buat GetUserController() untuk mengambil data pengguna dari basis data (ORM) dan menampilkan data dalam respons.
5. Buat CreateUserController() untuk mengikat data pengguna dari klien dan menyimpan pengguna ke database.
6. Routing, Tentukan routing API RESTful dengan memanggil Controller.

## Code Structure
Menyusun proyek Anda menggunakan Model-View-Controller (MVC)

Alasan menggunakan structure:
- Untuk mencapai aplikasi modular.
- Terapkan pemisahan perhatian.
- Konflik versi yang lebih sedikit.

Code Structuring:
1. Buat database di mysql ![Alt text](mvc1.png)
2. Buat folder structure ![Alt text](image-1.png)
3. Buat configuration ![Alt text](image-2.png)
4. Buat model ![Alt text](image-3.png)
5. Buat lib database 
6. Buat controller ![Alt text](image-5.png)
7. Buat router 
8. Buat main ![Alt text](image-7.png)
9. Run ![Alt text](image-9.png)

ORM (Object-Relational Mapping): Penjelasan Rinci Diluar yang Sudah Diketahui

**1. Lazy Loading vs. Eager Loading:

Definisi Umum: Lazy loading adalah strategi di mana data diambil dari database hanya saat diperlukan. Eager loading, sebaliknya, mengambil semua data terkait sejak awal.
Rincian Diluar Hal Umum:
Batch Loading: Menggunakan teknik batch untuk mengoptimalkan pengambilan data dalam jumlah besar.
**2. Caching Strategies:

Definisi Umum: Caching melibatkan penyimpanan sementara data yang sering diakses untuk mengoptimalkan kueri ke database.
Rincian Diluar Hal Umum:
Second-Level Caching: Caching yang melibatkan tingkat kedua, di luar tingkat pertama yang umumnya diimplementasikan oleh penyedia database.
**3. Composite Keys:

Definisi Umum: ORM mendukung penggunaan kunci komposit, di mana primary key terdiri dari dua atau lebih kolom.
Rincian Diluar Hal Umum:
Embedded Objects: Menggunakan objek yang bersifat "embedded" sebagai bagian dari entitas utama.
**4. Custom Query Language:

Definisi Umum: Beberapa ORM menyediakan bahasa kueri khusus yang memungkinkan pengguna menulis kueri kompleks.
Rincian Diluar Hal Umum:
Native SQL Queries: Menyediakan kemampuan untuk mengeksekusi kueri SQL mentah jika diperlukan.
**5. Inheritance Mapping:

Definisi Umum: ORM mendukung pemetaan pewarisan dalam pemodelan objek ke struktur database.
Rincian Diluar Hal Umum:
Table-per-Hierarchy: Mewakili seluruh hierarki pewarisan dalam satu tabel.
Table-per-Type: Mewakili setiap jenis objek dalam tabel terpisah.
**6. Database Migrations:

Definisi Umum: ORM sering menyediakan alat untuk mengelola dan menerapkan perubahan skema database.
Rincian Diluar Hal Umum:
Versioned Migrations: Menyimpan versi perubahan skema sehingga dapat diaplikasikan dalam urutan yang benar.
**7. Concurrency Control:

Definisi Umum: ORM dapat menangani masalah konkurensi, seperti konflik saat dua pengguna mencoba mengubah data yang sama secara bersamaan.
Rincian Diluar Hal Umum:
Optimistic Concurrency Control: Mengecek apakah data telah berubah sebelum menyimpan perubahan.
**8. Custom Data Types:

Definisi Umum: Mendukung penggunaan tipe data khusus yang tidak dikenali secara langsung oleh basis data.
Rincian Diluar Hal Umum:
User-Defined Types: Mewakili tipe data yang dibuat oleh pengguna dalam kode aplikasi.
Code Structure: Penjelasan Rinci Diluar yang Sudah Diketahui

**1. Dependency Injection (DI):

Definisi Umum: DI adalah praktik menyuntikkan dependensi ke dalam objek atau kelas untuk mengurangi ketergantungan dan memfasilitasi pengujian.
Rincian Diluar Hal Umum:
Constructor Injection: Menyuntikkan dependensi melalui konstruktor.
Setter Injection: Menyuntikkan dependensi melalui metode setter.
**2. Aspect-Oriented Programming (AOP):

Definisi Umum: AOP memisahkan aspek atau kepentingan tertentu dari logika bisnis utama.
Rincian Diluar Hal Umum:
Cross-Cutting Concerns: Memisahkan kepentingan yang melintasi berbagai bagian aplikasi, seperti logging atau keamanan.
**3. Repository Pattern:

Definisi Umum: Repository pattern menyediakan antarmuka untuk mengakses data dan menyembunyikan detail implementasinya.
Rincian Diluar Hal Umum:
Generic Repositories: Menggunakan repositori generik untuk mengakses berbagai jenis entitas.
**4. Service Layer:

Definisi Umum: Service layer menyediakan fungsi-fungsi bisnis dan memisahkan logika bisnis dari antarmuka pengguna.
Rincian Diluar Hal Umum:
Transaction Management: Menangani transaksi untuk memastikan keintegritasan data.
**5. Domain-Driven Design (DDD):

Definisi Umum: DDD adalah pendekatan untuk merancang aplikasi dengan memahami dan memodelkan domain bisnis.
Rincian Diluar Hal Umum:
Bounded Contexts: Menetapkan batasan pada domain tertentu untuk menghindari konflik model.
**6. Clean Architecture:

Definisi Umum: Clean Architecture mempromosikan pemisahan konsep dan bertujuan untuk membuat sistem yang dapat diuji dan mudah dimengerti.
Rincian Diluar Hal Umum:
Entities, Use Cases, Gateways: Menggunakan konsep ini untuk merancang lapisan inti aplikasi.
**7. Domain Events:

Definisi Umum: Domain events merepresentasikan peristiwa penting dalam domain bisnis dan digunakan untuk mengkomunikasikan perubahan ke berbagai bagian sistem.
Rincian Diluar Hal Umum:
Event Handlers: Menangani peristiwa dan meresponsnya dengan melakukan tindakan tertentu.
**8. Hexagonal Architecture (Ports and Adapters):

Definisi Umum: Hexagonal Architecture memisahkan inti aplikasi dari detail teknis dan infrastruktur.
Rincian Diluar Hal Umum:
Ports and Adapters: Menggunakan konsep ini untuk memisahkan antara inti aplikasi (ports) dan infrastruktur atau detail teknis (adapters).
**9. **Command Query Responsibility Segregation (CQRS)