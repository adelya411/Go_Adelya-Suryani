# Middleware

Middleware

# Middleware
Middleware adalah entitas yang terkait dalam pemrosesan permintaan/respon server.

Contoh middleware pihak ketiga:
- Negroni
- Echo
- Interpose
- Alice

- Setup middleware echo:

    | Echo#Pre() (Dieksekusi sebelum router memproses permintaan)   | Echo#Use() (Dieksekusi setelah router memproses permintaan dan memiliki akses penuh ke echo.Context API)|
    |-----------------|---------------------|
    | HTTPSRedirect          | BodyLimit         |
    | HTTPSWWWRedirect       | Logger  |
    | WWWRedirect         | Gzip  |
    | AddTrailingSlash          | Recover      |
    | RemoveTrailingSlash       | BasicAuth      |
    | MethodeOverride   | JWTAuth       |
    | Rewrite | Secure
    | |CORS
    | |Static

#### CORS Middleware
CORS adalah singkatan dari Cross Origin Resource Sharing yang merupakan sebuah middleware yang digunakan dalam aplikasi web untuk mengatur dan mengontrol izin akses yang diberikan oleh browser ketika permintaan HTTP dilakukan dari asal (origin) yang berbeda. 
- Konfigurasi CORS umum

   | Configuration   | Deskripsi |
    |-----------------|---------------------|
    | Access-Control-Allow-Origin          | Menentukan domain/origin yang dapat mengirim request ke server         |
    | Access-Control-Allow-Headers       | Menentukan request header yang dapat mengirim request ke server  |
    | Access-Control-Allow-Methods         | Menentukan HTTP method yang dapat digunakan ketika mengirim request ke server  |
    | Access-Control-Max-Age          | Menentukan durasi maksimum preflight request yang dapat dilakukan caching      

#### Rate Limiter Middleware
Rate Limiter Middleware adalah sebuah jenis middleware yang digunakan dalam aplikasi web untuk mengatur tingkat penggunaan (rate limit) dari suatu layanan atau endpoint tertentu. Middleware ini berfungsi untuk membatasi jumlah permintaan atau tindakan yang dapat dilakukan oleh pengguna dalam interval waktu tertentu.

#### Log Middleware
- Mencatat informasi permintaan HTTP
- Sebagai jejak kaki
- Sumber data untuk analisi

#### Auth Middleware
Mengapa menggunakan authentikasi? karena sebagai indetifikasi pengguna dan mengamankan data juga informasi. 

Authentication Middleware:
1. Basic Authentication: permintaan teknik otentikasi http, metode ini memerlukan informasi nama pengguna dan kata sandi untuk dimasukkan ke dalam header permintaan. 
![Alt text](image.png)
2. JSON Web Token: jenis middleware yang digunakan dalam aplikasi web untuk melakukan otentikasi dan otorisasi berdasarkan token JWT.
![Alt text](image-1.png)

Middleware: Penjelasan Rinci dan Konsep Diluar yang Umum

Definisi Umum:

Definisi Umum: Middleware adalah perangkat lunak yang berada di antara aplikasi atau sistem untuk membantu integrasi, koneksi, dan interaksi antara komponen-komponen yang berbeda.
HTTP Middleware:

Definisi Umum: Dalam konteks web development, HTTP middleware adalah lapisan perangkat lunak yang ditempatkan di antara aplikasi web dan server HTTP.
Rincian Diluar Hal Umum:
Logging Middleware: Merekam log dari setiap permintaan dan respons untuk keperluan audit atau pemecahan masalah.
Authentication Middleware: Memvalidasi identitas pengguna dan memberikan akses ke bagian-bagian tertentu dari aplikasi.
CORS Middleware: Menangani kebijakan Same-Origin Policy (SOP) untuk memungkinkan atau menolak permintaan lintas domain.
Message Queue Middleware:

Definisi Umum: Digunakan untuk mentransmisikan pesan antara komponen atau sistem terdistribusi secara asinkron.
Rincian Diluar Hal Umum:
Publish-Subscribe Pattern: Menggunakan antrian pesan untuk mendukung pola publish-subscribe di mana satu produsen pesan dapat memiliki beberapa konsumen.
Message Brokers: Solusi seperti RabbitMQ atau Apache Kafka yang menyediakan infrastruktur untuk mengirim dan menerima pesan.
Database Middleware:

