package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 Buat kalkulator dengan fungsi:
	- kali
	- bagi
	- tambah
	- kurang
	- akar
	- pangkat
	- luas persegi
	- luas lingkaran
	- volume tabung
	- volume balok
	- volume prisma
menggunkan:
	- struct & method
	- pointer
	- function
*/

//Persegi menggunakan struct
type Persegi struct {
	Sisi int
}

//Tabung menggunakan struct
type Tabung struct {
	Jarijari int
	Tinggi   int
}

//Lingkaran menggunakan struct
type Lingkaran struct {
	Jarijari int
}

//Balok menggunakan struct
type Balok struct {
	Panjang int
	Lebar   int
	Tinggi  int
}

//Prisma menggunakan struct, alas prisma bisa bermacam-macam
//di dalam kalkulator ini, alas prisma yang bisa dihitung adalah:
//persegi panjang (p*l), segitiga (a*t/2), trapesium((s1+s2)*t/2), dan belah ketupat (d1*d2/2)
type Prisma struct {
	LuasAlas float32
	Tinggi   float32
}

func main() {
	menu()
}

func menu() {
	var input int
	fmt.Println("==============================")
	fmt.Println("Welcome to Simpleen Calculator")
	fmt.Println("1. Perkalian")
	fmt.Println("2. Pembagian")
	fmt.Println("3. Penambahan")
	fmt.Println("4. Pengurangan")
	fmt.Println("5. Akar")
	fmt.Println("6. Pangkat")
	fmt.Println("7. Hitung Luas Persegi")
	fmt.Println("8. Hitung Luas Lingkaran")
	fmt.Println("9. Hitung Volume Tabung")
	fmt.Println("10. Hitung Volume Balok")
	fmt.Println("11. Hitung Volume Prisma")
	fmt.Println("12. Exit")
	fmt.Println("==============================")
	fmt.Println("Pilih menu :")
	fmt.Scan(&input)

	switch input {
	case 1:
		perkalian()
	case 2:
		pembagian()
	case 3:
		penambahan()
	case 4:
		pengurangan()
	case 5:
		pengakaran()
	case 6:
		pemangkatan()
	case 7:
		temppersegi := buatPersegi()
		temppersegi.hitungLuasPersegi()
	case 8:
		templingkaran := buatLingkaran()
		templingkaran.hitungLuasLingkaran()
	case 9:
		temptabung := buatTabung()
		temptabung.hitungVolumeTabung()
	case 10:
		tempbalok := buatBalok()
		tempbalok.hitungVolumeBalok()
	case 11:
		tempprisma := buatPrisma()
		tempprisma.hitungVolumePrisma()
	case 12:
		fmt.Println("Terima kasih sudah menggunakan Simpleen (Simple Shirleen) Calculator")
	default:
		fmt.Println("Mohon input menu sesuai angka")
		main()
	}
}

//-----------------------MENU 1-----------------------
//bisa memasukkan banyak angka. jika input berupa string selain "sudah", akan muncul error
//tetapi program tetap berjalan dan angka bukan string sebelumnya tetap disimpan
// 10*20*a akan menghasilkan 200 dengan error "Tidak boleh memasukkan huruf" ketika input a
func perkalian() {
	kumpulanangka := make([]int, 0)
	var angkastring string
	var angka int
	var error error
	fmt.Println("Input angka yang ingin anda kalikan, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)
		angka, error = strconv.Atoi(angkastring)
		if error != nil && angkastring != "sudah" {
			fmt.Println("Tidak boleh memasukkan huruf")
		} else {
			if angkastring == "sudah" {
				total := 1
				for _, v := range kumpulanangka {
					total = total * v
				}
				fmt.Println("Total Angka : ", total)
				main()
			} else {
				kumpulanangka = append(kumpulanangka, angka)
			}
		}
	}
}

