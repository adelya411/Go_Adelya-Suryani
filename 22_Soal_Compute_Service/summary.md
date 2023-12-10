# Compute Services

Infrastructure and Deployment

# AWS EC2 & RDS
## EC2
Amazon Elastic Compute Cloud (EC2) adalah layanan komputasi awan yang disediakan oleh Amazon Web Services (AWS). EC2 memungkinkan Anda untuk menyewa kapasitas komputasi virtual dalam lingkungan cloud AWS.

Berikut beberapa poin kunci tentang Amazon EC2:
1. Instans Komputasi Virtual
2. Fleksibilitas
3. Pembayaran Berdasarkan Penggunaan
4. Akses Root
5. Skalabilitas
6. Keamanan
7. Banyak OS Dukungan
8. Jangkauan Global

## RDS
Amazon Relational Database Service (RDS) adalah layanan manajemen database yang disediakan oleh Amazon Web Services (AWS). RDS dirancang untuk menyederhanakan pengelolaan database relasional, seperti MySQL, PostgreSQL, Oracle, Microsoft SQL Server, dan beberapa mesin database lainnya.

 Berikut beberapa poin penting tentang Amazon RDS:
 1. Database Relasional Manajemen
 2. Dukungan untuk Berbagai Mesin Database
 3. Skalabilitas Otomatis
 4. Pencadangan Otomatis
 5. Keamanan
 6. Kinerja dan Pemantauan
 7. Multi-AZ Deployment
 8. Pemeliharaan Dikelola
 9. Pilihan Penyimpanan
 10. Jangkauan Global

## GCP
Google Cloud Platform (GCP) adalah salah satu penyedia layanan cloud computing terkemuka di dunia yang disediakan oleh Google. GCP memberikan infrastruktur dan berbagai layanan cloud untuk membantu perusahaan dan individu dalam mengelola aplikasi, data, dan sumber daya komputasi mereka secara efisien.

 ## Continuous
 ### Continuous Integration
 Adalah proses otomatis. Hal ini dilakukan untuk mengintegrasikan berbagai kode dari berbagai sumber potensial untuk membangun atau mengujinya. Penerapan Continuous Integration membantu meningkatkan produktivitas, kualitas perangkat lunak, dan kecepatan pengiriman, serta mengurangi risiko kesalahan saat mengintegrasikan kode dari berbagai kontributor. Ini adalah praktik penting dalam pengembangan perangkat lunak modern yang mendukung pendekatan Agile dan DevOps.

 *Compute Services: Rincian yang Tidak Umum pada Infrastructure and Deployment*

1. *Serverless Computing:*
   - *Definisi Umum:* Paradigma komputasi di mana pengembang tidak perlu memikirkan atau mengelola infrastruktur server secara langsung, hanya perlu fokus pada kode yang dijalankan sebagai fungsi atau layanan.
   - *Rincian Diluar Hal Umum:*
      - *Event-Driven Architecture:* Membangun fungsi atau layanan yang merespons terhadap peristiwa atau trigger tertentu.

2. *Container Orchestration:*
   - *Definisi Umum:* Pengelolaan dan otomatisasi penyebaran, skala, dan operasi container (seperti Docker) dalam lingkungan produksi.
   - *Rincian Diluar Hal Umum:*
      - *Service Mesh:* Menambahkan lapisan manajemen trafik dan komunikasi antara layanan dalam infrastruktur kontainer.

3. *Edge Computing:*
   - *Definisi Umum:* Memindahkan pemrosesan data dan komputasi lebih dekat ke sumber data atau perangkat pengguna, mengurangi latensi dan meningkatkan kinerja.
   - *Rincian Diluar Hal Umum:*
      - *Fog Computing:* Konsep yang melibatkan komputasi di lapisan perbatasan (edge) dan di sepanjang jaringan (fog) untuk mengelola volume data yang besar secara efisien.

