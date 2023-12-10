# System Design

System Design

## Diagram
Yaitu representasi simbolis dari informasi yang menggunakan teknik visualisasi Tools design : smartdraw, Lucidchart, draw.io, Whimsical, Visio. 
Desain perangkat lunak yang umum digunakan: 
1. Flowchart, teorinya mewakili satu proses, terdapat process, decision, terminator.
2. Use Case, merangkum rincian penggunaan system (dikenal sebagai aktor) dan interaksi mereka dengan sistem.
3. ERD, jenis diagram alur yang menggambarkan bagaimana "entitas" seperti orang, objek, atau konsep berhubungan satu sama lain dengan sistem.
4. HLA, desain sistem secara keseluruhan, memeriksa audiens target. 

Karakteristik utama sistem terdistribusi
- Scalability (kemampuan)
- Reliability (keandalan)
- Availability (ketersediaan)
- Efficiency (efisiensi)
- Serviceability or Manageability (kemudahan servis atau pengelolaan)

#### Job/Work Queue
Dalam perangkat lunak sistem, job queue (kadang disebut antrian batch) adalah struktur data yang dikelola oleh perangkat lunak penjadwalan pekerjaan yang berisi pekerjaan yang akan dijalankan.

#### Load Balancing
Load Balancer (LB) adalah komponen kritis lainnya dari setiap sistem terdistribusi. Ini membantu menyebar lalu lintas di antara sekelompok server untuk meningkatkan responsivitas dan ketersediaan aplikasi, situs web, atau basis data.

Untuk memanfaatkan skalabilitas penuh dan redundansi, kita dapat mencoba menyeimbangkan beban di setiap lapisan sistem. Kita dapat menambahkan Load Balancer (LB) pada tiga tempat berikut:

1. Antara pengguna dan server web.
2. Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache.
3. Antara lapisan platform internal dan basis data.

#### Monolithic and Microservices
Aplikasi monolithic memiliki satu basis kode dengan beberapa modul. Modul-modul tersebut dibagi menjadi modul untuk fitur bisnis atau fitur teknis. Aplikasi ini memiliki satu sistem pembangunan yang membangun seluruh aplikasi dan/atau dependensinya. Ini juga memiliki satu binary yang dapat dieksekusi atau didistribusikan.

Microservices adalah layanan yang dapat dideploy secara independen yang didasarkan pada domain bisnis tertentu. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur, mereka menawarkan banyak opsi untuk menyelesaikan masalah yang mungkin Anda hadapi. Oleh karena itu, arsitektur microservices didasarkan pada beberapa microservices yang bekerja sama.

#### SQL dan NoSQL

SQL dan NoSQL (atau basis data relasional dan basis data non-relasional).

- Basis data relasional adalah terstruktur dan memiliki skema yang telah ditentukan sebelumnya, seperti buku telepon yang menyimpan nomor telepon dan alamat.
- Basis data non-relasional adalah tidak terstruktur dan memiliki skema dinamis, seperti folder berkas yang dapat berisi segala hal mulai dari alamat dan nomor telepon seseorang hingga 'like' Facebook dan preferensi berbelanja online mereka.

#### Caching
Sebuah cache mirip dengan ingatan jangka pendek: ia memiliki jumlah ruang yang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua tingkat dalam arsitektur, tetapi seringkali ditemukan di tingkat yang paling dekat dengan bagian depan di mana mereka diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani tingkat lebih bawah.

#### Database Replication
##### Redundancy dan Replication
Redundansi adalah duplikasi komponen atau fungsi kritis dari suatu sistem dengan tujuan meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau mekanisme keamanan, atau untuk meningkatkan kinerja sistem sebenarnya.

#### Database Indexing
Pengindeksan adalah cara untuk mengoptimalkan kinerja basis data dengan meminimalkan jumlah akses ke disk yang diperlukan saat sebuah query dijalankan. Ini adalah teknik struktur data yang digunakan untuk dengan cepat menemukan dan mengakses data dalam sebuah basis data.

System Design: Penjelasan Rinci Diluar Hal Umum:

1. Functional Decomposition:

Definisi Umum: Functional decomposition melibatkan pemisahan sistem besar menjadi bagian-bagian yang lebih kecil dan lebih mudah dikelola.
Rincian Diluar Hal Umum:
Top-Down Design: Memulai dengan gambaran besar sistem dan secara bertahap memecahnya menjadi komponen-komponen yang lebih kecil.
Bottom-Up Design: Memulai dengan komponen-komponen kecil dan menggabungkannya menjadi sistem yang lebih besar.
2. Microservices Architecture:

Definisi Umum: Microservices adalah pendekatan untuk membangun sistem yang terdiri dari layanan-layanan kecil dan independen yang dapat diimplementasikan dan dikelola secara terpisah.
Rincian Diluar Hal Umum:
Service Discovery: Sistem harus memiliki mekanisme untuk menemukan dan berkomunikasi dengan layanan-layanan yang ada.
API Gateway: Menyediakan titik akses tunggal untuk klien ke layanan-layanan di dalam sistem.
3. Data Sharding and Partitioning:

Definisi Umum: Data sharding dan partitioning adalah teknik untuk membagi dan mendistribusikan data di antara beberapa node atau server.
Rincian Diluar Hal Umum:
Horizontal Sharding: Membagi data berdasarkan baris atau rentang nilai.
Vertical Sharding: Membagi data berdasarkan kolom atau atribut.
4. Event-Driven Architecture (EDA):

Definisi Umum: EDA adalah pendekatan di mana komponen-komponen sistem berkomunikasi melalui peristiwa atau notifikasi.
Rincian Diluar Hal Umum:
Message Brokers: Digunakan untuk mengirim dan menerima pesan antar komponen.
Event Sourcing: Mencatat peristiwa sebagai urutan kejadian yang dapat merekonstruksi keadaan sistem.
5. Caching Strategies:

Definisi Umum: Caching digunakan untuk menyimpan data yang sering diakses secara lokal untuk meningkatkan kinerja sistem.
Rincian Diluar Hal Umum:
Cache Invalidation: Strategi untuk menghapus atau memperbarui cache saat data yang sesuai berubah.
Write-Through Caching: Menulis ke cache dan penyimpanan utama secara bersamaan.
Write-Behind Caching: Menunda penulisan ke penyimpanan utama hingga cache penuh atau pada interval waktu tertentu.
6. Consistency Models:

Definisi Umum: Konsistensi mengacu pada sejauh mana data di seluruh sistem tetap konsisten.
Rincian Diluar Hal Umum:
Eventual Consistency: Sistem akan menjadi konsisten setelah suatu waktu tanpa adanya pembaruan lebih lanjut.
Strong Consistency: Sistem menjamin konsistensi segera setelah suatu perubahan data.
7. Design Patterns for Scalability:

Definisi Umum: Design patterns adalah solusi umum untuk masalah desain yang sering muncul.
Rincian Diluar Hal Umum:
Load Balancing Patterns: Distribusi lalu lintas secara merata di antara beberapa server.
Sharding Patterns: Membagi data atau beban kerja menjadi unit-unit yang lebih kecil.
Circuit Breaker Patterns: Mengelola kesalahan dan kegagalan sistem dengan menghentikan sementara akses ke sumber daya yang bermasalah.
8. Polyglot Persistence:

Definisi Umum: Menggunakan lebih dari satu sistem manajemen database untuk menyimpan data.
Rincian Diluar Hal Umum:
Event Stores: Menggunakan basis data khusus untuk menyimpan peristiwa sebagai bentuk utama penyimpanan.
Graph Databases: Menggunakan basis data grafik untuk menyimpan dan menjelajahi hubungan antar data.
9. Chaos Engineering:

Definisi Umum: Chaos engineering adalah praktik untuk menguji dan mengukur keandalan sistem di bawah kondisi yang tidak normal atau bermasalah.
Rincian Diluar Hal Umum:
GameDays: Simulasi skenario kegagalan atau kondisi darurat untuk mengukur respons sistem.
10. Observability and Monitoring:
- Definisi Umum: Observability melibatkan kemampuan untuk memahami dan mengukur apa yang terjadi dalam sistem.
- Rincian Diluar Hal Umum:
- Distributed Tracing: Melacak permintaan atau transaksi melalui beberapa layanan dan server.
- Log Aggregation: Mengumpulkan dan menganalisis log dari berbagai sumber.

11. Continuous Integration and Deployment (CI/CD):
- Definisi Umum: CI/CD adalah praktik untuk mengotomatiskan dan menguji setiap perubahan kode sebelum diterapkan ke lingkungan produksi.
- Rincian Diluar Hal Umum:
- Blue-Green Deployments: Menerapkan versi baru aplikasi secara paralel dengan versi lama, dan beralih secara langsung saat dianggap stabil.
- Canary Deployments: Meluncurkan versi baru aplikasi kepada sekelompok pengguna terpilih sebelum merilis secara luas.

12. Stateful vs. Stateless Services:
- Definisi Umum: Stateful services menyimpan keadaan dan konteks pengguna, sedangkan stateless services tidak menyimpan keadaan.
- Rincian Diluar Hal Umum:
- Session Affinity: Menetapkan pengguna ke instance yang sama selama sesi mereka untuk menjaga keadaan.

Mempelajari aspek-aspek ini dari sistem design membantu seorang arsitek sistem atau insinyur perangkat lunak membuat keputusan yang lebih baik saat merancang dan membangun sistem yang skala dan dapat diandalkan.