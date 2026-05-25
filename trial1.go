package main

import "fmt"

type mahasiswa struct {
	NIM               string
	Nama              string
	tunggakan         int
	statuspembayaraan bool
	riwayat           [100]transaksi
	jumlahbayar       int
}

type Database struct {
	mahasiswa      [100]mahasiswa
	mahasiswaCount int
}

type transaksi struct {
	tanggal string
	nominal int
}

var DB Database

// Cari Bedasarkan NIM
func squential(A *Database, nim string, indeks *int) {
	var i int
	*indeks = -1

	for i = 0; i < A.mahasiswaCount; i++ {
		if A.mahasiswa[i].NIM == nim {
			*indeks = i
		}
	}
}

func main() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("           SISTEM KAS MAHASISWA")
		fmt.Println("========================================")
		fmt.Println("1. Kelola Data Mahasiswa")
		fmt.Println("2. Bayar")
		fmt.Println("3. Cari")
		fmt.Println("4. Urut")
		fmt.Println("0. keluar")
		fmt.Print("Pilih: ")

		var pilih int
		fmt.Scan(&pilih)
		fmt.Scanln()

		switch pilih {
		case 1:
			menusiswa(&DB)
		case 2:
			bayar(&DB)
		case 3:
			cari(&DB)
		case 4:
			urut(&DB)
		case 0:
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

// Fitur CRUD
func menusiswa(A *Database) {
	for {
		fmt.Println("\n========================================")
		fmt.Println("        KELOLA DATA MAHASISWA")
		fmt.Println("========================================")
		fmt.Println("Pilih Menu: ")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. tampil")
		fmt.Println("4. Hapus")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih (1-5): ")

		var menu int
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			tambah(A)
		case 2:
			ubah(A)
		case 3:
			tampil(A)
		case 4:
			hapus(A)
		case 0:
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
		}

	}
}

func tambah(A *Database) {
	fmt.Println("\n========================================")
	fmt.Println("         TAMBAH DATA MAHASISWA")
	fmt.Println("========================================")
	if A.mahasiswaCount < 100 {
		var mhs mahasiswa

		fmt.Print("Nama	: ")
		fmt.Scanln(&mhs.Nama)

		fmt.Print("NIM	: ")
		fmt.Scanln(&mhs.NIM)

		if mhs.NIM == "" || mhs.Nama == "" {
			fmt.Println("GAGAL: NIM dan Nama tidak boleh kosong!!")
			return
		}

		var cek int
		squential(A, mhs.NIM, &cek)
		if cek != -1 {
			fmt.Println("NIM sudah digunakan!!")
			return
		}

		for i := 0; i < A.mahasiswaCount; i++ {
			if A.mahasiswa[i].Nama == mhs.Nama {
				fmt.Println("Nama sudah digunakan!!")
				return
			}
		}

		mhs.tunggakan = 0
		mhs.statuspembayaraan = false
		mhs.jumlahbayar = 0
		A.mahasiswa[A.mahasiswaCount] = mhs
		A.mahasiswaCount++
		fmt.Println("== Mahasiswa Berhasil Ditambahkan ==")
	} else {
		fmt.Println("== Database Mahasiswa Penuh ==")
	}
}

func ubah(A *Database) {
	fmt.Println("\n========================================")
	fmt.Println("         UPDATE DATA MAHASISWA")
	fmt.Println("========================================")

	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan diupdate: ")
	fmt.Scanln(&nim)

	if nim == "" {
		fmt.Println("NIM tidak boleh kosong!")
		return
	}

	squential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("Data mahasiswa tidak ditemukan!")
		return
	}

	var Nimbaru, Namabaru string
	fmt.Print("Masukkan NIM baru	: ")
	fmt.Scanln(&Nimbaru)
	fmt.Print("Masukkan Nama baru	: ")
	fmt.Scanln(&Namabaru)

	if Nimbaru != "" {
		var cek int
		squential(A, Nimbaru, &cek)
		if cek != -1 && cek != indeks {
			fmt.Println("NIM baru sudah digunakan oleh mahasiswa lain!")
			return
		}
		A.mahasiswa[indeks].NIM = Nimbaru
	}

	if Namabaru != "" {
		A.mahasiswa[indeks].Nama = Namabaru
	}

	if Nimbaru == "" && Namabaru == "" {
		fmt.Println("Tidak ada perubahan data.")
	} else {
		fmt.Println("Data berhasil diubah!")
	}

}

func tampil(A *Database) {
	fmt.Println("\n========================================")
	fmt.Println("         DAFTAR MAHASISWA")
	fmt.Println("========================================")
	if A.mahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
	} else {
		fmt.Printf(" %-3s | %-10s | %-25s | %-12s | %-10s\n", "No", "NIM", "Nama", "Total Bayar", "Status")
		fmt.Println("-----+------------+---------------------------+--------------+------------")
		for i := 0; i < A.mahasiswaCount; i++ {
			status := "Belum Lunas"
			if A.mahasiswa[i].statuspembayaraan {
				status = "Lunas"
			}
			fmt.Printf(" %-3d | %-10s | %-25s | Rp %-9d | %-10s\n",
				i+1, A.mahasiswa[i].NIM, A.mahasiswa[i].Nama, A.mahasiswa[i].jumlahbayar, status)
		}
	}
	fmt.Println()
}

