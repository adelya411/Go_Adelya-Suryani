# Concurrency Programming

Concurrency in Golang


Concurrency programming yaitu menjalankan tugas secara bersamaan dan secara independent (berdiri sendiri) 

Feature Go :
1. Concurrent execution (Goroutines)
2. Synchronization and messaging (Channels)
3. Multi-way concurrent control (Select)

## Goroutines
Sebuah proses concurrent yang dijalankan di dalam golang. Goroutines adalah sebuah function yang ditentukan akan berjalan secara concurrently dibandingkan dengan function yang lain.

Gomaxprocs : Ini digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi goroutine dalam program Go tertentu.

## Channel & Select
Sebuah channel adalah objek komunikasi yang digunakan oleh goroutine untuk berkomunikasi satu sama lain.

Channel dalam Go dapat dibagi menjadi dua jenis:
1. *Buffered channel,  terdapat kapasitas tertentu untuk menyimpan pesan sebelum penerimaannya oleh goroutine penerima. Ketika buffer penuh, pengirim pesan tidak perlu menunggu penerima. Mereka dapat melanjutkan eksekusi mereka. Hanya ketika buffer penuh, pengiriman pesan akan memblokir pengirim sampai ada ruang di buffer untuk pesan baru. **Contoh:* "ch := make(chan int, 10)" akan membuat buffered channel dengan kapasitas buffer 10.
2. *Unbuffered channel,  setiap pengiriman pesan dari satu goroutine ke goroutine lainnya akan mengharuskan pengirim menunggu hingga penerima siap untuk menerima pesan tersebut. Ini memastikan bahwa pesan yang dikirim akan segera diterima oleh penerima sebelum goroutine pengirim melanjutkan eksekusi selanjutnya. **Contoh:* "ch := make(chan int)".

*Select* adalah pernyataan yang digunakan dalam Go untuk mengendalikan komunikasi data melalui satu atau banyak channel.

## Race Condition
Race condition (kondisi balapan) terjadi ketika dua thread mengakses memori pada waktu yang bersamaan, di mana salah satu dari mereka sedang menulis. Race condition terjadi karena akses yang tidak disinkronkan ke memori bersama.

Fixing Data Race :
1. WaitGroups : Cara yang paling langsung untuk mengatasi data race adalah dengan memblokir akses baca hingga operasi penulisan telah selesai.
2. Channels Blocking : Ini memungkinkan goroutine-goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sesaat.
3. Mutex : Jika kita tidak peduli dengan urutan baca dan tulis, yang penting adalah bahwa mereka tidak terjadi secara bersamaan, maka kita sebaiknya mempertimbangkan penggunaan mutex.

1. Goroutines:

Konsep Utama: Goroutines adalah fungsi yang dijalankan secara bersamaan dengan fungsi lainnya. Mereka sangat ringan dan dapat dibuat dengan menggunakan kata kunci go.
Contoh:      func main() {
    go myFunction() // Membuat goroutine
    // ...
}
             2. Channels:

Konsep Utama: Channels adalah sarana komunikasi antara goroutines. Mereka memungkinkan pengiriman data secara aman antar goroutines.
Contoh:           ch := make(chan int)

go func() {
    ch <- 42 // Mengirim data ke channel
}()

value := <-ch // Menerima data dari channel
  3. Select Statement:

Konsep Utama: Select statement digunakan untuk memilih satu dari beberapa operasi komunikasi yang siap dieksekusi.
Contoh:        select {
case msg1 := <-ch1:
    fmt.Println("Received", msg1)
case ch2 <- 42:
    fmt.Println("Sent 42 to ch2")
default:
    fmt.Println("No communication ready")
}
    4. Mutex (Mutual Exclusion):

Konsep Utama: Mutex digunakan untuk mengamankan akses ke bagian-bagian kode yang kritis, mencegah race condition.
Contoh:        var mu sync.Mutex

func myFunction() {
    mu.Lock()
    defer mu.Unlock()
    // Critical section
}
      5. Wait Groups:

Konsep Utama: Wait groups digunakan untuk menunggu selesainya sekelompok goroutines sebelum melanjutkan eksekusi.
Contoh:            var wg sync.WaitGroup

func main() {
    wg.Add(1)
    go myFunction()
    wg.Wait() // Menunggu selesai semua goroutines
}
               6. Atomic Operations:

Konsep Utama: Package atomic digunakan untuk melakukan operasi atomik, yang dapat dilakukan tanpa khawatir tentang race condition.
Contoh:           var counter int32

func incrementCounter() {
    atomic.AddInt32(&counter, 1)
}
   7. Context:

Konsep Utama: Package context digunakan untuk membatalkan, mengganti nilai timeout, atau menetapkan nilai-nilai khusus untuk goroutines.
Contoh:             ctx, cancel := context.WithCancel(context.Background())

go func() {
    // ...
    cancel() // Membatalkan konteks
}()
      8. Rate Limiting:

Konsep Utama: Golang menyediakan package golang.org/x/time/rate untuk mengimplementasikan rate limiting, yang dapat digunakan untuk mengendalikan frekuensi panggilan fungsi atau akses ke sumber daya tertentu.
Contoh:            limiter := rate.NewLimiter(rate.Limit(1), 5)

func myFunction() {
    if limiter.Allow() {
        // Do something
    }
}
   9. Context Switching:

Konsep Utama: Golang dirancang untuk mengelola context switching dengan efisien melalui goroutines, yang dapat berpindah konteks dengan cepat.
Contoh:           go func() {
    // ...
    runtime.Gosched() // Menyerahkan waktu eksekusi ke goroutine lain
}()
        10. Fan-out, Fan-in:

Konsep Utama: Fan-out adalah memulai beberapa goroutines untuk melakukan tugas yang sama, dan fan-in adalah menggabungkan hasil dari goroutines tersebut.
Contoh:                func fanOut(chIn <-chan int, chOut []chan int) {
    go func() {
        for v := range chIn {
            for _, c := range chOut {
                c <- v
            }
        }
    }()
}
     Concurrency programming in Golang memberikan kemampuan untuk menangani tugas-tugas konkuren secara efisien, membuat aplikasi yang scalable, dan mengoptimalkan penggunaan sumber daya. Penggunaan goroutines, channels, dan fitur-fitur lainnya menciptakan paradigma pemrograman konkuren yang kuat di dalam bahasa pemrograman Golang.

     