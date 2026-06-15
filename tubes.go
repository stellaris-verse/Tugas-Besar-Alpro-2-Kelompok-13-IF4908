package main

import "fmt"

const NMAX = 999

type kandidat struct {
	nomorUrut int
	nama      string
	visiMisi  string
	suara     int
}

type tabKandidat [NMAX]kandidat

func inputBilangan(teks string) int {
	var angka int

	angka = -1

	for angka < 0 {
		fmt.Print(teks)
		fmt.Scan(&angka)

		if angka < 0 {
			fmt.Println("Input harus suatu angka valid")
		}
	}

	return angka
}

func cariNomorSama(A tabKandidat, n int, nomor int) bool {
	var i int
	var ketemu bool

	i = 0
	ketemu = false

	for i < n && !ketemu {
		if A[i].nomorUrut == nomor {
			ketemu = true
		}
		i++
	}

	return ketemu
}

func tambahKandidat(A *tabKandidat, n *int) {
	var nomor int
	var valid bool

	if *n >= NMAX {
		fmt.Println("Data kandidat sudah penuh.")
		return
	}

	valid = false

	for !valid {
		nomor = inputBilangan("Nomor Urut : ")

		if nomor == 0 {
			fmt.Println("Nomor urut harus lebih dari 0")
		} else if cariNomorSama(*A, *n, nomor) {
			fmt.Println("Nomor urut sudah digunakan")
		} else {
			valid = true
		}
	}

	A[*n].nomorUrut = nomor

	fmt.Print("Nama       : ")
	fmt.Scan(&A[*n].nama)

	fmt.Print("Visi Misi  : ")
	fmt.Scan(&A[*n].visiMisi)

	A[*n].suara = inputBilangan("Suara      : ")

	*n = *n + 1

	fmt.Println("Data kandidat berhasil ditambahkan")
}

func tampilkanData(A tabKandidat, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data.")
	} else {
		i = 0
		for i < n {
			fmt.Println("--------------------")
			fmt.Println("Nomor :", A[i].nomorUrut)
			fmt.Println("Nama  :", A[i].nama)
			fmt.Println("Visi  :", A[i].visiMisi)
			fmt.Println("Suara :", A[i].suara)
			i++
		}
	}
}

func sequentialSearch(A tabKandidat, n int, nomor int) int {
	var i int
	var ketemu int

	ketemu = -1
	i = 0

	for ketemu == -1 && i < n {
		if A[i].nomorUrut == nomor {
			ketemu = i
		}
		i++
	}

	return ketemu
}

func ubahData(A *tabKandidat, n int) {
	var nomor int
	var idx int

	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	nomor = inputBilangan("Nomor urut yang dicari: ")

	idx = sequentialSearch(*A, n, nomor)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Nama baru      : ")
		fmt.Scan(&A[idx].nama)

		fmt.Print("Visi misi baru : ")
		fmt.Scan(&A[idx].visiMisi)

		A[idx].suara = inputBilangan("Suara baru     : ")

		fmt.Println("Data berhasil diubah")
	}
}

func hapusData(A *tabKandidat, n *int) {
	var nomor int
	var idx int
	var i int

	if *n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	nomor = inputBilangan("Nomor urut yang dicari: ")

	idx = sequentialSearch(*A, *n, nomor)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		i = idx

		for i <= *n-2 {
			A[i] = A[i+1]
			i++
		}

		*n = *n - 1

		fmt.Println("Data berhasil dihapus")
	}
}

func insertionSortNomorAsc(A *tabKandidat, n int) {
	var pass int
	var i int
	var temp kandidat

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.nomorUrut < A[i-1].nomorUrut {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
		pass++
	}
}

func insertionSortNomorDesc(A *tabKandidat, n int) {
	var pass int
	var i int
	var temp kandidat

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.nomorUrut > A[i-1].nomorUrut {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
		pass++
	}
}

func binarySearch(A tabKandidat, n int, nomor int) int {
	var kiri int
	var kanan int
	var tengah int
	var idx int

	kiri = 0
	kanan = n - 1
	idx = -1

	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2

		if nomor < A[tengah].nomorUrut {
			kanan = tengah - 1
		} else if nomor > A[tengah].nomorUrut {
			kiri = tengah + 1
		} else {
			idx = tengah
		}
	}

	return idx
}