//-----------------------MENU 2-----------------------
//bisa memasukkan banyak angka. jika input berupa string selain "sudah", akan muncul error
//tetapi program tetap berjalan dan angka bukan string sebelumnya tetap disimpan
//angka pertama yg diinput akan dibagi dengan angka-angka selanjutnya
// 100/5/a akan menghasilkan 20 dengan error "Tidak boleh memasukkan huruf" ketika input a
func pembagian() {
	kumpulanangka := make([]float32, 0)
	var angkastring string
	var angka int
	var error error
	fmt.Println("Input angka yang ingin anda bagi, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)
		angka, error = strconv.Atoi(angkastring)
		if error != nil && angkastring != "sudah" {
			fmt.Println("Tidak boleh memasukkan huruf")
		} else {
			if angkastring == "sudah" {
				total := kumpulanangka[0]
				for i := 1; i < len(kumpulanangka); i++ {
					total = total / kumpulanangka[i]
				}
				fmt.Println("Total Angka : ", total)
				main()
			} else {
				kumpulanangka = append(kumpulanangka, float32(angka))
			}
		}
	}
}

//-----------------------MENU 3-----------------------
//bisa memasukkan banyak angka. jika input berupa string selain "sudah", akan muncul error
//tetapi program tetap berjalan dan angka bukan string sebelumnya tetap disimpan
//10+20+a akan menghasilkan 30 dengan error "Tidak boleh memasukkan huruf" ketika input a
func penambahan() {
	kumpulanangka := make([]int, 0)
	var angkastring string
	var angka int
	var error error
	fmt.Println("Input angka yang ingin anda tambahkan, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)
		angka, error = strconv.Atoi(angkastring)
		if error != nil && angkastring != "sudah" {
			fmt.Println("Tidak boleh memasukkan huruf")
		} else {
			if angkastring == "sudah" {
				total := 0
				for _, v := range kumpulanangka {
					total += v
				}
				fmt.Println("Total Angka : ", total)
				main()
			} else {
				kumpulanangka = append(kumpulanangka, angka)
			}
		}
	}
}

//-----------------------MENU 4-----------------------
//bisa memasukkan banyak angka. jika input berupa string selain "sudah", akan muncul error
//tetapi program tetap berjalan dan angka bukan string sebelumnya tetap disimpan
//angka pertama yg diinput akan dikurangi dengan angka-angka selanjutnya
// 10-20-a akan menghasilkan -10 dengan error "Tidak boleh memasukkan huruf" ketika input a
func pengurangan() {
	kumpulanangka := make([]int, 0)
	var angkastring string
	var angka int
	var error error
	fmt.Println("Input angka yang ingin anda kurangkan, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)
		angka, error = strconv.Atoi(angkastring)
		if error != nil && angkastring != "sudah" {
			fmt.Println("Tidak boleh memasukkan huruf")
		} else {
			if angkastring == "sudah" {
				total := kumpulanangka[0]
				for i := 1; i < len(kumpulanangka); i++ {
					total = total - kumpulanangka[i]
				}
				fmt.Println("Total Angka : ", total)
				main()
			} else {
				kumpulanangka = append(kumpulanangka, angka)
			}
		}
	}
}

//-----------------------MENU 5-----------------------
func pengakaran() {
	var angka float64
	fmt.Println("Input angka yang ingin anda akarkan")
	fmt.Println("Masukkan angka :")
	fmt.Scan(&angka)
	fmt.Println("Akar dari", angka, "=", math.Sqrt(angka))
	main()
}

//-----------------------MENU 6-----------------------
func pemangkatan() {
	var angka, pangkat, gantipangkat float64
	pangkat = 2 //diberi nilai sementara, kemudian diganti dengan func change pointer value
	fmt.Println("Input angka yang ingin anda pangkatkan")
	fmt.Println("Masukkan angka :")
	fmt.Scan(&angka)
	fmt.Println("Masukkan pangkat :")
	fmt.Scan(&gantipangkat)
	change(&pangkat, gantipangkat)
	fmt.Println(angka, "^", pangkat, "=", math.Pow(angka, pangkat))
	main()
}

// hanya function untuk mencoba pointer (syarat tugas), function tambahan yang bersifat tidak perlu
func change(original *float64, value float64) {
	*original = value
}

//-----------------------MENU 7-----------------------
func buatPersegi() Persegi {
	var sisi int
	fmt.Println("Masukkan besar sisi persegi :")
	fmt.Scan(&sisi)
	return Persegi{Sisi: sisi}
}

func (persegibuatan Persegi) hitungLuasPersegi() {
	fmt.Println("Luas persegi :", (persegibuatan.Sisi * persegibuatan.Sisi))
	main()
}

