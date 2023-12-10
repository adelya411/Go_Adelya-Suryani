# Time Compexity & Space Complexity

TIME COMPLEXITY AND SPACE COMPLEXITY

*Time Complexity*

Penggunaan kompleksitas waktu memudahkan estimasi waktu eksekusi program dengan melihat jumlah maksimum operasi primitif yang mungkin dilakukan. Operasi primitif termasuk penambahan, perkalian, penugasan, dan lainnya. Operasi yang paling sering dilakukan disebut operasi dominan, sementara beberapa operasi lainnya bisa diabaikan.

*Comparison of Different Time Complexities*
1. Constant time - O(1), jumlah operasi yang tetap.
2. Linear time - O(n), Jika nilai pertama dari larik A adalah 0, program akan segera berakhir. Namun, dalam menganalisis kompleksitas waktu, fokus pada kasus terburuk diperlukan. Program akan memakan waktu paling lama untuk dieksekusi ketika larik A tidak mengandung angka 0.
3. Linear time - O(n + m), Waktu linear O(n+m) mengindikasikan kompleksitas waktu yang tumbuh sejalan dengan total elemen dalam dua input yang berbeda, yaitu n dan m.
4. Logarithmic time - O(log n), Program ini memiliki loop di mana nilai n dibagi dua setiap iterasinya. Jika n awalnya adalah 2^x, maka jumlah iterasi yang diperlukan adalah x (di mana log n = x). Waktu eksekusi program tergantung pada nilai n yang diberikan sebagai masukan.
5. Quadratic time - O(n^2), Ketika menghitung kompleksitas, fokus pada istilah yang tumbuh paling cepat, jadi abaikan konstanta dan istilah lainnya. Hasilnya adalah kompleksitas waktu kuadrat. Terkadang, kompleksitas juga bergantung pada lebih banyak variabel.

 *Exponential and Factorial Time*
 
 Ada jenis kompleksitas waktu lain, seperti kompleksitas faktorial O(n!) dan kompleksitas eksponensial O(2^n). Dimanamhanya cocok untuk masalah dengan nilai n yang sangat kecil karena terlalu lambat jika digunakan untuk nilai n yang besar.

*Space Complexity*

Batasan memori memberikan info tentang kompleksitas ruang yang diharapkan. Perkirakan jumlah variabel yang boleh dideklarasikan. Singkatnya, jumlah variabel tetap berarti kompleksitas ruang tetap.

Kompleksitas Waktu (Time Complexity):

Definisi:

Time complexity adalah ukuran seberapa cepat waktu eksekusi suatu algoritma tumbuh seiring dengan ukuran inputnya.
Notasi Besar-O (Big O Notation):

Representasi sintetis dari kompleksitas waktu yang memberikan batas atas pertumbuhan fungsi waktu algoritma.
Contoh Notasi Besar-O:

O(1): Konstanta, waktu eksekusi tetap konstan terlepas dari ukuran input.
O(log n): Logaritmik, pertumbuhan waktu yang lebih lambat seiring dengan pertumbuhan logaritma dari ukuran input.
O(n): Linier, pertumbuhan waktu sebanding dengan ukuran input.
O(n log n): Superlinier, umumnya terjadi dalam algoritma pengurutan cepat.
O(n^2): Kuadratik, pertumbuhan waktu sebanding dengan kuadrat dari ukuran input.
Analisis Kasus Terbaik, Rata-rata, dan Terburuk:

Terbaik: Kompleksitas waktu saat input memberikan hasil terbaik.
Rata-rata: Kompleksitas waktu rata-rata untuk semua kemungkinan input.
Terburuk: Kompleksitas waktu saat input memberikan hasil terburuk.
Pentingnya Kompleksitas Waktu:

Membantu memilih algoritma yang efisien.
Memahami bagaimana performa algoritma berubah dengan ukuran input.
Kompleksitas Ruang (Space Complexity):

Definisi:

Space complexity adalah ukuran seberapa banyak ruang memori yang dibutuhkan oleh suatu algoritma seiring dengan pertumbuhan ukuran inputnya.
Notasi Besar-O untuk Ruang (Space):

Sama dengan kompleksitas waktu, tetapi mengukur kebutuhan ruang memori algoritma.
Contoh Notasi Besar-O untuk Ruang:

O(1): Kebutuhan ruang konstan, tidak tergantung pada ukuran input.
O(n): Kebutuhan ruang tumbuh sebanding dengan ukuran input.
O(n^2): Kebutuhan ruang tumbuh sebanding dengan kuadrat dari ukuran input.
Overhead dan Penggunaan Memori Tambahan:

Overhead adalah penggunaan memori tetap yang tidak tergantung pada ukuran input.
Penggunaan memori tambahan dapat terjadi karena struktur data tambahan atau variabel sementara.
Pentingnya Kompleksitas Ruang:

Membantu memilih algoritma yang meminimalkan penggunaan memori.
Penting dalam lingkungan dengan batasan memori.
Hal-hal Terkait:

Trade-off Antara Waktu dan Ruang:

Terkadang, algoritma yang lebih cepat menggunakan lebih banyak ruang dan sebaliknya. Perlu memilih berdasarkan kebutuhan spesifik.
Optimisasi dan Pengembangan Algoritma:

Pengoptimalan dapat mencakup meningkatkan kompleksitas waktu atau ruang, tergantung pada kebutuhan aplikasi.
Analisis Empiris vs. Analisis Teoritis:

Analisis teoritis menggunakan notasi besar-O, sementara analisis empiris melibatkan pengujian algoritma pada input yang sebenarnya.
Pentingnya Pemahaman Asimtotik:

Asimtotik adalah studi tentang pertumbuhan fungsi saat input mendekati tak hingga. Pemahaman ini penting untuk memilih algoritma yang sesuai.
Penerapan di Kehidupan Nyata:

Kompleksitas waktu dan ruang memainkan peran kunci dalam desain algoritma dan pengembangan perangkat lunak di dunia nyata.
Pemahaman tentang kompleksitas waktu dan ruang adalah kunci untuk menghasilkan algoritma yang efisien dan memahami bagaimana kinerja suatu program berubah seiring dengan ukuran dan sifat inputnya.