func selectionSortSuaraAsc(A *tabKandidat, n int) {
	var pass int
	var i int
	var idx int
	var temp kandidat

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass

		for i < n {
			if A[idx].suara > A[i].suara {
				idx = i
			}
			i++
		}

		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func selectionSortSuaraDesc(A *tabKandidat, n int) {
	var pass int
	var i int
	var idx int
	var temp kandidat

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass

		for i < n {
			if A[idx].suara < A[i].suara {
				idx = i
			}
			i++
		}

		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func statistikSuara(A tabKandidat, n int) {
	var i int
	var total int
	var persen float64

	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	total = 0
	i = 0

	for i < n {
		total = total + A[i].suara
		i++
	}

	if total <= 0 {
		fmt.Println("Belum ada suara")
		return
	}

	i = 0

	for i < n {
		persen = float64(A[i].suara) / float64(total) * 100
		fmt.Printf("%s : %.2f%%\n", A[i].nama, persen)
		i++
	}

	fmt.Println("Total suara:", total)
}

func main() {
	var data tabKandidat
	var n int
	var pilih int
	var nomor int
	var idx int
	var metode int

	n = 0

	for {
		fmt.Println()
		fmt.Println("APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Ubah Kandidat")
		fmt.Println("4. Hapus Kandidat")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Sorting")
		fmt.Println("8. Statistik")
		fmt.Println("0. Keluar")

		pilih = inputBilangan("Pilih : ")

		switch pilih {

		case 1:
			tambahKandidat(&data, &n)

		case 2:
			tampilkanData(data, n)

		case 3:
			ubahData(&data, n)

		case 4:
			hapusData(&data, &n)

		case 5:
			if n == 0 {
				fmt.Println("Belum ada data.")
			} else {
				nomor = inputBilangan("Nomor urut : ")

				idx = sequentialSearch(data, n, nomor)

				if idx == -1 {
					fmt.Println("Tidak ditemukan")
				} else {
					fmt.Println("Data ditemukan")
					fmt.Println("Nomor :", data[idx].nomorUrut)
					fmt.Println("Nama  :", data[idx].nama)
					fmt.Println("Visi  :", data[idx].visiMisi)
					fmt.Println("Suara :", data[idx].suara)
				}
			}

		case 6:
			if n == 0 {
				fmt.Println("Belum ada data.")
			} else {
				insertionSortNomorAsc(&data, n)

				nomor = inputBilangan("Nomor urut : ")

				idx = binarySearch(data, n, nomor)

				if idx == -1 {
					fmt.Println("Tidak ditemukan")
				} else {
					fmt.Println("Data ditemukan")
					fmt.Println("Nomor :", data[idx].nomorUrut)
					fmt.Println("Nama  :", data[idx].nama)
					fmt.Println("Visi  :", data[idx].visiMisi)
					fmt.Println("Suara :", data[idx].suara)
				}
			}

		case 7:
			if n == 0 {
				fmt.Println("Belum ada data.")
			} else {
				fmt.Println("1. Suara ascending")
				fmt.Println("2. Suara descending")
				fmt.Println("3. Nomor urut ascending")
				fmt.Println("4. Nomor urut descending")

				metode = inputBilangan("Pilih : ")

				switch metode {

				case 1:
					selectionSortSuaraAsc(&data, n)
					fmt.Println("Data diurutkan berdasarkan suara dari kecil ke besar")
					tampilkanData(data, n)

				case 2:
					selectionSortSuaraDesc(&data, n)
					fmt.Println("Data diurutkan berdasarkan suara dari besar ke kecil")
					tampilkanData(data, n)

				case 3:
					insertionSortNomorAsc(&data, n)
					fmt.Println("Data diurutkan berdasarkan nomor urut dari kecil ke besar")
					tampilkanData(data, n)

				case 4:
					insertionSortNomorDesc(&data, n)
					fmt.Println("Data diurutkan berdasarkan nomor urut dari besar ke kecil")
					tampilkanData(data, n)

				default:
					fmt.Println("Pilihan sorting tidak tersedia")
				}
			}

		case 8:
			statistikSuara(data, n)

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
