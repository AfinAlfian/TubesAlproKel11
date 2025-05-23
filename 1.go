package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const NMAX int = 1024

type arrAkun [NMAX]akun
type arrKripto [NMAX]kripto
type arrRiwayat [NMAX]riwayat

type akun struct {
	username    string
	password    string
	recovery    string
	saldo       float64
	aset        float64
	nAsetKripto int
	nRiwayat    int
	arrKripto
	arrRiwayat
}

type kripto struct {
	nama        string
	harga       float64
	jumlah      float64
	totalKripto int
}

type riwayat struct {
	jenis  string
	nama   string
	jumlah float64
	waktu  string
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
	var k arrKripto
	var b, nAkun, idxAcc int
	var nKripto int = 10
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

	menuAkun(&a, &nAkun, &idxAcc, &exit)
	if exit {
		return
	}
	for {
		clearScreen()
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Menu Utama               |")
		fmt.Println("-----------------------------------------")
		fmt.Println("username:", a[idxAcc].username)
		hitungAset(&a, k, idxAcc, nKripto)
		fmt.Println("Total Aset: Rp", int(a[idxAcc].aset))
		fmt.Println("Saldo: Rp", int(a[idxAcc].saldo))
		fmt.Println("-----------------------------------------")
		for i := 0; i < 3; i++ {
			fmt.Println(i+1, k[i].nama, "Harga:", int(k[i].harga))
		}
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Lihat Daftar Kripto")
		fmt.Println("2. Cari Kripto")
		fmt.Println("3. Edit Kripto")
		fmt.Println("4. Lihat Portofolio")
		fmt.Println("5. Deposit")
		fmt.Println("6. Withdraw")
		fmt.Println("7. Lihat Riwayat Transaksi")
		fmt.Println("8. Log Out")
		fmt.Println("9. Keluar Program")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&b)

		switch b {
		case 1:
			lihatKripto(k, nKripto)
			enterKembali()
		case 2:
			cariKriptoSeq(k, nKripto)
		case 3:
			menuKripto(&k, &nKripto)
		case 4:
			portofolio(&a, &k, idxAcc, nKripto)
		case 5:
			deposit(&a, idxAcc)
		case 6:
			withdraw(&a, idxAcc)
		case 7:
			LihatRiwayatTransaksi(&a, idxAcc)
		case 8:
			menuAkun(&a, &nAkun, &idxAcc, &exit)
			if exit {
				return
			}
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid")
			enterKembali()
		}
	}

}

func menuAkun(A *arrAkun, nAkun, idx *int, exit *bool) {
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
			enterKembali()
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
	var user string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&user)
	for i := 0; i < *n; i++ {
		if A[i].username == user {
			fmt.Println("Username sudah terdaftar!")
			enterKembali()
			return
		}
	}
	A[*n].username = user
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

func menuKripto(k *arrKripto, n *int) {
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
		fmt.Println("4. Kembali ke Menu Utama")
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
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func lihatKripto(k arrKripto, n int) {
	//var dummy string
	clearScreen()
	fmt.Println("-----------------------------------------")
	fmt.Println("|              Daftar Kripto            |")
	fmt.Println("-----------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, k[i].nama, "Harga:", int(k[i].harga))
	}
	fmt.Println("-----------------------------------------")
}

func inputKripto(k *arrKripto, n *int) {
	var x, nama string
	var idx int

	for {
		fmt.Print("Masukkan nama kripto: ")
		fmt.Scan(&nama)
		idx = sequentialSearchStr(*k, *n, nama)
		if idx != -1 {
			fmt.Println("Kripto sudah terdaftar!")
		} else {
			k[*n].nama = nama
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
		}
		fmt.Print("Apakah ingin menambahkan lagi(y/n)? ")
		fmt.Scan(&x)
		fmt.Println("-----------------------------------------")
		if x == "n" {
			break
		}
	}
}

func hapusKripto(k *arrKripto, n *int, x string) {
	var i, j int
	for i = 0; i < *n; i++ {
		if k[i].nama == x {
			for j = i; j < *n-1; j++ {
				k[j] = k[j+1]
			}
			*n--
		}
	}
	fmt.Printf("Kripto \"%s\" sudah terhapus.\n", x)
	enterKembali()
}

func cariKriptoBin(k *arrKripto, n int) {
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

func urutkanKripto(k *arrKripto, n int) {
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

func cariKriptoSeq(k arrKripto, n int) {
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
			fmt.Println("Nama:", k[idx].nama, "Harga:", int(k[idx].harga))
		} else {
			fmt.Println("Kripto tidak ditemukan")
		}
	} else if y == "nama" {
		fmt.Print("Masukkan nama kripto yang ingin dicari: ")
		fmt.Scan(&search2)
		idx = sequentialSearchStr(k, n, search2)
		if idx != -1 {
			fmt.Println("Kripto ditemukan di index:", idx)
			fmt.Println("Nama:", k[idx].nama, "Harga:", int(k[idx].harga))
		} else {
			fmt.Println("Kripto tidak ditemukan")
		}
	}
	enterKembali()
}

func binarySearchStrAsc(A arrKripto, n int, x string) int {
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

func sequentialSearchInt(A arrKripto, n, x int) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if int(A[i].harga) == x {
			idx = i
		}
	}
	return idx
}

