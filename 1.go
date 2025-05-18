package main

import (
	"fmt"
	//"math/rand"
	//"time"
	"os"
	"os/exec"
)

const NMAX int = 1024

type arrAkun [NMAX]Akun
type kripto [NMAX]Akripto
type porto [NMAX]Aportofolio
type riwayat [NMAX]Ariwayat

type Akun struct {
	username    string
	password    string
	recovery    string
	saldo       int
	aset        float64
	nkriptoAset int
	namaKrip    string
	jHarga      int
	kripto
}

type Akripto struct {
	nama        string
	harga       int
	jumlah      float64
	totalKripto int
}

type Aportofolio struct {
	keuntungan int
	kerugian   int
}

type Ariwayat struct {
	status  string
	nkripto string
	jumlah  float64
	harga   int
	tanggal string
}

func clearScreen() {
	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func enterKembali() {
	var x string
	fmt.Println("Tekan enter untuk kembali...")
	fmt.Scanln()
	fmt.Scanln(&x)
}

func main() {
	var a arrAkun
	var k kripto
	var b, nAkun, idxAcc int
	var nkripto int = 10
	var exit bool = false

	k[0].nama = "Bitcoin"
	k[1].nama = "Ethereum"
	k[2].nama = "Litecoin"
	k[3].nama = "Solana"
	k[4].nama = "Fartcoin"
	k[5].nama = "Dogecoin"
	k[6].nama = "CoinNavigator"
	k[7].nama = "Cardano"
	k[8].nama = "Ondo"
	k[9].nama = "BNB"
	k[0].harga = 1700500000
	k[1].harga = 43000000
	k[2].harga = 1600000
	k[3].harga = 2900000
	k[4].harga = 21800
	k[5].harga = 3800
	k[6].harga = 195000
	k[7].harga = 13100
	k[8].harga = 16700
	k[9].harga = 10700000

	akun(&a, &nAkun, &idxAcc, &exit)
	if exit {
		return
	}
	for {
		clearScreen()
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Menu Utama               |")
		fmt.Println("-----------------------------------------")
		fmt.Println("username:", a[idxAcc].username)
		fmt.Println("Total Aset: Rp", int(a[idxAcc].aset))
		fmt.Println(idxAcc)
		fmt.Println("-----------------------------------------")
		for i := 0; i < 3; i++ {
			fmt.Println(i+1, k[i].nama, "Harga:", k[i].harga)
		}
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Lihat Daftar Kripto")
		fmt.Println("2. Cari Kripto")
		fmt.Println("3. Edit Kripto")
		fmt.Println("4. Lihat Portofolio")
		fmt.Println("5. Lihat Riwayat Transaksi")
		fmt.Println("6. Beli Kripto")
		fmt.Println("7. LogOut")
		fmt.Println("8. Keluar")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&b)

		switch b {
		case 1:
			lihatKripto(k, nkripto)
			enterKembali()
		case 2:
			cariKriptoSeq(k, nkripto)
		case 3:
			menuKripto(&k, &nkripto)
		case 4:
			lihatPortofolio(a, idxAcc)
		case 5:
		case 6:
			Beli(&a, k, idxAcc, nkripto)
		case 7:
			akun(&a, &nAkun, &idxAcc, &exit)
			if exit {
				return
			}
		case 8:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

}

func akun(A *arrAkun, nAkun, idx *int, exit *bool) {
	var a int
	var cek bool
	*exit = false
	for {
		clearScreen()
		fmt.Println("-----------------------------------------")
		fmt.Println("|   Selamat datang di aplikasi kripto   |")
		fmt.Println("-----------------------------------------")
		fmt.Println("1. login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&a)
		switch a {
		case 1:
			login(*A, *nAkun, idx, &cek)
			if cek {
				return
			}
		case 2:
			register(A, nAkun)
		case 3:
			*exit = true
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func login(A arrAkun, n int, idx *int, cek *bool) {
	var user, pass, in string
	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)
		*idx = -1
		for j := 0; j < n; j++ {
			if A[j].username == user && A[j].password == pass {
				*idx = j
				*cek = true
				fmt.Println("Login berhasil!")
				return
			}
		}
		fmt.Println("Username atau Password salah!")
	}

	fmt.Print("Lupa Password? (y/n)? ")
	fmt.Scan(&in)
	if in == "y" {
		var recUser, recFood string
		fmt.Print("Masukkan username: ")
		fmt.Scan(&recUser)
		fmt.Print("Apa makanan kesukaanmu? ")
		fmt.Scan(&recFood)
		for j := 0; j < n; j++ {
			if A[j].username == recUser && A[j].recovery == recFood {
				fmt.Println("Username:", A[j].username)
				fmt.Println("Password:", A[j].password)
				enterKembali()
			}
		}
	}
	*cek = false
}

func register(A *arrAkun, n *int) {
	clearScreen()
	fmt.Print("Masukkan username: ")
	fmt.Scan(&A[*n].username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&A[*n].password)
	fmt.Println("-----------------------------------------")
	fmt.Println("Recovery")
	fmt.Println("Masukkan nama makanan favorit anda: ")
	fmt.Scan(&A[*n].recovery)
	fmt.Println("Akun berhasil dibuat!")
	enterKembali()
	*n++
}

func menuKripto(k *kripto, n *int) {
	var a int
	var x string

	for {
		clearScreen()
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Menu Edit Kripto         |")
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Lihat Daftar Kripto")
		fmt.Println("2. Tambah Kripto")
		fmt.Println("3. Hapus Kripto")
		//fmt.Println("4. Urutkan Kripto")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&a)
		switch a {
		case 1:
			lihatKripto(*k, *n)
			fmt.Print("Urutkan Kripto (y/n)? ")
			fmt.Scan(&x)
			if x == "y" {
				urutkanKripto(k, *n)
				lihatKripto(*k, *n)
			} else {
				return
			}
		case 2:
			inputKripto(k, n)
		case 3:
			fmt.Print("Masukkan nama kripto yang ingin dihapus: ")
			fmt.Scan(&x)
			hapusKripto(k, n, x)
		case 4:

		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func lihatKripto(k kripto, n int) {
	//var dummy string
	clearScreen()
	fmt.Println("-----------------------------------------")
	fmt.Println("|              Daftar Kripto            |")
	fmt.Println("-----------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, k[i].nama, "Harga:", k[i].harga)
	}
	fmt.Println("-----------------------------------------")
}

func inputKripto(k *kripto, n *int) {
	var x string

	for {
		clearScreen()
		fmt.Print("Masukkan nama kripto: ")
		fmt.Scan(&k[*n].nama)
		fmt.Print("Masukkan harga kripto: ")
		fmt.Scan(&k[*n].harga)
		if k[*n].harga > 0 {
			*n++
		} else {
			fmt.Println("Harga tidak boleh 0 atau negatif")
			fmt.Print("Masukkan harga kripto: ")
			fmt.Scan(&k[*n].harga)
			*n++
		}
		fmt.Print("Apakah ingin menambahkan lagi(y/n)? ")
		fmt.Scan(&x)
		if x == "n" {
			break
		}
	}
}

func hapusKripto(k *kripto, n *int, x string) {
	var i, j int
	for i = 0; i < *n; i++ {
		if k[i].nama == x {
			for j = i; j < *n-1; j++ {
				k[j] = k[j+1]
			}
			*n--
		}
	}
	fmt.Println("Kripto sudah terhapus.")
}

func cariKriptoBin(k *kripto, n int) {
	var search string
	var idx int

	fmt.Print("Masukkan nama kripto yang ingin dicari: ")
	fmt.Scan(&search)

	SelectionSortStrAsc(k, n)
	idx = binarySearchStrAsc(*k, n, search)
	if idx != -1 {
		fmt.Println("Kripto ditemukan!")
		fmt.Println("Nama:", k[idx].nama, "Harga:", k[idx].harga)
	} else {
		fmt.Println("Kripto tidak ditemukan!")
	}
	enterKembali()
}

func urutkanKripto(k *kripto, n int) {
	var y, s string
	fmt.Print("Urutkan berdasarkan harga (tertinggi/terendah)?: ")
	fmt.Scan(&y)
	if y == "tertinggi" {
		SelectionSortIntDes(k, n)
		fmt.Println("Data kripto sudah diurutkan")
	} else if y == "terendah" {
		insertionSortIntAsc(k, n)
		fmt.Println("Data kripto sudah diurutkan")
	}
	lihatKripto(*k, n)
	fmt.Print("Cari Kripto (y/n)? ")
	fmt.Scan(&s)
	if s == "y" {
		cariKriptoBin(k, n)
	} else if s == "n" {
		return
	}
}

func cariKriptoSeq(k kripto, n int) {
	var y, search2 string
	var search1, idx int

	fmt.Print("Cari berdasarkan (harga/nama): ")
	fmt.Scan(&y)
	if y == "harga" {
		fmt.Print("Masukkan harga kripto yang ingin dicari: ")
		fmt.Scan(&search1)
		idx = sequentialSearchInt(k, n, search1)
		if idx != -1 {
			fmt.Println("Kripto ditemukan di index:", idx)
			fmt.Println("Nama:", k[idx].nama, "Harga:", k[idx].harga)
		} else {
			fmt.Println("Kripto tidak ditemukan")
		}
	} else if y == "nama" {
		fmt.Print("Masukkan nama kripto yang ingin dicari: ")
		fmt.Scan(&search2)
		idx = sequentialSearchStr(k, n, search2)
		if idx != -1 {
			fmt.Println("Kripto ditemukan di index:", idx)
			fmt.Println("Nama:", k[idx].nama, "Harga:", k[idx].harga)
		} else {
			fmt.Println("Kripto tidak ditemukan")
		}
	}
	enterKembali()
}

func binarySearchStrAsc(A kripto, n int, x string) int {
	var left, right, mid, idx int

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x > A[mid].nama {
			right = mid - 1
		} else if x < A[mid].nama {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func sequentialSearchInt(A kripto, n, x int) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].harga == x {
			idx = i
		}
	}
	return idx
}

func sequentialSearchStr(A kripto, n int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].nama == x {
			idx = i
		}
	}
	return idx
}

func SelectionSortIntDes(A *kripto, n int) {
	var i, idx, pass int
	var temp Akripto

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].harga > A[idx].harga {
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

func SelectionSortStrAsc(A *kripto, n int) {
	var i, idx, pass int
	var temp Akripto

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nama > A[idx].nama {
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

func insertionSortIntAsc(A *kripto, n int) {
	var pass, i int
	var temp Akripto

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.harga < A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func Beli(A *arrAkun, B kripto, idxAcc, nkripto int) {
	var jumlah, idx int
	var beliKripto, x string
	x = "y"
	clearScreen()
	for x == "y" {
		for {
			fmt.Print("Masukkan nama kripto yang akan dibeli: ")
			fmt.Scan(&beliKripto)
			idx = sequentialSearchStr(B, nkripto, beliKripto)
			if idx != -1 {
				break
			} else {
				fmt.Println("Kripto tidak ditemukan!")
			}
		}
		fmt.Print("Masukkan jumlah kripto yang akan dibeli: ")
		fmt.Scan(&jumlah)

		A[idxAcc].kripto[A[idxAcc].nkriptoAset].nama = beliKripto
		A[idxAcc].kripto[A[idxAcc].nkriptoAset].jumlah = float64(jumlah) / float64(B[idx].harga)
		A[idxAcc].aset += A[idxAcc].kripto[A[idxAcc].nkriptoAset].jumlah * float64(B[idx].harga)
		A[idxAcc].nkriptoAset++
		fmt.Print("Ingin melakukan transaksi lagi (y/n)?")
		fmt.Scan(&x)
	}
}

func lihatPortofolio(A arrAkun, idx int) {
	var i int
	clearScreen()
	fmt.Println("-----------------------------------------")
	fmt.Println("|              Portofolio               |")
	fmt.Println("-----------------------------------------")
	for i = 0; i < A[idx].nkriptoAset; i++ {
		fmt.Printf("%d. %s - Jumlah: %f\n", i+1, A[idx].kripto[i].nama, A[idx].kripto[i].jumlah)
	}
	fmt.Printf("Total Aset: Rp %d\n", int(A[idx].aset))
	fmt.Println("-----------------------------------------")
	enterKembali()
}
