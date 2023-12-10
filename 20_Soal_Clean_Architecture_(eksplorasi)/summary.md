# Clean Architecture dan Hexagonal Architecture

## Clean Architecture
Arsitektur yang baik adalah pemisahan perhatian menggunakan lapisan untuk membangun aplikasi yang modular, terukur, dapat dipelihara dan dapat diuji.
#### Benefit clean architecture
1. memiliki standar struktur
2. lebih cepat dalam development
3. mocking depedensi
4. easy switching

#### CA Layer
1. entities layer -> bisnis object
2. use case layer / domain layer/services -> bisnis logic
3. controller/presenter/handler -> user interaction, keberadaan framework
4. drivers-data layer - menerima data, mengambil, menyimpan data.

myapp/
├── entities/
│   └── user.go
├── usecases/
│   └── user_usecase.go
├── interfaces/
│   ├── controllers/
│   │   └── controller.go
│   └── repositories/
│       └── repository.go
├── main.go

## Contex Golang
Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline. Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request). 

Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses
Membuat context kosong.
ctx :=  context.Background()

Parent dan child context
Context menganut konsep parent dan child Artinya, saat kita membuat context, kita bisa membuat child context dari context yang sudah ada
Parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context. fitur yang dimiliki oleh parent maka dimiliki juga oleh childnya
Context merupakan object yang Immutable, artinya setelah Context dibuat, dia tidak bisa diubah lagi.

Contex with value(menambahkan isi contex)
kita bisa menambahkan value ke sebuah root context. Saat menambah value ke context, secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah sama sekali Untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)

 ctx :=  context.Background()
 ctxA = context.WithValue(ctx, key, value)
 ctxB = context.WithValue(ctxA, key, value)

 Clean Architecture: Penjelasan Diluar yang Tidak Umum

Entities, Use Cases, Gateways:

Definisi Umum: Dalam Clean Architecture, inti dari aplikasi terdiri dari tiga elemen utama: Entities, Use Cases, dan Gateways.
Rincian Diluar Hal Umum:
Entities yang Murni: Memastikan bahwa Entities tidak memiliki logika bisnis tambahan dan hanya berfokus pada representasi data.
Use Cases yang Independen: Use Cases seharusnya tidak bergantung pada detail teknis atau infrastruktur.
Dependency Rule:

Definisi Umum: Prinsip ini menyatakan bahwa arah dependensi harus selalu menuju ke arah inti aplikasi, sehingga lapisan luarannya tidak tergantung pada detail implementasi lapisan dalamannya.
Rincian Diluar Hal Umum:
Penanganan Dependensi Tersyarat: Menerapkan kebijakan untuk membatasi penggunaan dependensi yang merugikan arah dependensi.
Interface Adapters:

Definisi Umum: Interface Adapters menghubungkan antara lapisan inti aplikasi dan elemen-elemen luarannya, seperti framework, database, atau antarmuka pengguna.
Rincian Diluar Hal Umum:
Presenter sebagai Adapter: Menggunakan presenter sebagai bagian dari lapisan adapter untuk mengubah data dari Use Cases menjadi format yang sesuai dengan kebutuhan UI.
Frameworks Independency:

Definisi Umum: Clean Architecture menekankan kemandirian dari framework, sehingga perubahan dalam framework tidak mengganggu inti bisnis aplikasi.
Rincian Diluar Hal Umum:
Custom Framework Abstractions: Membuat abstraksi khusus untuk menangani interaksi dengan framework tertentu, membatasi dampak perubahan framework.
Entities vs. Value Objects:

Definisi Umum: Clean Architecture membedakan antara Entities yang memiliki identitas dan Value Objects yang hanya bernilai.
Rincian Diluar Hal Umum:
Immutable Value Objects: Menjamin nilai objek tidak dapat diubah setelah dibuat.
Hexagonal Architecture (Ports and Adapters): Penjelasan Diluar yang Tidak Umum

Hexagonal vs. Onion Architecture:

Definisi Umum: Hexagonal Architecture (juga dikenal sebagai Ports and Adapters) menekankan pada pemisahan inti aplikasi dari detail teknis, dan ini berbeda dengan Onion Architecture yang fokus pada lapisan-lapisan terkonsentrasi di sekitar inti.
Rincian Diluar Hal Umum:
Layer Independence: Mencapai independensi lapisan yang lebih besar, memungkinkan perubahan pada satu lapisan tanpa memengaruhi lainnya.
Ports and Adapters:

Definisi Umum: Hexagonal Architecture membedakan antara "ports" sebagai antarmuka aplikasi dan "adapters" sebagai implementasi dari port tersebut.
Rincian Diluar Hal Umum:
Multiple Adapters for a Port: Memungkinkan adanya beberapa implementasi dari satu port, memberikan fleksibilitas untuk berbagai konteks penggunaan.
Hexagon Shape:

Definisi Umum: Representasi fisik dari Hexagonal Architecture yang menunjukkan inti aplikasi di tengah dengan port di sisi-sisinya.
Rincian Diluar Hal Umum:
Asymmetrical Hexagon: Menyajikan inti aplikasi dengan lebih dari enam sisi untuk menangani ketergantungan yang tidak seimbang.
Domain Events and Notifications:

Definisi Umum: Hexagonal Architecture mendukung penggunaan domain events untuk mengkomunikasikan perubahan dalam inti aplikasi.
Rincian Diluar Hal Umum:
Asynchronous Event Handling: Menggunakan mekanisme asinkron untuk menangani peristiwa, mengurangi ketergantungan pada respon waktu nyata.
Testing Hexagonal Architecture:

Definisi Umum: Testing dalam Hexagonal Architecture melibatkan pengujian port untuk memastikan antarmuka aplikasi berfungsi seperti yang diharapkan.
Rincian Diluar Hal Umum:
Contract Testing: Menguji kontrak antara port dan adapter untuk memastikan kepatuhan.
Dynamic Dependency Injection:

Definisi Umum: Hexagonal Architecture mendorong penggunaan dependency injection untuk menggantikan ketergantungan statis.
Rincian Diluar Hal Umum:
Runtime Dependency Injection: Menyediakan mekanisme untuk mengganti atau menambahkan implementasi port pada waktu eksekusi.
Hexagonal Persistence:

Definisi Umum: Menggunakan prinsip Hexagonal Architecture pada lapisan persistensi untuk memastikan independensi dari detail teknis database.
Rincian Diluar Hal Umum:
Event-Sourcing for Persistence: Menggunakan event-sourcing untuk mengelola perubahan data secara logis, mengurangi ketergantungan pada struktur tabel database.
Pemahaman dan penerapan konsep-konsep ini di luar definisi umumnya membantu dalam mencapai kelebihan dan kelebihan yang dijanjikan oleh Clean Architecture dan Hexagonal Architecture, seperti fleksibilitas, maintainability, dan independensi dari detail teknis.