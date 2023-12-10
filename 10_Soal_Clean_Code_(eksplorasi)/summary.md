# Clean Code

Clean Code

Clean code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah  oleh programmer. 

Karakteristik:
1. Mudah dipahami
2. Mudah dieja dan dicari
3. Singkat namun mendeskripsikan konteks
4. Konsisten
5. Hindari penambahan konteks yang tidak perlu
6. Komentar
7. Good function
8. Gunakan konvensi
9. Formatting

Principle:
1. KISS (Keep It So Simple), Hindari membuat fungsi yang dibuat untuk melakukan A, sekalian memodifikasi B, mengecek fungsi C, dst.
2. DRY (Don't Repeat Yourself), Duplikasi code terjadii karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

Refactoring, adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal.Prinsip KISS dan DRY bisa dicapai dengan cara refactor.

Teknik refactoring:

- Membuat sebuah abstraksi
- Memcah kode dengan fungsi/class
- Perbaiki penamaan dan lokasi kode
- Deteksi kode yang memiliki duplikasi

Clean Code:

Clean Code adalah konsep pengkodean yang bertujuan untuk meningkatkan kualitas dan kejelasan kode sumber. Hal ini mencakup sejumlah prinsip dan praktik untuk membuat kode lebih mudah dipahami, mudah dikelola, dan lebih mudah dipelihara.

1. DRY (Don't Repeat Yourself):

Konsep Utama: Hindari pengulangan kode. Jika Anda menemukan diri Anda menulis logika yang sama atau serupa di beberapa tempat, pindahkan ke dalam fungsi atau kelas yang sesuai.
2. KISS (Keep It Simple, Stupid):

Konsep Utama: Usahakan untuk membuat kode sejelas mungkin dan hindari kompleksitas yang tidak perlu. Solusi yang sederhana lebih mudah dipahami dan lebih mudah dipelihara.
3. YAGNI (You Ain't Gonna Need It):

Konsep Utama: Hindari menambahkan fitur atau fungsi yang tidak diperlukan atau belum diperlukan. Fokuslah pada kebutuhan saat ini, bukan spekulasi kebutuhan di masa depan.
4. SOLID Principles:

Single Responsibility Principle (SRP): Sebuah kelas harus hanya memiliki satu alasan untuk berubah.
Open/Closed Principle (OCP): Kode harus terbuka untuk ekstensi, tetapi tertutup untuk modifikasi.
Liskov Substitution Principle (LSP): Subtipe harus dapat diganti untuk tipe induknya tanpa mengubah kebenaran program.
Interface Segregation Principle (ISP): Klien tidak boleh dipaksa untuk mengimplementasikan metode yang tidak digunakan.
Dependency Inversion Principle (DIP): Tergantung pada abstraksi, bukan implementasi.
5. Code Smells:

Konsep Utama: Pahami tanda-tanda yang menunjukkan kemungkinan masalah di dalam kode. Contoh code smells termasuk metode yang terlalu panjang, kelas yang terlalu besar, dan nama variabel yang tidak deskriptif.
6. Meaningful Names:

Konsep Utama: Gunakan nama variabel, fungsi, dan kelas yang jelas dan mendeskripsikan tujuan atau fungsi mereka. Hindari singkatan yang tidak jelas.
7. Comments:

Konsep Utama: Gunakan komentar dengan bijak dan hanya jika diperlukan. Kode yang baik seharusnya dapat menjelaskan dirinya sendiri. Hindari komentar yang menjelaskan "apa" daripada "mengapa".
8. Code Formatting:

Konsep Utama: Memiliki format kode yang konsisten dan mudah dibaca. Gunakan aturan indentasi dan spasi yang konsisten di seluruh proyek.
9. Unit Testing:

Konsep Utama: Tulis unit test yang menyeluruh dan relevan. Ini membantu memastikan bahwa kode berfungsi sebagaimana mestinya dan memudahkan perubahan tanpa menimbulkan efek samping yang tidak diinginkan.
10. Refactoring:

Konsep Utama: Terus-menerus tingkatkan kualitas kode dengan melakukan refactoring. Perbaiki code smells, hilangkan redundansi, dan buat kode lebih bersih seiring waktu.
11. Consistency:

Konsep Utama: Pertahankan konsistensi dalam seluruh proyek. Gunakan gaya penulisan yang sama, patuhi konvensi penamaan, dan terapkan standar pengkodean yang telah ditetapkan.
Clean Code adalah lebih dari sekadar aturan-aturan. Ini mencerminkan filosofi bahwa kode harus mudah dipahami oleh programmer lain dan diri sendiri, serta dapat beradaptasi dengan perubahan tanpa menimbulkan masalah. Menerapkan prinsip-prinsip Clean Code membantu membangun perangkat lunak yang lebih robust, mudah dipelihara, dan berkelanjutan.