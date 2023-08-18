Deskripsi
Ini adalah proyek Assement Moonlay Technologies,dengan Cloudinary untuk penyimpanan file, dan integrasi dengan PostgreSQL menggunakan GORM .


Instalasi
Clone repository ini ke direktori lokal Anda.
Buka terminal dan ketik perintah go mod tidy untuk mengelola dependensi Go.
Pastikan Anda memiliki akun Cloudinary dan sesuaikan konfigurasi pada file .env(di sini saya memakai akun saya).
Install PostgreSQL dan PGAdmin jika belum terpasang.
Jalankan perintah go run main.go di terminal untuk memulai server.
Akses API melalui URL http://localhost:5000/api/v1/sublist menggunakan metode HTTP POST. Tambahkan key-value di dalam body form-data: list_id (dengan ID list yang tersedia).
