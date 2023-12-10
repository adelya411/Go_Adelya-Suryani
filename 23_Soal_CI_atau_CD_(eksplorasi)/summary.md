# CI/CD

Infrastructure and Deployment

# Continuous Integration
Proses yang diotomatisasi ini dilakukan untuk mengintegrasikan berbagai kode dari berbagai sumber potensial dengan tujuan membangun atau mengujinya.

## Cycle dalam CI/CD
1. *Continuous Integration (CI):* Development: Tim pengembangan bekerja pada fitur atau perbaikan bug dalam kode. Version Control: Kode dikelola dalam sistem kontrol versi seperti Git. Automated Build: Setiap perubahan kode yang di-commit ke repositori Git akan memicu build otomatis untuk menghasilkan binary atau paket perangkat lunak. Automated Testing: Setelah build selesai, berbagai tes otomatis dijalankan untuk memastikan bahwa perangkat lunak berfungsi seperti yang diharapkan. Reporting: Hasil dari build dan tes diunggah ke sistem pelaporan untuk diperiksa oleh pengembang.
2. *Continuous Delivery (CD):* Staging Environment: Setelah melewati CI, perangkat lunak dapat ditempatkan di lingkungan staging (mirip dengan produksi) untuk pengujian lebih lanjut. User Acceptance Testing (UAT): Tim QA atau pemangku kepentingan melaksanakan pengujian UAT untuk memastikan perangkat lunak siap untuk produksi. Manual Approval: Jika semua pengujian berhasil, tim pengembangan atau pemangku kepentingan memberikan persetujuan manual untuk melanjutkan ke produksi. Deployment to Production: Perangkat lunak diterapkan di lingkungan produksi.
3. *Continuous Deployment (CD):* Automated Deployment: Seiring dengan CI/CD, perangkat lunak secara otomatis diterapkan ke lingkungan produksi setelah lulus semua pengujian, tanpa persetujuan manual. Monitoring and Feedback: Setelah di produksi, perangkat lunak terus dimonitor. Jika terdeteksi masalah, dapat dilakukan rollback atau tindakan perbaikan segera.
4. *Continuous Monitoring:* Monitoring: Lingkungan produksi terus dimonitor untuk kinerja, keamanan, dan keandalan. Feedback Loop: Jika terjadi masalah, informasi kembali ke tim pengembangan untuk perbaikan lebih lanjut.
5. *Continuous Feedback and Improvement:* Feedback Gathering: Pemangku kepentingan dan pengguna memberikan umpan balik yang digunakan untuk perbaikan berkelanjutan. Back to Development: Berdasarkan umpan balik, tim pengembangan membuat perubahan dan memulai siklus CI/CD lagi untuk memperbarui perangkat lunak. Siklus CI/CD adalah pendekatan berkelanjutan yang memungkinkan pengembang untuk mengirimkan perubahan perangkat lunak secara cepat, aman, dan efisien. Ini menggabungkan otomatisasi, pengujian, pemantauan, dan umpan balik dari pengguna untuk memastikan perangkat lunak tetap berkualitas tinggi dan siap digunakan dalam waktu singkat.

What is CI CD deployment?
What Is CI/CD and How Does It Work? | Synopsys
CI and CD stand for continuous integration and continuous delivery/continuous deployment. In very simple terms, CI is a modern software development practice in which incremental code changes are made frequently and reliably. termasuk dalam hal yang tidak umum dalam beberapa bentuk program

CI/CD atau Continuous Integration dan Continuous Delivery/Continuous Deployment adalah praktik pengembangan perangkat lunak modern di mana perubahan kode dilakukan secara berkala dan dapat diandalkan. Dalam bahasa yang lebih sederhana, CI merupakan suatu pendekatan pengembangan perangkat lunak di mana perubahan kode kecil dilakukan secara rutin dan dapat diuji dengan cepat. Berikut adalah penjelasan lebih lanjut dalam bahasa Indonesia:

1. *Continuous Integration (CI):*
   - *Definisi:* Continuous Integration (Pengintegrasian Berkelanjutan) adalah praktik di mana pengembang secara teratur menggabungkan perubahan ke dalam repositori kode bersama dan secara otomatis menjalankan rangkaian tes otomatis untuk memverifikasi keberhasilan integrasi.
   - *Cara Kerja:*
      - Pengembang membuat perubahan ke dalam kode sumber.
      - Sistem CI secara otomatis menggabungkan perubahan tersebut ke dalam repositori bersama.
      - Rangkaian tes otomatis dijalankan untuk memastikan bahwa perubahan tersebut tidak mengenai fungsi yang sudah ada.

2. *Continuous Delivery (CD):*
   - *Definisi:* Continuous Delivery (Pengiriman Berkelanjutan) melibatkan otomatisasi proses pengiriman perangkat lunak ke lingkungan produksi setelah melalui tahap-tahap pengujian yang otomatis.
   - *Cara Kerja:*
      - Setelah perubahan melewati CI dan lulus semua tes, perangkat lunak tersebut dianggap "siap untuk pengiriman."
      - Tim pengembangan dapat memutuskan untuk melepaskan perangkat lunak tersebut ke lingkungan produksi dengan mengeksekusi proses pengiriman secara otomatis.

3. *Continuous Deployment (CD):*
   - *Definisi:* Continuous Deployment (Pengiriman Berkelanjutan) mirip dengan Continuous Delivery, namun dengan langkah tambahan di mana perubahan otomatis dideploy ke lingkungan produksi tanpa perlu campur tangan manusia.
   - *Cara Kerja:*
      - Setelah perubahan lulus semua tes dalam proses Continuous Delivery, perangkat lunak secara otomatis dideploy ke lingkungan produksi tanpa perlu intervensi manusia.

*Manfaat CI/CD:*
- *Ketanggapan Cepat:* Perubahan kode dapat diuji dan diterapkan dengan cepat, memungkinkan pengembang untuk mendapatkan umpan balik segera.
- *Pengurangan Risiko:* Tes otomatis dapat membantu mengidentifikasi masalah sejak dini, mengurangi risiko kesalahan saat perubahan diintegrasikan.
- *Keterlibatan Tim yang Efisien:* CI/CD memungkinkan tim untuk fokus pada inovasi daripada tugas manual pengujian dan penyebaran.

Dengan menerapkan CI/CD, tim pengembangan dapat meningkatkan efisiensi, keandalan, dan kecepatan pengembangan perangkat lunak mereka.