func sequentialSearchStr(A arrKripto, n int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].nama == x {
			idx = i
		}
	}
	return idx
}

func SelectionSortIntDes(A *arrKripto, n int) {
	var i, idx, pass int
	var temp kripto

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

func SelectionSortStrAsc(A *arrKripto, n int) {
	var i, idx, pass int
	var temp kripto

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

func insertionSortIntAsc(A *arrKripto, n int) {
	var pass, i int
	var temp kripto

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

func Beli(A *arrAkun, B arrKripto, idxAcc, nkripto int) {
	var idx int
	var beliKripto, x string
	var jumlah float64

	for {
		fmt.Print("Masukkan nama kripto yang akan dibeli: ")
		fmt.Scan(&beliKripto)
		idx = sequentialSearchStr(B, nkripto, beliKripto)
		if idx == -1 {
			fmt.Println("Kripto tidak ditemukan!")
			enterKembali()
			return
		}

		fmt.Print("Masukkan jumlah harga kripto yang akan dibeli: ")
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println("Jumlah tidak boleh 0 atau negatif")
			enterKembali()
			return
		} else {
			if A[idxAcc].saldo < jumlah {
				fmt.Println("Saldo tidak cukup")
				fmt.Println("Saldo anda: Rp", int(A[idxAcc].saldo))
				fmt.Println("Silahkan isi saldo terlebih dahulu")
				enterKembali()
				return
			}
			A[idxAcc].arrKripto[A[idxAcc].nAsetKripto].nama = beliKripto
			A[idxAcc].arrKripto[A[idxAcc].nAsetKripto].jumlah = float64(jumlah) / float64(B[idx].harga)
			A[idxAcc].saldo -= jumlah
			A[idxAcc].nAsetKripto++
			randomHarga(&B, nkripto)

			A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jenis = "Beli"
			A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].nama = beliKripto
			A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jumlah = jumlah
			A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].waktu = time.Now().Format("2006-01-02 15:04")
			A[idxAcc].nRiwayat++

			fmt.Print("Ingin melakukan transaksi lagi (y/n)? ")
			fmt.Scan(&x)
			if x == "n" {
				return
			}
		}
	}
}

func portofolio(A *arrAkun, B *arrKripto, idxAcc, nKripto int) {
	var i, a int
	for {
		clearScreen()
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Portofolio               |")
		fmt.Println("-----------------------------------------")
		for i = 0; i < A[idxAcc].nAsetKripto; i++ {
			fmt.Printf("%d. %s - Jumlah: %f\n", i+1, A[idxAcc].arrKripto[i].nama, A[idxAcc].arrKripto[i].jumlah)
		}
		fmt.Printf("Saldo: Rp %d\n", int(A[idxAcc].saldo))
		hitungAset(A, *B, idxAcc, nKripto)
		fmt.Printf("Total Aset: Rp %d\n", int(A[idxAcc].aset))
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Beli Kripto")
		fmt.Println("2. Jual Kripto")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&a)
		switch a {
		case 1:

			Beli(A, *B, idxAcc, nKripto)
		case 2:
			jual(A, B, idxAcc, nKripto)

		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func deposit(A *arrAkun, idxAcc int) {
	var depo float64
	fmt.Print("Masukkan jumlah yang ingin didepositkan: ")
	fmt.Scan(&depo)
	A[idxAcc].saldo += depo

	A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jenis = "Deposit"
	A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jumlah = depo
	A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].waktu = time.Now().Format("2006-01-02 15:04")
	A[idxAcc].nRiwayat++
	fmt.Println("Deposit berhasil!")
	enterKembali()
}

func withdraw(A *arrAkun, idxAcc int) {
	var narik float64
	fmt.Print("Masukkan jumlah yang ingin diambil: ")
	fmt.Scan(&narik)
	if A[idxAcc].saldo < narik {
		fmt.Println("Saldo tidak cukup")
		fmt.Println("Saldo anda: Rp", int(A[idxAcc].saldo))
		enterKembali()
	} else {
		A[idxAcc].saldo -= narik
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jenis = "Withdraw"
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jumlah = narik
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].waktu = time.Now().Format("2006-01-02 15:04")
		A[idxAcc].nRiwayat++
		fmt.Println("Withdraw berhasil!")
		enterKembali()
	}
}

