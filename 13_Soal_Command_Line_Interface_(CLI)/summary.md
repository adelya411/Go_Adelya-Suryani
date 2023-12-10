# Configuration Management and CLI

Basic Command Line Interface

## Command Line
Command line adalah antarmuka berbasis teks yang cepat dan kuat yang digunakan pengembang untuk berkomunikasi secara lebih efektif dan efisien dengan komputer guna menyelesaikan serangkaian tugas yang lebih luas. 

Kenapa menggunakan command line?
1. Kontrol granular dari suatu OS atau aplikasi 
2. Manajemen lebih cepat dari sejumlah besar sistem operasi
3. Kemampuan untuk menyimpan skrip untuk diotomatisasi tugas rutin
4. Menghubungkan pengetahuan untuk penyelesaian masalah

Command line interface :
1. UNIX Shell
2. Command Prompt (MSDOS)
3. Python REPL
4. MySQL CLI client
5. Mongo Shell
6. redis-cli, etc.

## UNIX Shell
Normal User vs Root User
- Normal User = Diizinkan memanipulasi dir/home/username dir only
- Root User = Diizinkan memanipulasi / dir (semua directory)
- Normal User +  Sudoers = Diizinkan bertindak sebagai root sementara. 

Directory :
- pwd 
- ls
- mkdir
- cd
- rm
- cp
- mv
- ln

Files :
- create : touch
- view : head, cat, tail, less
- editor : vim, nano
- permission : chown, chmod
- different : diff

Network :
- ping
- ssh
- netstat
- nmap
- ip addr (ifconfig)
- wget
- curl
- telnet
- netcat

Utility :
- man
- env
- echo
- date
- which
- watch
- sudo
- history
- grep
- locate

Shell, Program yang berfungsi sebagai jembatan antara user dan kernel. Shell script adalah bahasa pemrograman yang dikompilasi berdasarkan perintah shell.

Configuration Management:

1. Infrastructure as Code (IaC):

Definisi Umum: IaC adalah pendekatan untuk mengelola dan mengotomatiskan infrastruktur melalui kode. Ini memungkinkan konfigurasi dan pengelolaan infrastruktur melalui file teks atau skrip.
2. Desired State Configuration (DSC):

Definisi Umum: DSC adalah konsep di mana Anda mendefinisikan keadaan yang diinginkan dari sistem, dan alat konfigurasi akan mengelola sistem untuk mencapai keadaan tersebut.
3. Infrastructure as Configuration (IaC):

Definisi Umum: IaC tidak hanya mencakup mengotomatiskan infrastruktur, tetapi juga mencakup pengelolaan konfigurasi perangkat lunak dan aplikasi di dalam infrastruktur tersebut.
4. Configuration Drift:

Definisi Umum: Configuration drift terjadi ketika keadaan aktual sistem tidak sesuai dengan keadaan yang diinginkan atau didefinisikan. Alat konfigurasi cenderung mengelola dan memperbaiki drift ini.
5. Continuous Configuration Automation:

Definisi Umum: Konsep ini mencakup penggunaan otomatisasi untuk menjaga keadaan konfigurasi di seluruh siklus hidup perangkat lunak dan infrastruktur.
6. Idempotent Configuration:

Definisi Umum: Konfigurasi idempoten berarti bahwa, jika konfigurasi dijalankan lebih dari satu kali, hasil akhirnya tetap sama. Ini penting untuk menghindari perubahan tak terduga.
7. Configuration Management Tools:

Rincian Diluar Hal Umum:
Ansible: Ansible menggunakan playbook untuk mendefinisikan dan mengotomatiskan konfigurasi.
Chef: Chef menggunakan "recipes" untuk mendefinisikan konfigurasi.
Puppet: Puppet menggunakan "manifests" untuk mendefinisikan konfigurasi.
Terraform: Terraform fokus pada provisioning dan manajemen infrastruktur melalui kode.
Command Line Interface (CLI):

1. Pipes and Redirection:

Definisi Umum: Pipes (|) digunakan untuk mengalirkan output dari satu perintah ke perintah lainnya, sedangkan redireksi (>, <, >>, 2>) digunakan untuk mengarahkan output atau input ke atau dari file.
2. Environment Variables:

Definisi Umum: Variabel lingkungan adalah nilai yang dapat diakses oleh program yang berjalan di dalam lingkungan sistem operasi. Di CLI, Anda dapat menggunakan perintah seperti export (Linux) atau set (Windows) untuk mengatur variabel lingkungan.
3. Command Substitution:

Definisi Umum: Command substitution memungkinkan Anda menyisipkan output dari suatu perintah ke dalam perintah lain. Dalam bash (Linux), ini dilakukan dengan menggunakan tanda kutip backticks (``) atau $().
4. Aliases and Functions:

Definisi Umum: Aliases memungkinkan Anda membuat singkatan untuk perintah yang sering digunakan, sementara fungsi memungkinkan Anda membuat blok kode yang dapat dipanggil sebagai satu perintah.
5. Job Control:

Definisi Umum: Job control memungkinkan Anda menjalankan perintah di latar belakang (&), menghentikan perintah yang sedang berjalan (Ctrl + Z), dan mengelola pekerjaan yang sedang berjalan (bg, fg).
6. Scripting and Automation:

Definisi Umum: CLI dapat digunakan untuk menulis skrip atau perintah yang menjalankan serangkaian tugas otomatis. Pemahaman tentang pengelolaan variabel, percabangan, dan perulangan diperlukan untuk scripting yang efektif.
7. Tab Completion:

Definisi Umum: Tab completion memungkinkan Anda menyelesaikan perintah atau nama file secara otomatis dengan menekan tombol tab. Ini meningkatkan efisiensi dan keakuratan saat mengetik perintah.
8. Remote Access and SSH:

Definisi Umum: CLI dapat digunakan untuk mengakses sistem jarak jauh melalui protokol SSH (Secure Shell). Ini memungkinkan administrasi jarak jauh dan pengelolaan server.
9. Regular Expressions in CLI:

Definisi Umum: Pemahaman tentang ekspresi reguler membantu dalam pencarian dan manipulasi teks di CLI. Perintah seperti grep, sed, dan awk menggunakan ekspresi reguler.
10. Debugging Tools:
- Rincian Diluar Hal Umum:
- strace (Linux): Menyediakan informasi rinci tentang panggilan sistem yang dilakukan oleh program.
- dtrace (Unix-like systems): Memberikan pemantauan dinamis dan tracing untuk sistem operasi dan aplikasi.

Pemahaman mendalam tentang Configuration Management dan Command Line Interface adalah keterampilan penting bagi seorang ahli informatika, terutama dalam konteks pengelolaan infrastruktur dan otomatisasi tugas-tugas sistem.