# Recursive - Number Theory - Sorting - Searching

ALGORITHM AND DATA STRUCTURE

## Recursive
Yaitu dimana sebuah fungsi dapat memanggil dirinya sendiri untuk menyelesaikan suatu masalah. Jika terdapat masalah yang ringan, maka hasil akan didapat dengan cepat. Namun untuk masalah yang kompleks dapat dipecah menjadi masalah yang lebih kecil yang serupa. Setiap pemanggilan rekursif akan memproses masalah yang lebih kecil hingga akhirnya mencapai kondisi dasar yang disebut sebagai "base case". Base case merupakan kondisi di mana rekursi berhenti dan nilai dikembalikan tanpa adanya pemanggilan rekursif lebih lanjut.

## Number Theory
Number theory atau teori bilangan adalah cabang dalam matematika yang mempelajari sifat dan hubungan antara bilangan bulat. Subjek ini mencakup berbagai konsep dan teorema terkait bilangan bulat, termasuk sifat-sifat bilangan prima, faktorisasi, sifat-sifat divisibilitas, masalah-masalah aritmetika, dan banyak lagi.

## Searching
Searching atau pencarian adalah proses mencari suatu elemen tertentu dalam suatu kumpulan data atau struktur data. Tujuannya adalah untuk menemukan apakah elemen yang dicari ada dalam kumpulan data tersebut dan jika iya, di mana posisinya. Pencarian adalah salah satu operasi dasar dalam pemrosesan data dan sering digunakan dalam pemrograman untuk menemukan informasi tertentu dari data yang tersimpan.

Berikut beberapa algoritma pencarian yang umum digunakan:

1. *Linear Search* : Algoritma ini melakukan pencarian elemen dengan cara memeriksa satu per satu elemen dalam urutan, mulai dari awal hingga akhir. 
2. *Binary Search* : Hanya berlaku untuk data yang telah diurutkan secara teratur (biasanya dalam urutan menaik). Binary search membagi kumpulan data menjadi setengah bagian setiap langkah, membandingkan elemen tengah dengan elemen yang dicari, dan kemudian mengabaikan salah satu setengahnya berdasarkan hasil perbandingan. 

## Sorting
Sorting atau pengurutan adalah proses mengatur kumpulan data dalam urutan tertentu, baik itu menaik (dari nilai terkecil ke terbesar) atau menurun (dari nilai terbesar ke terkecil). Tujuan utama pengurutan adalah membuat data lebih mudah dicari, diproses, dan dipahami.

Berikut beberapa algoritma pengurutan yang umum digunakan:

1. *Selection Sort* : Algoritma yang bekerja dengan memilih elemen minimum (atau maksimum) dari sisa data dan menukarnya dengan elemen pertama, kemudian memilih elemen minimum dari sisa data dan menukarnya dengan elemen kedua, dan seterusnya.
2. *Counting Sort* : Algoritma yang bekerja dengan menghitung berapa kali setiap elemen muncul, membangun array kumulatif dari hitungan tersebut, dan kemudian memposisikan setiap elemen di tempat yang benar.
3. *Merge Sort* : Algoritma pengurutan yang menggunakan pendekatan rekursif dengan membagi data menjadi dua bagian, mengurutkan setiap bagian, dan kemudian menggabungkannya kembali menjadi satu data yang terurut.

Algorithm and Data Structure (Algoritma dan Struktur Data):

Algorithm (Algoritma):

Divide and Conquer:

Algoritma ini memecah masalah besar menjadi submasalah yang lebih kecil, menyelesaikan masing-masing secara terpisah, dan kemudian menggabungkan solusi submasalah tersebut.
Greedy Algorithm:

Algoritma Greedy membuat keputusan lokal optimal pada setiap langkah dalam harapan bahwa keputusan ini akan menghasilkan solusi optimal secara global.
Dynamic Programming:

Menerapkan pemrograman dinamis untuk memecahkan masalah dengan memecahnya menjadi submasalah yang lebih kecil dan menyimpan solusi submasalah untuk menghindari perhitungan berulang.
Backtracking:

Strategi pencarian yang mencoba menemukan solusi melalui serangkaian keputusan, dan jika solusi tidak memuaskan, kembali ke keputusan sebelumnya dan mencoba alternatif lain.
Data Structure (Struktur Data):

Linked List (Daftar Berantai):

Singly Linked List:

Setiap elemen memiliki dua bagian: data dan pointer ke elemen berikutnya.
Doubly Linked List:

Setiap elemen memiliki dua pointer: satu ke elemen berikutnya dan satu ke elemen sebelumnya.
Circular Linked List:

Elemen terakhir menunjuk kembali ke elemen pertama, membentuk lingkaran.
Tree (Pohon):

Binary Tree:

Setiap node memiliki maksimal dua anak, yaitu anak kiri dan anak kanan.
Binary Search Tree (BST):

Binary tree di mana nilai di node kiri lebih kecil daripada nilai di node tersebut, dan nilai di node kanan lebih besar.
AVL Tree:

Bentuk pohon biner pencarian yang dijaga dalam keseimbangan dengan mengatur tinggi subpohon dengan perbedaan tinggi maksimum satu.
Heap:

Max Heap:

Heap di mana setiap node anak memiliki nilai kurang dari atau sama dengan nilai orang tuanya, sehingga nilai maksimum berada di root.
Min Heap:

Heap di mana setiap node anak memiliki nilai lebih besar dari atau sama dengan nilai orang tuanya, sehingga nilai minimum berada di root.
Graph (Graf):

Directed Graph:

Graf di mana setiap tautan atau sisi memiliki arah tertentu, dari satu node ke node lain.
Undirected Graph:

Graf di mana tautan atau sisi tidak memiliki arah tertentu.
Weighted Graph:

Graf di mana setiap tautan atau sisi memiliki bobot atau nilai terkait.
Hash Table:

Collision Handling:
Teknik untuk menangani tumpang tindih, seperti chaining (menempatkan elemen yang bertumpuk dalam daftar terkait) atau open addressing (mencari lokasi kosong terdekat).
Queue:

Priority Queue:
Antrian di mana setiap elemen memiliki prioritas, dan elemen dengan prioritas tertinggi diambil terlebih dahulu.
Stack:

Application of Stack:
Digunakan dalam evaluasi ekspresi matematika, pemanggilan fungsi, manajemen memori, dan pembalikan string.
Setiap struktur data memiliki karakteristik, kelebihan, dan kekurangan tertentu yang membuatnya sesuai untuk penggunaan tertentu. Pemahaman yang mendalam tentang algoritma dan struktur data memungkinkan pengembang untuk memilih pendekatan terbaik dalam merancang dan mengoptimalkan solusi perangkat lunak.
