# Version Control and Branch Management

VERSION CONTROL AND BRANCH MANAGEMENT (GIT)

Git : Salah satu vcs populer yang digunakan para developer untuk mengembangkan software secara bersama-sama.

The stagging area : working directory => (git add) => stagging area => (git commit) => repository

=> git add untuk menambahkan perubahan ke staging area.
=> git status untuk memeriksa apakah perubahan telah ditambahkan ke staging area.
=> git commit untuk membuat snapshot permanen dari perubahan dengan pesan commit yang relevan.
=> git diff dan git stash adalah perintah-perintah lain yang umum digunakan dalam Git untuk melihat perbedaan antara revisi-revisi atau menyimpan perubahan sementara.
=> git checkout adalah perintah dalam Git yang memiliki beberapa fungsi tergantung pada bagaimana menggunakannya. Perintah ini digunakan untuk melakukan beberapa tugas yang berkaitan dengan navigasi, cabang, commit, dan pengelolaan perubahan dalam repositori.
=> git push untuk mengirim perubahan lokal yang telah di-commit ke repositori jarak jauh.
=> git pull untuk mengambil perubahan dari repositori jarak jauh dan menggabungkannya dengan cabang lokal yang sedang aktif. 
=> git branch digunakan untuk melihat, membuat, menghapus, atau mengelola cabang-cabang dalam repositori.
=> git merge digunakan untuk menggabungkan perubahan dari satu cabang ke dalam cabang lain.
=> Pull Request (PR) adalah cara bagi pengembang untuk mengajukan perubahan kode yang ada di cabang mereka ke cabang utama (biasanya disebut "master" atau "main") atau cabang lainnya dalam repositori, dengan tujuan untuk menggabungkan perubahan tersebut setelah diperiksa dan diulas.

Version Control:

Version control, juga dikenal sebagai kontrol sumber atau kontrol revisi, adalah sistem yang mencatat perubahan pada satu atau sekelompok berkas dari waktu ke waktu sehingga Anda dapat mengingat versi tertentu nantinya. Ini memungkinkan beberapa kontributor untuk bekerja pada suatu proyek secara bersamaan dan membantu melacak serta mengelola perubahan. Ada dua jenis utama kontrol versi: terpusat dan terdistribusi.

Sistem Kontrol Versi Terpusat (CVCS):

Pada CVCS, ada satu repositori pusat yang menyimpan seluruh riwayat proyek.
Pengguna mengambil berkas dari repositori pusat ini, melakukan perubahan, dan kemudian mengirimkan perubahan tersebut kembali ke repositori.
Contoh: CVS (Concurrent Versions System), SVN (Apache Subversion).
Sistem Kontrol Versi Terdistribusi (DVCS):

Pada DVCS, setiap pengguna memiliki salinan lengkap dari repositori, termasuk seluruh riwayatnya.
Pengguna dapat bekerja secara independen dan mengirimkan perubahan ke repositori lokal mereka. Mereka juga dapat berbagi perubahan dengan orang lain dengan cara menarik dan mendorong perubahan antar repositori.
Contoh: Git, Mercurial.


Branching:

Branching adalah konsep penting dalam sistem kontrol versi, memungkinkan pengembang untuk bekerja pada versi yang berbeda dari kode secara bersamaan. Ini memungkinkan isolasi fitur, perbaikan bug, atau eksperimen tanpa mempengaruhi kode utama sampai perubahan siap untuk digabungkan.

Master/Main/Branch:

Cabang utama (sering disebut "master" atau "main") mewakili versi kode yang stabil dan siap produksi.
Semua cabang fitur biasanya dibuat dari dan digabungkan kembali ke cabang utama.
Branch Fitur:

Pengembang membuat cabang fitur untuk bekerja pada fitur atau perubahan baru.
Cabang-cabang ini biasanya memiliki jangka waktu pendek dan digabungkan kembali ke cabang utama setelah fitur selesai.
Cabang Rilis:

Cabang rilis dibuat saat mempersiapkan rilis perangkat lunak baru.
Ini memungkinkan stabilisasi dan pengujian kode sebelum digabungkan ke cabang utama untuk rilis.
Cabang Perbaikan Darurat:

Cabang perbaikan darurat dibuat untuk mengatasi isu kritis dalam kode produksi.
Ini memungkinkan perbaikan cepat tanpa mengganggu pengembangan yang sedang berlangsung di cabang utama.
Git - Sistem Kontrol Versi Terdistribusi:

Git adalah salah satu sistem kontrol versi terdistribusi yang paling populer. Ini memiliki model branching yang kuat dan memberikan fleksibilitas dalam mengelola perubahan kode.

Alur Kerja Dasar Git:

Clone: Salin repositori ke mesin lokal Anda.
Add: Menyiapkan perubahan untuk di-commit.
Commit: Menyimpan perubahan ke repositori lokal.
Push: Mengunggah perubahan ke repositori jarak jauh.
Branching di Git:

git branch: Membuat cabang baru.
git checkout atau git switch: Beralih antara cabang.
git merge: Menggabungkan perubahan dari satu cabang ke cabang lain.
git pull: Mengambil perubahan dari repositori jarak jauh dan menggabungkannya ke dalam cabang saat ini.
Memahami kontrol versi dan branching sangat penting untuk pengembangan perangkat lunak kolaboratif, memungkinkan tim untuk bekerja secara efisien, melacak perubahan, dan mengelola riwayat proyek dengan efektif.