//-----------------------MENU 8-----------------------
func buatLingkaran() Lingkaran {
	var jarijari int
	fmt.Println("Masukkan besar jari-jari lingkaran :")
	fmt.Scan(&jarijari)
	return Lingkaran{Jarijari: jarijari}
}

func (lingkaranbuatan Lingkaran) hitungLuasLingkaran() {
	const phi float32 = 3.14
	fmt.Println("Luas Lingkaran :", phi*float32(lingkaranbuatan.Jarijari*lingkaranbuatan.Jarijari))
	main()
}

//-----------------------MENU 9-----------------------
func buatTabung() Tabung {
	var jarijari, tinggi int
	fmt.Println("Masukkan besar jari-jari tabung :")
	fmt.Scan(&jarijari)
	fmt.Println("Masukkan besar tinggi tabung :")
	fmt.Scan(&tinggi)
	return Tabung{Jarijari: jarijari, Tinggi: tinggi}
}

func (tabungbuatan Tabung) hitungVolumeTabung() {
	const phi float32 = 3.14
	fmt.Println("Luas Tabung :", phi*float32(tabungbuatan.Jarijari*tabungbuatan.Jarijari*tabungbuatan.Tinggi))
	main()
}

//-----------------------MENU 10----------------------

func buatBalok() Balok {
	var panjang, lebar, tinggi int
	fmt.Println("Masukkan besar panjang balok :")
	fmt.Scan(&panjang)
	fmt.Println("Masukkan besar lebar balok :")
	fmt.Scan(&lebar)
	fmt.Println("Masukkan besar tinggi balok :")
	fmt.Scan(&tinggi)
	return Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}
}

func (balokbuatan Balok) hitungVolumeBalok() {
	fmt.Println("Volume Balok :", balokbuatan.Lebar*balokbuatan.Panjang*balokbuatan.Tinggi)
	main()
}

//-----------------------MENU 11-----------------------
//karena prisma bisa memiliki berbagai macam bentuk alas, untuk menentukan besar LuasAlas,
// maka dibuat menu untuk input bentuk alas dan unsur-unsurnya
func buatPrisma() Prisma {
	var angka, luasalas, tinggi float32

	fmt.Println("Apa bentuk alas prisma anda?")
	fmt.Println("1. Persegi Panjang")
	fmt.Println("2. Segitiga")
	fmt.Println("3. Trapesium")
	fmt.Println("4. Belah Ketupat")
	fmt.Scan(&angka)
	switch angka {
	case 1:
		var panjang, lebar float32
		fmt.Println("Masukkan besar panjang persegi panjang :")
		fmt.Scan(&panjang)
		fmt.Println("Masukkan besar lebar persegi panjang :")
		fmt.Scan(&lebar)
		luasalas = panjang * lebar
	case 2:
		var alas, tinggi float32
		fmt.Println("Masukkan besar alas segitiga :")
		fmt.Scan(&alas)
		fmt.Println("Masukkan besar tinggi segitiga :")
		fmt.Scan(&tinggi)
		luasalas = alas * tinggi / 2
	case 3:
		var sisi1, sisi2, tinggi float32
		fmt.Println("Masukkan besar sisi sejajar satu trapesium :")
		fmt.Scan(&sisi1)
		fmt.Println("Masukkan besar sisi sejajar dua trapesium :")
		fmt.Scan(&sisi2)
		fmt.Println("Masukkan besar tinggi trapesium :")
		fmt.Scan(&tinggi)
		luasalas = (sisi1 + sisi2) * tinggi / 2
	case 4:
		var diagonal1, diagonal2 float32
		fmt.Println("Masukkan besar diagonal satu belah ketupat :")
		fmt.Scan(&diagonal1)
		fmt.Println("Masukkan besar diagonal dua belah ketupat :")
		fmt.Scan(&diagonal2)
		luasalas = diagonal1 * diagonal2 / 2
	default:
		fmt.Println("Mohon input bentuk alas yang diinginkan")
		buatPrisma()
	}

	fmt.Println("Masukkan tinggi prisma :")
	fmt.Scan(&tinggi)

	return Prisma{LuasAlas: luasalas, Tinggi: tinggi}
}

func (prismabuatan Prisma) hitungVolumePrisma() {
	fmt.Println("Volume Prisma :", prismabuatan.LuasAlas*prismabuatan.Tinggi)
	main()
}
