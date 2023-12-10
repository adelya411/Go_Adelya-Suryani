# Docker

Infrastructure and Deployment

Docker adalah platform perangkat lunak yang digunakan untuk mengembangkan, mengemas, dan menjalankan aplikasi dalam wadah (container) yang ringan. Ini memungkinkan pengembang untuk mengisolasi aplikasi dan semua dependensinya dalam lingkungan yang konsisten, portabel, dan dapat diimplementasikan di berbagai platform, termasuk mesin fisik, mesin virtual, atau cloud.

*Docker basic*
1. Image
2. Container
3. Engine
4. Registry
5. Control Plane

## Container
Kontainer adalah unit standar yang dapat dijalankan yang berisi aplikasi dan semua dependensinya. Kontainer memungkinkan pengembang untuk mengemas aplikasi dan semua komponennya (pustaka, kode, variabel lingkungan, dll.) ke dalam lingkungan terisolasi yang konsisten dan portabel.

#### Perbedaan Container dan Virtual Machines
Container :

1. Abstraksi pada lapisan aplikasi
2. Kontainer memakan lebih sedikit ruang dibandingkan VM
3. Menangani lebih banyak aplikasi dan memerlukan lebih sedikit
4. VM dan Sistem Operasi

Virtual Machines :
1. Abstraksi perangkat keras fisik
2. Setiap VM menyertakan salinan lengkap Sistem operasi
3. Juga lambat untuk boot

#### Syntax Docker
- FROM -> Mendapatkan gambar dari registri buruh pelabuhan
- RUN -> Jalankan perintah bash saat membuat Kontainer
- ENV -> Tetapkan variabel di dalam Kontainer
- ADD -> Salin file dengan beberapa proses lainnya
- COPY -> Salin file
- WORKDIR -> Atur direktori file kerja
- ENTRYPOINT -> Jalankan perintah setelah selesai membangun Kontainer
- CMD -> Jalankan perintah tetapi dapat ditimpa

## Introduction Deployment
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. 

*Strategi Deployment*
1. Big-Bang Deployment Strategy
2. Rollout Deployment Strategy
3. Blue/Green Deployment Strategy
4. Canary Deployment Strategy

### Docker Volume
Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container. Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker.

### Docker Network
Defaultnya container pada docker akan saling terisolasi satu sama lain. Kita tidak dapat melakukan request api (misal) dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama.

*Infrastructure and Deployment: Rincian yang Tidak Termasuk Hal Umum*

1. *Infrastructure as Code (IaC):*
   - *Definisi Umum:* IaC melibatkan mendefinisikan dan mengelola infrastruktur secara otomatis melalui kode.
   - *Rincian Diluar Hal Umum:*
      - *Immutable Infrastructure:* Menerapkan prinsip di mana infrastruktur tidak diubah secara langsung; sebaliknya, setiap perubahan menciptakan instance baru.

2. *Continuous Deployment vs. Continuous Delivery:*
   - *Definisi Umum:* Continuous Deployment (CD) adalah praktik merilis perubahan ke produksi secara otomatis setelah melalui tahap pengujian, sementara Continuous Delivery (CD) membawa perubahan ke tahap produksi dengan kemungkinan rilis manual.
   - *Rincian Diluar Hal Umum:*
      - *Feature Toggles:* Memungkinkan mengaktifkan atau menonaktifkan fitur tertentu secara dinamis dalam lingkungan produksi.

3. *Blue-Green Deployment:*
   - *Definisi Umum:* Metode di mana dua lingkungan produksi identik (blue dan green) berjalan bersamaan, dan perpindahan dilakukan dengan mengalihkan lalu lintas dari satu ke yang lain.
   - *Rincian Diluar Hal Umum:*
      - *Canary Releases:* Melibatkan merilis perubahan kepada sebagian kecil pengguna untuk menguji dampaknya sebelum merilis ke seluruh lingkungan.

