# 🗳️ Aplikasi Sistem Pemungutan Suara Digital

## E-Voting Berbasis Terminal Menggunakan Go

Aplikasi ini merupakan program sederhana berbasis terminal yang digunakan untuk mengelola data kandidat dan menghitung suara dalam proses pemungutan suara digital. Program dibuat menggunakan bahasa pemrograman **Go** sebagai implementasi dari materi **Algoritma Pemrograman 2**.

---

## 📌 Deskripsi Program

Aplikasi **E-Voting** ini dirancang untuk membantu proses pengelolaan data kandidat dalam pemungutan suara sederhana. Program berjalan melalui terminal dan menyediakan fitur utama seperti menambah data kandidat, menampilkan data, mengubah data, menghapus data, mencari data, mengurutkan data, serta menampilkan statistik suara.

Program ini masih bersifat sederhana dan cocok digunakan sebagai media pembelajaran konsep dasar pemrograman, struktur data array, record atau struct, searching, sorting, dan modularisasi program.

---

## ✨ Fitur Program

| No | Fitur              | Keterangan                                             |
| -- | ------------------ | ------------------------------------------------------ |
| 1  | Tambah Kandidat    | Menambahkan data kandidat baru ke dalam sistem         |
| 2  | Tampilkan Kandidat | Menampilkan seluruh data kandidat yang tersimpan       |
| 3  | Ubah Kandidat      | Mengubah data kandidat berdasarkan nomor urut          |
| 4  | Hapus Kandidat     | Menghapus data kandidat berdasarkan nomor urut         |
| 5  | Sequential Search  | Mencari kandidat secara berurutan                      |
| 6  | Binary Search      | Mencari kandidat pada data yang sudah terurut          |
| 7  | Sorting            | Mengurutkan data kandidat berdasarkan pilihan tertentu |
| 8  | Statistik Suara    | Menampilkan total suara dan persentase suara kandidat  |

---

## 🧾 Struktur Data Kandidat

Setiap kandidat memiliki beberapa data utama, yaitu:

* Nomor urut
* Nama kandidat
* Visi misi
* Jumlah suara

Struktur data kandidat dalam program:

```go
type kandidat struct {
	nomorUrut int
	nama      string
	visiMisi  string
	suara     int
}
```

Array yang digunakan untuk menyimpan data kandidat:

```go
const NMAX = 999

type tabKandidat [NMAX]kandidat
```

---

## 🧠 Algoritma yang Digunakan

Program ini menerapkan beberapa algoritma dasar yang dipelajari dalam mata kuliah **Algoritma Pemrograman 2**.

| Algoritma         | Kegunaan                                                             |
| ----------------- | -------------------------------------------------------------------- |
| Sequential Search | Mencari kandidat satu per satu berdasarkan nomor urut                |
| Binary Search     | Mencari kandidat pada data yang sudah terurut berdasarkan nomor urut |
| Selection Sort    | Mengurutkan kandidat berdasarkan jumlah suara                        |
| Insertion Sort    | Mengurutkan kandidat berdasarkan nomor urut                          |

---

## 🔃 Pilihan Sorting

Program menyediakan beberapa pilihan pengurutan data kandidat:

1. Mengurutkan suara secara ascending
2. Mengurutkan suara secara descending
3. Mengurutkan nomor urut secara ascending
4. Mengurutkan nomor urut secara descending

---

## 🖥️ Tampilan Menu Program

```text
APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL

1. Tambah Kandidat
2. Tampilkan Kandidat
3. Ubah Kandidat
4. Hapus Kandidat
5. Sequential Search
6. Binary Search
7. Sorting
8. Statistik
0. Keluar

Pilih :
```

---

## 🚀 Cara Menjalankan Program

Pastikan bahasa pemrograman **Go** sudah terpasang pada perangkat.

Cek versi Go:

```bash
go version
```

Clone repository:

```bash
git clone https://github.com/stellaris-verse/Tugas-Besar-Alpro-2-Kelompok-13-IF4908.git
```

Masuk ke folder repository:

```bash
cd Tugas-Besar-Alpro-2-Kelompok-13-IF4908
```

Jalankan program:

```bash
go run tubes.go
```

---

## ✅ Validasi Program

Program memiliki beberapa validasi sederhana agar data yang dimasukkan lebih aman dan sesuai kebutuhan.

* Nomor urut tidak boleh bernilai negatif.
* Nomor urut tidak boleh sama dengan kandidat lain.
* Input jumlah suara tidak boleh bernilai negatif.
* Program menampilkan pesan jika data kandidat belum tersedia.
* Program menampilkan pesan jika kandidat yang dicari tidak ditemukan.
* Statistik suara tidak ditampilkan jika belum ada suara yang masuk.

---

## ⚠️ Keterbatasan Program

Program ini masih memiliki beberapa keterbatasan, yaitu:

* Belum memiliki fitur login pengguna.
* Belum memiliki data pemilih secara terpisah.
* Data belum tersimpan secara permanen ke file atau database.
* Program masih berjalan melalui terminal.
* Input teks masih sederhana karena menggunakan input dasar dari Go.

---

## 👥 Kontributor

| Nama                          | NIM          |
| ----------------------------- | ------------ |
| Muhammad Zehan Royyan Wahyudi | 103012500281 |
| Muhammad Ariq Alkosstra       | 103012500030 |

---

## 🎓 Mata Kuliah

**Algoritma Pemrograman 2**
Program Studi S1 Informatika
Fakultas Informatika
Telkom University

---

## 📄 Lisensi

Repository ini dibuat untuk keperluan tugas akademik.
