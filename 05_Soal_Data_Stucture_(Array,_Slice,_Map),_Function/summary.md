# Data Structure

STRUCTURE DATA (ARRAY, SLICE, & MAP)

## Array
*Array* adalah struktur data yang berisi sekelompok elemen, dapat berisi satu jenis variabel dengan ukuran alokasi yang tetap. Berbagai jenis data yang berbeda dapat ditangani sebagai elemen dalam array seperti numerik, string, boolean.

*Deklarasi array -* untuk mendeklarasikan sebuah array perlu menentukan jumlah elemen yang di dalamnya di dalam tanda kurung siku [], diikuti oleh tipe elemen yang dipegang oleh array tersebut.

## Slice
*Slice* adalah struktur data yang berisi sekelompok elemen, dapat berisi satu jenis variabel (seperti array) tetapi memiliki ukuran alokasi dinamis.

Slice sebenarnya bukanlah array dinamis, melainkan tipe data referensi. Slice dideklarasikan seperti array, namun tidak perlu menentukan ukuran dalam tanda kurung siku [].

*Append() :*  meningkatkan kapasitas (capacity) dari sebuah slice.

*Copy() :* Isi dari slice sumber akan disalin ke dalam slice tujuan.


## Map
*Map* adalah struktur data yang menyimpan data dalam bentuk pasangan kunci dan nilai, di mana setiap kunci bersifat unik. 

| Keys | Value |
| ----------- | ----------- |
| November | 30 |
| Desember | 31 |

## Function
*Fungsi* adalah potongan kode yang dipanggil dengan nama. Fungsi merupakan cara yang nyaman untuk membagi kode Anda menjadi blok-blok yang berguna. Ini memungkinkan kita untuk menulis kode yang bersih, rapi, dan modular.

Struktur Data: Array, Slice, dan Map

Array:

Definisi:

Array adalah struktur data statis yang menyimpan elemen-elemen dengan tipe data yang sama. Ukuran array biasanya tetap dan harus ditentukan saat deklarasi.
Karakteristik:

Ukuran tetap, tidak dapat diubah setelah deklarasi.
Akses elemen menggunakan indeks.
Alamat memori yang berurutan untuk setiap elemen.
Kelebihan:

Akses elemen cepat dengan indeks.
Penyimpanan elemen secara kontinu memudahkan pengelolaan memori.
Kekurangan:

Ukuran tetap, sulit disesuaikan dengan perubahan dinamis.
Memerlukan pengalokasian memori statis.
Slice:

Definisi:

Slice adalah struktur data dinamis yang memungkinkan Anda mengelola potongan atau bagian dari array. Slices lebih fleksibel karena ukurannya dapat berubah.
Karakteristik:

Referensi ke bagian dari array.
Ukuran dapat berubah.
Dibuat menggunakan operator slicing a[start:end].
Kelebihan:

Dinamis, dapat disesuaikan dengan perubahan ukuran.
Lebih fleksibel daripada array.
Kekurangan:

Memerlukan alokasi memori dinamis.
Slightly lebih lambat dalam akses elemen dibandingkan array.
Map (Peta atau Pemetaan):

Definisi:

Map (peta atau pemetaan) adalah struktur data yang memetakan kunci ke nilai. Dikenal juga sebagai associative array, dictionary, atau hash map.
Karakteristik:

Setiap elemen terdiri dari pasangan kunci-nilai.
Kunci harus unik di dalam map.
Tidak memiliki urutan tertentu.
Kelebihan:

Akses data cepat melalui kunci.
Memudahkan pengelolaan hubungan antara data.
Kekurangan:

Tidak memiliki urutan tertentu, tidak cocok untuk pengurutan.
Memerlukan alokasi memori untuk setiap elemen.
Hal-Hal Penting:

Kompleksitas Waktu:

Array dan Slice: O(1) untuk mengakses elemen dengan indeks.
Map: Rata-rata O(1) untuk operasi penyimpanan dan pengambilan data.
Operasi Dasar:

Array: Akses, ubah nilai, inisialisasi.
Slice: Slicing, append, copy, len.
Map: Insert, delete, retrieve, check keberadaan kunci.
Penggunaan yang Tepat:

Gunakan array jika ukuran tetap dan diketahui sebelumnya.
Gunakan slice untuk dinamisitas dan potongan data dari array.
Gunakan map untuk hubungan antara kunci dan nilai.
Pentingnya Hashing pada Map:

Hashing digunakan untuk mengonversi kunci menjadi indeks untuk akses cepat.
Dinamika Pertumbuhan:

Slice lebih dinamis dan fleksibel dibandingkan array.
Map dapat dengan mudah berkembang seiring waktu tanpa perlu menentukan ukuran sebelumnya.
Pentingnya Keseragaman Tipe Data:

Array harus memiliki elemen dengan tipe data yang sama.
Map dapat memiliki kunci dan nilai dengan tipe data yang berbeda.
Pemahaman yang baik tentang array, slice, dan map membantu dalam desain dan implementasi program, memilih struktur data yang paling sesuai untuk tugas tertentu, dan memaksimalkan efisiensi dan fleksibilitas kode.
