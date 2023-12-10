# Unit Testing

Unit Testing in Golang

# Unit Testing 
## Survey
Apa itu software testing? Yaitu suatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang diperlukan (yaitu cacat) dan untuk mengevaluasi fitur perangkat lunak.

Tujuan testing:
- Mencegah regresi
- Tingkat keyakinan dalam refactoring
- Meningkatkan desain kode
- Dokumentasi kode
- Kode penskalaan tim

Tingkat pengujian:
1. UI, (End To End) menguji interaksi antar keseluruhan melalui antarmuka pengguna.
2. Integration, menguji modul atau subsistem tertentu melalui api.
3. Unit, menguji bagian terkecil yang dapat diuji(operasi logis tunggal) dari suatu aplikasi melalui metode.

### Framework
Framework menyediakan alat dan struktur yang diperlukan untuk menjalankan pengujian secara efisien.

### Structure
2 pola biasa:
1. Pusatkan file pengujian di dalam folder pengujian.
2. Simpan file pengujian bersama dengan file produksi.

### Runner
- alat yang menjalankan file uji
- gunakan mode jam tangan (jika ada perubahan dalam basis kode, secara otomatis jalankan tes lagi, buat tdd lebih efisien)
- pilih pelari yang bisa berlari paling cepat

### Mocking
Kasing uji anda harus mandiri, anda tidak harus menguji:
- modul pihak ke-3
- basis data
- api pihak ke-3 atau sistem eksternal lainnya
- kelas objek yang telah anda uji di tempat lain

### Coverage
Pembuat kode perlu memastikan apakah mereka telah membuat cukup pengujian.
Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan.

Unit Testing in Golang: Rincian Diluar Definisi Umum

Table-Driven Tests:

Definisi Umum: Table-driven tests adalah teknik pengujian di mana serangkaian kasus uji dan hasil yang diharapkan didefinisikan dalam tabel.
Rincian Diluar Hal Umum:
Data-Driven Tests: Menguji fungsi atau metode dengan berbagai input dan ekspektasi hasil.
Test Coverage:

Definisi Umum: Test coverage adalah metrik yang mengukur sejauh mana kode program diuji oleh unit test.
Rincian Diluar Hal Umum:
Code Smells: Mengidentifikasi bagian-bagian kode yang belum diuji dan mungkin memerlukan perhatian ekstra.
Mocking:

Definisi Umum: Mocking melibatkan penggunaan objek palsu atau tiruan untuk menggantikan komponen eksternal dalam pengujian.
Rincian Diluar Hal Umum:
Test Doubles: Termasuk jenis-jenis objek palsu seperti dummy, stub, spy, dan fake.
Subtests and Sub-benchmarks:

Definisi Umum: Subtests dan sub-benchmarks memungkinkan pengelompokan dan penamaan lebih rinci dalam fungsi uji atau uji kinerja.
Rincian Diluar Hal Umum:
Parallel Testing: Menjalankan subtests secara bersamaan untuk mengurangi waktu eksekusi.
Property-Based Testing:

Definisi Umum: Property-based testing adalah teknik di mana properti atau invarian yang seharusnya selalu berlaku diuji dengan input acak.
Rincian Diluar Hal Umum:
Generators: Menghasilkan data uji secara otomatis berdasarkan aturan yang ditentukan.
Benchmarking:

Definisi Umum: Benchmarking digunakan untuk mengukur kinerja suatu fungsi atau bagian dari kode.
Rincian Diluar Hal Umum:
Memory Profiling: Menilai penggunaan memori selama eksekusi benchmark.
Test Tagging:

Definisi Umum: Test tagging melibatkan penambahan metadata atau tag ke unit test untuk mengelompokkan atau menjalankan subset tertentu dari tes.
Rincian Diluar Hal Umum:
Conditional Testing: Menjalankan atau menghindari tes berdasarkan kondisi atau konfigurasi tertentu.
Golden Files:

Definisi Umum: Golden files adalah file referensi yang berisi hasil yang diharapkan dari suatu operasi.
Rincian Diluar Hal Umum:
Updating Golden Files: Menguji dan memperbarui file-file ini saat logika atau perilaku berubah.
Test Fixture Management:

Definisi Umum: Test fixtures adalah kondisi awal atau konteks yang diperlukan untuk menjalankan tes unit atau fungsional.
Rincian Diluar Hal Umum:
Test Setup and Teardown: Menangani persiapan dan pembersihan setiap tes.
Test Coverage Tools:

Definisi Umum: Alat pengukuran cakupan pengujian digunakan untuk memberikan informasi tentang sejauh mana kode telah diuji.
Rincian Diluar Hal Umum:
Integrasi dengan CI/CD: Mengintegrasikan alat cakupan pengujian dengan alur kerja CI/CD.
Parameterized Tests:

Definisi Umum: Parameterized tests memungkinkan pengujian fungsi dengan berbagai input dan ekspektasi.
Rincian Diluar Hal Umum:
Table-Driven Parameterized Tests: Menggabungkan konsep tabel-driven tests dengan parameterized tests.
Error Assertion:

Definisi Umum: Error assertion melibatkan pengujian fungsi atau metode dengan menguji jenis dan nilai dari error yang dihasilkan.
Rincian Diluar Hal Umum:
Error Wrapping: Pengujian ketika error dibungkus atau diperluas dalam fungsi atau lapisan lain.
Memahami dan menerapkan konsep-konsep ini dalam pengujian unit Golang dapat membantu pengembang menciptakan tes yang lebih kuat, efisien, dan efektif dalam memastikan kualitas dan keandalan perangkat lunak yang dibangun dengan Golang.
