# String - Advance Function - Pointer - Method - Struct and Interface

STRING, ADVANCE FUNCTION, POINTER, STRUCT, METHOD & INTERFACE

## String
*String* adalah tipe data yang digunakan untuk merepresentasikan urutan karakter dalam pemrograman.

*Menggunakan String*
1. len: Mengukur panjang (jumlah karakter) dari sebuah string.
2. compare: Membandingkan dua string untuk menentukan urutan leksikal mereka.
3. contains: Memeriksa apakah sebuah string mengandung substring tertentu.
4. substrings: Membuat substring baru dari string yang ada.
5. replace: Mengganti kemunculan substring tertentu dengan substring lainnya dalam string.
6. insert: Menyisipkan sebuah substring ke dalam posisi tertentu dalam string.

## Function
*Variadic Function:*
- Untuk menghindari membuat slice sementara hanya untuk dilewatkan ke sebuah fungsi.
- Ketika jumlah parameter input tidak diketahui.
- Untuk menyatakan niat Anda agar kode lebih mudah dibaca.

*Closure* adalah jenis khusus dari fungsi tanpa nama yang merujuk pada variabel yang dideklarasikan di luar fungsi itu sendiri. Dalam kasus ini, kita akan menggunakan variabel yang tidak diteruskan ke dalam fungsi sebagai parameter, tetapi sebaliknya variabel tersebut tersedia ketika fungsi dideklarasikan.

*Fungsi defer*, sebuah fungsi hanya dieksekusi setelah fungsi induknya selesai. Penggunaan return ganda juga dapat digunakan, mereka dijalankan sebagai tumpukan, satu per satu.

## Pointer
*Pointer* adalah variabel yang menyimpan alamat memori dari variabel lain. Pointer memiliki kemampuan untuk mengubah data yang mereka tunjuk.

1. Pointer pada Tipe Data Biasa: Merujuk pada alamat memori variabel dari tipe data yang sama. Contohnya, int pointer merujuk pada alamat memori variabel int, string pointer merujuk pada alamat memori variabel string, dan seterusnya.
2. Pointer Struct: Anda dapat menggunakan pointer pada tipe data struct untuk merujuk pada instansi struktur.
3. Pointer ke Pointer: Anda bisa memiliki pointer yang merujuk pada pointer lain. Ini berguna ketika Anda ingin membuat perubahan pada variabel melalui beberapa tingkat pointer.
4. Pointer ke Fungsi: Anda bisa memiliki pointer yang merujuk pada fungsi. Ini memungkinkan Anda untuk meneruskan fungsi sebagai argumen ke fungsi lain.

## Method
*Method* adalah sebuah fungsi yang dilampirkan pada suatu tipe data (dapat berupa struktur atau tipe data lainnya).

Deklarasi Method mirip dengan fungsi, hanya saja deklarasi variabel objek perlu ditambahkan di antara kata kunci func dan nama fungsi.

Mengapa Menggunakan Method Daripada Fungsi?
1. Membantu Anda Menulis Kode Berorientasi Objek di Go.
2. Method Membantu Menghindari Konflik Nama.
3. Pemanggilan Method Lebih Mudah Dibaca dan Dimengerti Daripada Pemanggilan Fungsi.

## Interface
*Interface* adalah kumpulan dari tanda tangan method yang dapat diimplementasikan oleh suatu objek. Oleh karena itu, interface mendefinisikan perilaku dari objek tersebut.

*Package* adalah kumpulan dari fungsi-fungsi dan data.

Ketika runtime Go mendeteksi kesalahan-kesalahan ini, itu memicu *panic. Untuk menambahkan kemampuan **recorver* dari error panic, Anda dapat menambahkan sebuah fungsi tanpa nama (anonymous function) atau mendefinisikan sebuah fungsi kustom dan memanggilnya dengan kata kunci "defer" dari dalam method, di mana panic mungkin terjadi dari pemanggilan internal lainnya.

String:

Mutability:

Sebagian besar bahasa pemrograman memperlakukan string sebagai tipe data yang tidak dapat diubah (immutable). Ini berarti setelah string dibuat, Anda tidak dapat mengubah karakter di dalamnya. Operasi seperti concatenation membuat string baru.
Encoding:

Karakter dalam string direpresentasikan oleh kode karakter. Encoding, seperti UTF-8, UTF-16, dan ASCII, menentukan cara representasi ini di dalam komputer.
Regex (Regular Expressions):

Regex adalah alat kuat untuk memanipulasi string dengan mencocokkan pola tertentu. Memahami ekspresi reguler penting untuk pencarian dan manipulasi string yang kompleks.
Advanced Functions:

First-Class Functions:

Fungsi dianggap sebagai entitas first-class, yang berarti dapat dioperasikan seperti tipe data lainnya. Ini memungkinkan penggunaan fungsi sebagai parameter, nilai kembali dari fungsi, dan penyimpanan fungsi dalam struktur data.
Higher-Order Functions:

Fungsi yang dapat menerima fungsi lain sebagai argumen atau mengembalikan fungsi sebagai nilai kembali disebut higher-order functions. Ini memungkinkan paradigma pemrograman fungsional.
Anonymous Functions (Lambda):

Anonymous functions, atau lambda functions, memungkinkan definisi fungsi tanpa harus memberikan nama. Berguna dalam situasi di mana fungsi kecil diperlukan secara sementara.
Pointer:

Pointer Arithmetic:

Pointer arithmetic melibatkan penambahan atau pengurangan nilai pointer. Ini khas dalam bahasa pemrograman C dan C++, di mana pointer dapat digunakan untuk berinteraksi dengan array dan struktur data secara langsung.
Void Pointer:

Void pointer adalah pointer yang tidak memiliki tipe data yang terkait dengannya. Dapat digunakan untuk menangani alamat memori tanpa perlu mengetahui tipe data sebenarnya.
Double Pointer:

Double pointer adalah pointer yang menunjuk ke pointer lain. Berguna dalam situasi di mana kita ingin mengubah nilai dari pointer melalui fungsi.
Struct:

Nested Struct:

Struct dapat memiliki anggota yang juga merupakan struct. Ini disebut nested struct dan memungkinkan penyusunan data yang lebih kompleks dan terorganisir.
Bitfields:

Bitfields memungkinkan penggunaan sejumlah bit yang ditentukan dalam struktur data. Berguna untuk menghemat ruang memori dalam representasi data tertentu.
Method:

Receiver Types:

Method dalam pemrograman berorientasi objek dapat memiliki receiver types. Receiver types menunjukkan tipe data yang bisa memiliki metode tersebut. Ada receiver types by value dan by pointer.
Method Chaining:

Method chaining melibatkan pemanggilan serangkaian metode secara berurutan pada objek yang sama. Ini meningkatkan ekspresivitas dan membantu membuat kode lebih ringkas.
Interface:

Type Assertion:

Dalam Go dan beberapa bahasa pemrograman lainnya, untuk menggunakan nilai dari antarmuka, type assertion digunakan. Ini memungkinkan pengambilan nilai konkret dari antarmuka.
Empty Interface:

Empty interface, antarmuka tanpa metode, dikenal sebagai interface{}. Dapat menyimpan nilai dari tipe apa pun dan berguna dalam situasi di mana tipe tidak diketahui sebelumnya.
Type Switch:

Type switch adalah mekanisme dalam bahasa pemrograman seperti Go yang memungkinkan pengujian tipe untuk menentukan tipe konkret dari antarmuka atau nilai interface kosong.
Setiap konsep yang lebih lanjut ini membuka pintu untuk fungsionalitas dan paradigma pemrograman yang lebih canggih. Pemahaman mendalam tentang string, fungsi tingkat tinggi, pointer, struct, metode, dan antarmuka dapat meningkatkan kemampuan pemrograman Anda secara signifikan.