4. *Serverless Architecture:*
   - *Definisi Umum:* Serverless Architecture menghilangkan kebutuhan untuk mengelola infrastruktur server secara langsung, dengan fokus pada menjalankan kode tanpa harus memikirkan tentang server.
   - *Rincian Diluar Hal Umum:*
      - *Cold Starts:* Waktu yang diperlukan untuk memulai fungsi serverless yang tidak aktif.

5. *Infrastructure Monitoring and Observability:*
   - *Definisi Umum:* Memantau dan mengamati infrastruktur untuk memahami dan mengatasi isu kinerja dan keandalan.
   - *Rincian Diluar Hal Umum:*
      - *Distributed Tracing:* Pelacakan dan analisis permintaan lintas-servis untuk pemahaman yang lebih baik tentang kinerja aplikasi.

6. *Chaos Engineering:*
   - *Definisi Umum:* Praktik yang melibatkan uji sistem dengan menyengaja memasukkan kegagalan untuk mengidentifikasi dan memitigasi risiko keandalan.
   - *Rincian Diluar Hal Umum:*
      - *GameDays:* Simulasi skenario kegagalan atau kondisi darurat untuk mengukur respons dan keandalan sistem.

7. *Rolling Deployments:*
   - *Definisi Umum:* Jenis implementasi di mana perubahan perlahan diperkenalkan ke dalam lingkungan produksi dengan menggantikan instance lama dengan yang baru secara bertahap.
   - *Rincian Diluar Hal Umum:*
      - *Readiness and Liveness Probes:* Menerapkan pengujian yang otomatis untuk memastikan bahwa aplikasi siap dan hidup.

8. *Infrastructure Security:*
   - *Definisi Umum:* Menyelidiki dan mengamankan infrastruktur dari ancaman keamanan.
   - *Rincian Diluar Hal Umum:*
      - *Zero Trust Security Model:* Mengasumsikan bahwa tidak ada entitas atau sumber yang dapat dipercaya sepenuhnya, bahkan dalam jaringan internal.

9. *DevSecOps:*
   - *Definisi Umum:* Integrasi keamanan dalam siklus hidup pengembangan dan operasional (DevOps) untuk memastikan bahwa keamanan menjadi perhatian utama.
   - *Rincian Diluar Hal Umum:*
      - *Automated Security Testing:* Menggunakan alat otomatis untuk mengidentifikasi dan memitigasi kerentanan keamanan.

10. *Cattle vs. Pets Mentality:*
    - *Definisi Umum:* Konsep dalam manajemen infrastruktur di mana server dianggap sebagai hewan ternak (cattle) yang dapat diganti dan dikelola dengan otomatis, bukan sebagai hewan peliharaan (pets) yang memerlukan perhatian individual.
    - *Rincian Diluar Hal Umum:*
       - *Auto-Healing:* Kemampuan infrastruktur untuk mendeteksi dan memperbaiki otomatis isu atau kegagalan.

11. *Deployment Pipelines:*
    - *Definisi Umum:* Rangkaian otomatis proses pengujian dan rilis yang membawa perubahan dari tahap pengembangan ke produksi.
    - *Rincian Diluar Hal Umum:*
       - *Canary Pipelines:* Memiliki pipeline terpisah untuk rilis canary untuk memastikan tahap pengujian khusus sebelum merilis ke lingkungan produksi utama.

12. *Secrets Management:*
    - *Definisi Umum:* Menangani penyimpanan dan distribusi informasi rahasia seperti kunci API, kata sandi, atau token.
    - *Rincian Diluar Hal Umum:*
       - *Rotation of Secrets:* Rutin mengganti rahasia untuk meningkatkan keamanan.

Memahami dan mengimplementasikan aspek-aspek ini di luar definisi umumnya membantu dalam mengelola infrastruktur dengan efisien, memastikan keamanan, dan memfasilitasi implementasi yang lancar dan andal.