Definisi Umum: Melibatkan perangkat lunak yang menyediakan antarmuka antara aplikasi dan database.
Rincian Diluar Hal Umum:
Object-Relational Mapping (ORM): Sebuah teknik yang memetakan data antara model objek dalam aplikasi dan tabel dalam basis data relasional.
Connection Pooling Middleware: Mengelola koneksi ke database untuk meningkatkan kinerja dengan mengurangi overhead pembukaan dan penutupan koneksi.
Distributed Middleware:

Definisi Umum: Digunakan dalam sistem terdistribusi untuk menyediakan mekanisme komunikasi dan koordinasi antara layanan atau node.
Rincian Diluar Hal Umum:
Remote Procedure Call (RPC): Memungkinkan pemanggilan fungsi atau metode di server yang berjalan di mesin yang berbeda.
Service Discovery Middleware: Mencari dan menemukan layanan yang tersedia secara dinamis dalam lingkungan terdistribusi.
Security Middleware:

Definisi Umum: Middleware yang fokus pada perlindungan dan penanganan aspek keamanan aplikasi atau sistem.
Rincian Diluar Hal Umum:
Rate Limiting Middleware: Mengontrol jumlah permintaan yang dapat diterima dalam periode waktu tertentu untuk melindungi dari serangan pencarian celah keamanan.
Web Application Firewall (WAF) Middleware: Melindungi aplikasi web dari serangan seperti SQL injection atau cross-site scripting (XSS).
Authentication and Authorization Middleware:

Definisi Umum: Mengelola proses otentikasi pengguna dan memberikan izin akses berdasarkan peran atau hak tertentu.
Rincian Diluar Hal Umum:
OAuth Middleware: Mendukung alur otentikasi OAuth untuk memberikan akses ke pihak ketiga.
Role-Based Access Control (RBAC) Middleware: Mengatur hak akses berdasarkan peran pengguna dalam aplikasi.
Konsep dan Penjelasan Lain Diluar Hal Umum:

Idempotency:

Definisi Umum: Konsep di mana operasi dapat diulang tanpa menghasilkan perubahan yang berbeda.
Rincian Diluar Hal Umum:
Implementing Idempotency in APIs: Menggunakan token atau tanda unik dalam permintaan untuk memastikan bahwa operasi yang sama dapat diulang tanpa efek samping.
Chaos Engineering:

Definisi Umum: Praktik yang melibatkan uji sistem dengan sengaja memasukkan kegagalan untuk mengidentifikasi dan memperbaiki masalah potensial.
Rincian Diluar Hal Umum:
GameDays: Simulasi skenario kegagalan atau kondisi darurat untuk mengukur respons dan keandalan sistem.
Polyglot Programming:

Definisi Umum: Pendekatan di mana pengembang menggunakan berbagai bahasa pemrograman untuk membangun komponen atau layanan yang paling sesuai dengan tugas tertentu.
Rincian Diluar Hal Umum:
Polyglot Persistence: Menggunakan lebih dari satu jenis database atau sistem penyimpanan data sesuai dengan kebutuhan aplikasi.
Infrastructure as Code (IaC):

Definisi Umum: Pendekatan untuk mengelola dan mengotomatiskan infrastruktur melalui kode.
Rincian Diluar Hal Umum:
Configuration Drift: Memantau perubahan konfigurasi sistem untuk memastikan kesesuaian dengan keadaan yang diinginkan.
Cloud-Native Architecture:

Definisi Umum: Pendekatan untuk merancang, mengembangkan, dan mengelola aplikasi yang dioptimalkan untuk lingkungan cloud.
Rincian Diluar Hal Umum:
Microservices Architecture: Membangun aplikasi sebagai kumpulan layanan independen yang dapat diimplementasikan dan dikelola secara terpisah.
Memahami konsep-konsep ini di luar hal umum membantu para profesional teknologi untuk merancang, mengelola, dan meningkatkan sistem secara efisien, serta menghadapi tantangan yang kompleks dalam pengembangan dan pengelolaan aplikasi.

