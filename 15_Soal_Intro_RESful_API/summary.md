# Intro RESTful API

Introduction to RESTful API

# RESTful API
API adalah sekumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data sistem operasi, aplikasi atau layanan lainnya.

Request & Response format :
- JSON
- XML
- SOAP

HTTP request Method :
- GET
- POST
- PUT
- DELETE
- HEAD
- OPTION
- PATCH

*Postman* klien HTTP untuk menguji layanan web. postman memudahkan pengujian, pengembangan dan dokumentasi API dengan memungkinkan pengguna dengan cepat menyusun permintaan HTTP yang sederhana dan kompleks.

# Intro restful API
- introduction Api
- introduction to postman
- REST
- JSON
- HTTP respon code
- package net/http

api bekerja dengan client(browser) mengirimkan request ke server dan kemudians server akan memberikan response yang sesuai dengan requestnya.
 
Implementasi dan kegunaan rest API:
- mengintegrasikan aplikasi yang berbasis front end (web. mobile, dll) dan backend. front end ada banyak dan be hanya 1 tetap bisa digunakan
- mengintegrasikan BE ke BE
misalnya  service sepulsa penyedia layanan pembelian pulsa, sepulsa dapat menggunakan fitur pembayaran, seperti yang disediakan oleh midtrans(service yang melayani pembayaran)
 
### open API
open API adalah API yang dapat digunakan dengan bebas.
contoh-contoh open Api: https://github.com/public-apis/public-apis

1. OpenWeatherMap API: [OpenWeatherMap API](https://openweathermap.org/api)

2. Google Maps API: [Google Maps API](https://developers.google.com/maps)

3. Twitter API: [Twitter Developer](https://developer.twitter.com/en)

4. GitHub API: [GitHub Developer](https://developer.github.com/)

5. YouTube API: [YouTube API](https://developers.google.com/youtube/)

### Best Practice RestAPI
1. use Nouns instead of verbs

    | Baik (Nouns)    | Kurang Baik (Verbs) |
    |-----------------|---------------------|
    | /users          | /createUser         |
    | /products       | /getProductDetails  |
    | /orders         | /updateOrderStatus  |
    | /posts          | /createNewPost      |
    | /comments       | /deleteComment      |
    | /categories     | /addCategory        |

2. Naming using plurals Nouns
   
    | Baik (Plurals Nouns) | Kurang Baik (Singular Nouns) |
    |-----------------------|------------------------------|
    | /users                | /user                        |
    | /products             | /product                     |
    | /orders               | /order                       |
    | /posts                | /post                        |
    | /comments             | /comment                     |
    | /categories           | /category                    |

3. using resourse nesting to show relations or hierarchy
   
   /users <- user`s list
   /user/1 <- spesific user
   /user/1/orders <- order list that belongs to spesific user 
   /user/1/orders/12<- spesific order to spesifis user

   RESTful API: Pengantar dan Rincian Diluar Hal Umum untuk Teknik Informatika

**1. Definisi Umum: REST (Representational State Transfer) adalah arsitektur untuk komunikasi sistem terdistribusi melalui protokol HTTP. RESTful API merujuk pada implementasi API yang mengikuti prinsip-prinsip REST.

**2. Resource dan URI:

Definisi Umum: Resource dalam RESTful API direpresentasikan sebagai objek atau data yang dapat diakses melalui URI (Uniform Resource Identifier).
Rincian Diluar Hal Umum:
Hierarchical URI: Desain URI dengan struktur hierarki yang mencerminkan relasi antar resource.
Collection Resource: Resource yang merepresentasikan kumpulan objek, misalnya /users.
Nested Resources: Penggunaan URI bersarang untuk merepresentasikan hubungan antar resource, misalnya /users/{userID}/posts.
**3. HTTP Methods:

Definisi Umum: HTTP methods (GET, POST, PUT, DELETE, dll.) digunakan untuk menentukan operasi yang diinginkan pada resource.
Rincian Diluar Hal Umum:
PATCH Method: Digunakan untuk memperbarui sebagian dari resource.
OPTIONS Method: Digunakan untuk mendapatkan informasi tentang resource, seperti metode yang didukung atau izin akses.
**4. Status Codes:

Definisi Umum: Status codes memberikan informasi tentang hasil dari permintaan ke server.
Rincian Diluar Hal Umum:
201 Created: Digunakan ketika resource berhasil dibuat.
204 No Content: Digunakan ketika permintaan berhasil dijalankan tetapi tidak ada konten yang dikembalikan.
**5. Content Negotiation:

Definisi Umum: Content negotiation melibatkan proses negosiasi format data antara client dan server.
Rincian Diluar Hal Umum:
Accept Header: Digunakan oleh client untuk menyatakan format yang diinginkan untuk respons.
Content-Type Header: Digunakan oleh client untuk menyatakan format data yang dikirim dalam permintaan.
**6. HATEOAS (Hypermedia As The Engine Of Application State):

Definisi Umum: HATEOAS adalah prinsip RESTful API di mana setiap respons berisi tautan (link) ke resource terkait.
Rincian Diluar Hal Umum:
Link Relations: Tautan memiliki relasi yang menjelaskan hubungan antar resource.
**7. Statelessness:

Definisi Umum: RESTful API bersifat stateless, artinya setiap permintaan dari client ke server berisi semua informasi yang diperlukan.
Rincian Diluar Hal Umum:
Stateless Authentication: Menggunakan token atau kredensial dalam setiap permintaan untuk otentikasi.
**8. Pagination:

Definisi Umum: Ketika resource menghasilkan sejumlah besar data, API dapat menyediakan mekanisme paginasi untuk mengambil data secara bertahap.
Rincian Diluar Hal Umum:
Cursor-based Pagination: Menggunakan nilai-nilai yang berfungsi sebagai "kursor" untuk mengambil halaman berikutnya.
**9. Rate Limiting:

Definisi Umum: Rate limiting melibatkan pembatasan jumlah permintaan yang dapat dilakukan oleh client dalam periode waktu tertentu.
Rincian Diluar Hal Umum:
Token Bucket Algorithm: Mekanisme rate limiting di mana client diberikan "token" untuk setiap permintaan yang kemudian dikonsumsi.
**10. Versioning:
- Definisi Umum: Versioning digunakan untuk mengelola evolusi API dan perubahan struktur data.
- Rincian Diluar Hal Umum:
- URL Versioning: Menambahkan nomor versi ke dalam URI, misalnya /v1/users.
- Header Versioning: Menentukan versi melalui header permintaan atau respons.

**11. Security (OAuth, JWT):
- Definisi Umum: Keamanan dalam RESTful API melibatkan otentikasi dan otorisasi.
- Rincian Diluar Hal Umum:
- OAuth 2.0: Protokol otentikasi yang digunakan untuk memberikan akses ke aplikasi pihak ketiga.
- JWT (JSON Web Tokens): Format token yang digunakan untuk mentransfer klaim antara dua pihak.

**12. Webhooks:
- Definisi Umum: Webhooks adalah mekanisme di mana server dapat memberi tahu client secara asinkron tentang perubahan atau peristiwa tertentu.
- Rincian Diluar Hal Umum:
- Idempotent Webhooks: Mekanisme yang memastikan bahwa pengulangan pesan tidak memiliki dampak ganda.

Memahami rincian diluar hal umum dari RESTful API memungkinkan seorang profesional teknik informatika untuk merancang, mengimplementasikan, dan mengelola API dengan lebih efektif, serta memenuhi kebutuhan bisnis yang kompleks.

