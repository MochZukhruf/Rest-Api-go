# Rest-API-Go

Rest-API-Go adalah sebuah implementasi sederhana REST API menggunakan bahasa pemrograman Go (Golang) dengan arsitektur yang terorganisir (repository-service-controller). API ini dirancang untuk melakukan operasi CRUD pada data produk menggunakan database MySQL.

---

## **Fitur**

- **CRUD Produk**:
  - **Buat Produk**: Tambahkan produk baru ke dalam database.
  - **Baca Produk**: Tampilkan semua produk atau produk tertentu berdasarkan ID.
  - **Perbarui Produk**: Ubah informasi produk.
  - **Hapus Produk**: Hapus produk dari database.
- **Error Handling**: Menampilkan pesan error yang sesuai untuk setiap permintaan yang gagal.

---

## **Teknologi yang Digunakan**

- **Bahasa Pemrograman**: Go (Golang)
- **Framework**: Gin
- **Database**: MySQL
- **ORM**: GORM
- **Dependency Management**: Go Modules

---

## **Instalasi**

### **1. Prasyarat**

Pastikan Anda sudah menginstal:

- [Go](https://go.dev/dl/) versi 1.21 atau yang lebih baru.
- MySQL Database.
- XAMPP (opsional, untuk pengelolaan MySQL lokal).
- Git (untuk mengelola repositori).

### **2. Clone Repositori**

Clone proyek ke dalam direktori lokal:

```bash
git clone https://github.com/username/rest-api-go.git
cd rest-api-go
```