func jual(A *arrAkun, B *arrKripto, idxAcc, nKripto int) {
	var jumlah float64
	var kripto string
	var idxKripto, idxKriptoJual int
	fmt.Print("Masukkan nama kripto yang ingin dijual: ")
	fmt.Scan(&kripto)
	idxKripto = sequentialSearchStr(*B, nKripto, kripto)
	idxKriptoJual = sequentialSearchStr(A[idxAcc].arrKripto, A[idxAcc].nAsetKripto, kripto)
	if idxKriptoJual != -1 {
		fmt.Print("Masukkan jumlah yang ingin dijual: ")
		fmt.Scan(&jumlah)
		A[idxAcc].arrKripto[idxKriptoJual].jumlah -= jumlah
		A[idxAcc].saldo += jumlah * float64(B[idxKripto].harga)
		randomHarga(B, nKripto)
		if A[idxAcc].arrKripto[idxKriptoJual].jumlah <= 0 {
			hapusAsetKripto(A, idxAcc, idxKriptoJual)
		}

		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jenis = "Jual"
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].nama = kripto
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].jumlah = jumlah
		A[idxAcc].arrRiwayat[A[idxAcc].nRiwayat].waktu = time.Now().Format("2006-01-02 15:04")
		A[idxAcc].nRiwayat++
	} else {
		fmt.Println("Kripto tidak ditemukan!")
	}
	fmt.Println("Jual berhasil!")
	enterKembali()
}

func hitungAset(A *arrAkun, B arrKripto, idxAcc, nKripto int) {
	var i, idxKripto int
	A[idxAcc].aset = 0
	for i = 0; i < A[idxAcc].nAsetKripto; i++ {
		idxKripto = sequentialSearchStr(B, nKripto, A[idxAcc].arrKripto[i].nama)
		A[idxAcc].aset += A[idxAcc].arrKripto[i].jumlah * B[idxKripto].harga
	}
}

func hapusAsetKripto(A *arrAkun, idxAcc int, idxKriptoJual int) {
	var i, j int
	for i = 0; i < A[idxAcc].nAsetKripto; i++ {
		if A[idxAcc].arrKripto[i].nama == A[idxAcc].arrKripto[idxKriptoJual].nama {
			for j = i; j < A[idxAcc].nAsetKripto-1; j++ {
				A[idxAcc].arrKripto[j] = A[idxAcc].arrKripto[j+1]
			}
			A[idxAcc].nAsetKripto--
		}
	}
}

func randomHarga(A *arrKripto, nKripto int) {
	var i int
	for i = 0; i < nKripto; i++ {
		A[i].harga += A[i].harga * (float64(rand.Intn(151)-50) / 100)
	}
}

func LihatRiwayatTransaksi(A *arrAkun, idxAcc int) {
	var i int
	clearScreen()
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("|                           Riwayat Transaksi Akun                            |")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-13s | %-12s | %-18s | %-17s |\n", "No", "Jenis", "Nama", "Jumlah", "Waktu")
	fmt.Println("-------------------------------------------------------------------------------")
	for i = 0; i < A[idxAcc].nRiwayat; i++ {
		fmt.Printf("| %3d | %-13s | %-12s | %-18.0f | %-17s |\n",
			i+1,
			A[idxAcc].arrRiwayat[i].jenis,
			A[idxAcc].arrRiwayat[i].nama,
			A[idxAcc].arrRiwayat[i].jumlah,
			A[idxAcc].arrRiwayat[i].waktu)
	}
	fmt.Println("-------------------------------------------------------------------------------")
	enterKembali()
}
