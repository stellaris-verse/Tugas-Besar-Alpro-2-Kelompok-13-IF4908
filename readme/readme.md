\# 🗳️ Aplikasi Sistem Pemungutan Suara Digital



\## E-Voting



Program sederhana berbasis terminal untuk mengelola data kandidat dan menghitung suara dalam proses pemungutan suara digital.



\---



\## 📌 Deskripsi



Aplikasi \*\*E-Voting\*\* ini dibuat menggunakan bahasa pemrograman \*\*Go\*\*. Program berjalan melalui terminal dan digunakan untuk mengelola data kandidat dalam proses pemungutan suara sederhana.



Program ini dapat menambahkan, menampilkan, mengubah, menghapus, mencari, mengurutkan, dan menampilkan statistik suara kandidat.



\---



\## ✨ Fitur Program



| No | Fitur              | Keterangan                                     |

| -- | ------------------ | ---------------------------------------------- |

| 1  | Tambah Kandidat    | Menambahkan data kandidat baru                 |

| 2  | Tampilkan Kandidat | Menampilkan seluruh data kandidat              |

| 3  | Ubah Kandidat      | Mengubah data kandidat berdasarkan nomor urut  |

| 4  | Hapus Kandidat     | Menghapus data kandidat berdasarkan nomor urut |

| 5  | Sequential Search  | Mencari kandidat secara berurutan              |

| 6  | Binary Search      | Mencari kandidat pada data yang sudah terurut  |

| 7  | Sorting            | Mengurutkan data kandidat                      |

| 8  | Statistik Suara    | Menampilkan persentase dan total suara         |



\---



\## 🧾 Data Kandidat



Setiap kandidat memiliki data sebagai berikut:



\* Nomor urut

\* Nama kandidat

\* Visi misi

\* Jumlah suara



Struktur data kandidat:



```go

type kandidat struct {

&#x09;nomorUrut int

&#x09;nama      string

&#x09;visiMisi  string

&#x09;suara     int

}

```



Array data kandidat:



```go

const NMAX = 999



type tabKandidat \[NMAX]kandidat

```



\---



\## 🧠 Algoritma yang Digunakan



Program ini menerapkan beberapa algoritma dasar dalam mata kuliah Algoritma Pemrograman 2.



| Algoritma         | Kegunaan                                              |

| ----------------- | ----------------------------------------------------- |

| Sequential Search | Mencari kandidat satu per satu berdasarkan nomor urut |

| Binary Search     | Mencari kandidat pada data yang sudah terurut         |

| Selection Sort    | Mengurutkan kandidat berdasarkan jumlah suara         |

| Insertion Sort    | Mengurutkan kandidat berdasarkan nomor urut           |



\---



\## 🔃 Pilihan Sorting



Program menyediakan empat pilihan pengurutan data:



1\. Suara ascending

2\. Suara descending

3\. Nomor urut ascending

4\. Nomor urut descending



\---



\## 🖥️ Tampilan Menu



```text

APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL

1\. Tambah Kandidat

2\. Tampilkan Kandidat

3\. Ubah Kandidat

4\. Hapus Kandidat

5\. Sequential Search

6\. Binary Search

7\. Sorting

8\. Statistik

0\. Keluar

Pilih :

```



\---



\## 🚀 Cara Menjalankan Program



Pastikan Go sudah terpasang di perangkat.



Cek versi Go:



```bash

go version

```



Jalankan program:



```bash

go run main.go

```



\---



\## ✅ Validasi Program



Program memiliki beberapa validasi sederhana:



\* Nomor urut tidak boleh `0`.

\* Nomor urut tidak boleh sama.

\* Input angka tidak boleh bernilai negatif.

\* Program menampilkan pesan jika data belum tersedia.

\* Program menampilkan pesan jika data yang dicari tidak ditemukan.

\* Statistik tidak ditampilkan jika belum ada suara.



\---



\## ⚠️ Keterbatasan Program



Program ini masih memiliki beberapa keterbatasan:



\* Belum memiliki fitur login.

\* Belum memiliki data pemilih secara terpisah.

\* Data belum tersimpan secara permanen ke file atau database.

\* Input nama dan visi misi masih terbatas satu kata karena menggunakan `fmt.Scan`.



\---



\## 👥 Kontributor



| Nama                          | NIM          |

| ----------------------------- | ------------ |

| Muhammad Zehan Royyan Wahyudi | 103012500281 |

| Muhammad Ariq Alkosstra       | 103012500030 |



\---



\## 🎓 Mata Kuliah



\*\*Algoritma Pemrograman 2\*\*

Program Studi S1 Informatika

Fakultas Informatika

Telkom University



\---



\## 📄 Lisensi



Repository ini dibuat untuk keperluan tugas akademik.



\---



