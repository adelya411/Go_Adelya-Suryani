# Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

PROBLEM SOLVING PARADIGM

## Problem Solving Paradigm
Bagaimana menyelesaikan sebuah masalah. Terdapat jenis-jenis untuk menyelesaikan sebuah masalah, yaitu :
1. *Complete Search* atau yang dikenal dengan brute force dimana cara penyelesaian masalahnya dengan melakukan sebuah pengecekan pada semua data. Lebih baik digunakan ketika tidak ada algoritma yang lebih cepat digunakan.
2. *Divide and Conquer* dimana ketika terdapat sebuah masalah yang besar maka akan dibagi menjadi masalah yang lebih kecil dan langsung diselesaikan. Contohnya pada penggunaan binary search.
3. *Greedy* yaitu sebuah algoritma yang menyelesaikan masalah dengan mencari nilai lokal optimal. Algoritma greedy yang lain yaitu huffman coding, activity selection, dijkstra algorithm, dan lain-lain.
4. *Dynamic Programming* adalah algoritma teknik untuk menyelesaikan sebuah problem secara optimal dengan memecahkan problem tersebut ke dalam problem-problem yang lebih mudah dikerjakan akan tetapi harus mengoptimalkan fakta optimalnya seperti apa dari setiap problem tersebut. Karakteristiknya yaitu :
- Overlapping Subproblems, ketika mencari solusinya, subproblems akan dipanggil beberapa kali. 
- Optimal Substructure Property, menyatakan bahwa solusi optimal untuk masalah utama dapat ditemukan dengan menggabungkan solusi optimal untuk submasalah yang lebih kecil.
Methodnya yaitu :
- Top-down with Memoization, adalah pendekatan yang menggunakan rekursi dengan penyimpanan (caching) hasil dari submasalah yang sudah dipecahkan. 
- Bottom-up with Tabulation, metode ini melibatkan pembuatan tabel dan mengisi seluruh tabel dengan hasil dari submasalah dalam urutan yang tepat, mulai dari submasalah terkecil hingga masalah utama.