func hapus(A *Database) {
	fmt.Println("\n========================================")
	fmt.Println("         HAPUS DATA MAHASISWA")
	fmt.Println("========================================")
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM mahasiswa : ")
	fmt.Scanln(&nim)

	if nim == "" {
		fmt.Println("NIM tidak boleh kosong!")
		return
	}

	squential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("Data mahasiswa tidak ditemukan!")
		return
	}

	for i := indeks; i < A.mahasiswaCount; i++ {
		A.mahasiswa[i] = A.mahasiswa[i+1]
	}
	A.mahasiswaCount--
	fmt.Println("Data mahasiswa berhasil dihapus!")
}

// Fitur Bayar
func bayar(A *Database) {
	fmt.Println("\n========================================")
	fmt.Println("           PEMBAYARAN KAS")
	fmt.Println("========================================")

	var nim string
	var indeks int
	var bayar transaksi

	fmt.Print("Masukkan NIM Mahasiswa yang akan melakukan pembayaran: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)

	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		fmt.Print("Masukan Tanggal Pembayaran (DD/MM/YYYY): ")
		fmt.Scan(&bayar.tanggal)

		fmt.Print("Masukan Nominal Pembayaran: ")
		fmt.Scan(&bayar.nominal)

		A.mahasiswa[indeks].riwayat[A.mahasiswa[indeks].jumlahbayar] = bayar
		A.mahasiswa[indeks].jumlahbayar++
		A.mahasiswa[indeks].statuspembayaraan = true
		fmt.Println("== Pembayaran Berhasil Dicatat ==")
	}
}

// Fitur Cari
func cari(A *Database) {
	for {
		fmt.Println("====================")
		fmt.Println("   Cari Mahasiswa")
		fmt.Println("====================")
		fmt.Println("Pilih Menu: ")
		fmt.Println("1. Sequntial")
		fmt.Println("2. Binary")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var menu int
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			cari_squential(A)
		case 2:
			cari_binary(A)
		case 0:
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func cari_squential(A *Database) {
	var nim string
	var indeks int

	fmt.Println("====")
	fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		if A.mahasiswa[indeks].statuspembayaraan == false {
			fmt.Println("== Daftar Mahasiswa Belum Lunas ==")
			fmt.Println("NIM  :", A.mahasiswa[indeks].NIM)
			fmt.Println("Nama :", A.mahasiswa[indeks].Nama)
		} else {
			fmt.Println("Mahasiswa Sudah Membayar secara Lunas")
			fmt.Println("NIM  :", A.mahasiswa[indeks].NIM)
			fmt.Println("Nama :", A.mahasiswa[indeks].Nama)
		}
	}
}

func cari_binary(A *Database) {
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
	fmt.Scan(&nim)
	insertionsort(A)
	binary(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		fmt.Println("NIM  :", A.mahasiswa[indeks].NIM)
		fmt.Println("Nama :", A.mahasiswa[indeks].Nama)
	}
}

func binary(A *Database, nim string, indeks *int) {
	var kanan, kiri, tengah int

	*indeks = -1
	kiri = 0
	kanan = A.mahasiswaCount - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if A.mahasiswa[tengah].NIM == nim {
			*indeks = tengah
			kiri = kanan + 1
		} else if A.mahasiswa[tengah].NIM < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
}

// Fitur Urut
func urut(A *Database) {
	for {
		fmt.Println("=========================")
		fmt.Println("   Urut Data Mahasiswa")
		fmt.Println("=========================")
		fmt.Println("Pilih Menu: ")
		fmt.Println("1. Selection")
		fmt.Println("2. Insertion")
		fmt.Println("0. kembali")
		fmt.Print("Pilih: ")

		var menu int
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			selectionsort(A)
		case 2:
			insertionsort(A)
		case 0:
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func selectionsort(A *Database) {
	for i := 0; i < A.mahasiswaCount-1; i++ {
		min := i
		for j := i + 1; j < A.mahasiswaCount; j++ {
			if A.mahasiswa[min].Nama > A.mahasiswa[j].Nama || A.mahasiswa[min].tunggakan > A.mahasiswa[j].tunggakan {
				min = j
			}
		}
		temp := A.mahasiswa[i]
		A.mahasiswa[i] = A.mahasiswa[min]
		A.mahasiswa[min] = temp
	}
}
func insertionsort(A *Database) {
	var i, j int
	var sementara mahasiswa

	for i = 1; i < A.mahasiswaCount; i++ {
		sementara = A.mahasiswa[i]
		j = i - 1

		for j >= 0 && A.mahasiswa[j].Nama > sementara.Nama || A.mahasiswa[j].tunggakan > sementara.tunggakan {
			A.mahasiswa[j+1] = A.mahasiswa[j]
			j = j - 1
		}
		A.mahasiswa[j+1] = sementara
	}
}