4. *Function as a Service (FaaS):*
   - *Definisi Umum:* Layanan komputasi di mana pengembang mendeploy fungsi ke cloud tanpa harus mengelola server atau infrastruktur.
   - *Rincian Diluar Hal Umum:*
      - *Stateless Functionality:* Fungsi yang tidak menyimpan status antara panggilan, memungkinkan skalabilitas horizontal yang lebih baik.

5. *Bare Metal Cloud Computing:*
   - *Definisi Umum:* Menyediakan akses langsung ke sumber daya fisik tanpa virtualisasi, memberikan kinerja yang lebih tinggi dan kontrol lebih besar.
   - *Rincian Diluar Hal Umum:*
      - *Custom Hardware Provisioning:* Kemampuan untuk menyediakan dan mengelola perangkat keras khusus untuk memenuhi kebutuhan aplikasi tertentu.

6. *Immutable Infrastructure:*
   - *Definisi Umum:* Pendekatan di mana setiap perubahan pada infrastruktur menciptakan instance baru daripada mengubah instance yang ada.
   - *Rincian Diluar Hal Umum:*
      - *Phoenix Servers:* Menerapkan prinsip bahwa server dapat "dibangun kembali" atau direkonfigurasi sepenuhnya dari awal.

7. *Multi-Cloud Deployments:*
   - *Definisi Umum:* Strategi menggunakan lebih dari satu penyedia layanan cloud untuk meningkatkan keandalan dan menghindari ketergantungan penuh pada satu penyedia.
   - *Rincian Diluar Hal Umum:*
      - *Cloud Interoperability Standards:* Menerapkan standar untuk memastikan portabilitas aplikasi dan data di antara penyedia cloud yang berbeda.

8. *Dark Launching:*
   - *Definisi Umum:* Mekanisme peluncuran fitur atau perubahan secara bertahap, diaktifkan hanya untuk sebagian pengguna atau dalam kondisi tertentu.
   - *Rincian Diluar Hal Umum:*
      - *A/B Testing Infrastructure:* Menyediakan infrastruktur yang memungkinkan uji A/B untuk mengukur kinerja fitur baru atau perubahan.

9. *Serverless Containers:*
   - *Definisi Umum:* Menggabungkan konsep serverless dengan containerization, memungkinkan pengembang menjalankan kode dalam wadah tanpa harus mengelola infrastruktur kontainer secara eksplisit.
   - *Rincian Diluar Hal Umum:*
      - *Container Orchestration Integration:* Integrasi dengan platform orkestrasi kontainer (seperti Kubernetes) untuk manajemen dan otomatisasi.

10. *Serverless Frameworks:*
    - *Definisi Umum:* Frameworks yang menyederhanakan pengembangan aplikasi serverless dengan menyediakan alat, layanan, dan integrasi.
    - *Rincian Diluar Hal Umum:*
       - *Local Emulation:* Kemampuan untuk menguji dan mengembangkan fungsi serverless secara lokal sebelum dideploy ke cloud.

11. *Hybrid Cloud Deployments:*
    - *Definisi Umum:* Menggabungkan penggunaan infrastruktur cloud publik dan privat, memungkinkan perusahaan untuk memanfaatkan keuntungan dari keduanya.
    - *Rincian Diluar Hal Umum:*
       - *Data Gravity Considerations:* Pertimbangan yang berkaitan dengan lokasi dan volume data dalam pengambilan keputusan tentang penempatan cloud.

12. *Infrastructure Resilience Engineering:*
    - *Definisi Umum:* Disiplin yang berkaitan dengan merancang, mengelola, dan menguji infrastruktur agar tahan terhadap kegagalan dan meningkatkan keandalan.
    - *Rincian Diluar Hal Umum:*
       - *Chaos Engineering Practices:* Melibatkan uji dan eksperimen terkontrol untuk mengidentifikasi dan memitigasi kerentanan sistem.

Memahami dan menerapkan aspek-aspek ini di luar definisi umumnya membantu organisasi untuk mengelola infrastruktur dengan lebih efektif, meningkatkan kinerja aplikasi, dan meningkatkan ketangguhan infrastruktur terhadap kegagalan.
