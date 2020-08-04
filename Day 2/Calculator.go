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

func main() {
	menu()
}

func menu() {
	input := 1

	fmt.Println("Welcome to Simpleen Calculator")
	fmt.Println("1. Perkalian")
	fmt.Println("3. Penambahan")
	fmt.Println("5. Akar")
	fmt.Println("6. Pangkat")
	fmt.Println("7. Hitung Luas Persegi")
	fmt.Println("9. Hitung Volume Tabung")

	fmt.Println("Pilih menu :")
	fmt.Scan(&input)

	switch input {
	case 1:
		perkalian()
	case 3:
		penambahan()
	case 5:
		pengakaran()
	case 6:
		pemangkatan()
	case 7:
		temppersegi := buatPersegi()
		temppersegi.hitungLuasPersegi()
	case 9:
		temptabung := buatTabung()
		temptabung.hitungVolumeTabung()
	}

}

func (tabungbuatan Tabung) hitungVolumeTabung() {
	const phi float32 = 3.14
	fmt.Println("Luas Tabung :", phi*float32(tabungbuatan.Jarijari*tabungbuatan.Jarijari*tabungbuatan.Tinggi))
	main()
}

func buatTabung() Tabung {
	var jarijari, tinggi int
	fmt.Println("Masukkan besar jari-jari tabung :")
	fmt.Scan(&jarijari)
	fmt.Println("Masukkan besar tinggi tabung :")
	fmt.Scan(&tinggi)
	return Tabung{Jarijari: jarijari, Tinggi: tinggi}
}

func (persegibuatan Persegi) hitungLuasPersegi() {
	fmt.Println("Luas persegi :", (persegibuatan.Sisi * persegibuatan.Sisi))
	main()
}

func buatPersegi() Persegi {
	var sisi int
	fmt.Println("Masukkan besar sisi persegi :")
	fmt.Scan(&sisi)
	return Persegi{Sisi: sisi}
}

func penambahan() {
	kumpulanangka := make([]int, 1)
	var angkastring string
	var angka int
	fmt.Println("Input angka yang ingin anda tambahkan, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)

		if angkastring == "sudah" {
			total := 0
			for _, v := range kumpulanangka {
				total += v
			}
			fmt.Println("Total Angka : ", total)
			main()
		}

		angka, _ = strconv.Atoi(angkastring)
		kumpulanangka = append(kumpulanangka, angka)
	}

}

func perkalian() {
	kumpulanangka := make([]int, 0)
	var angkastring string
	var angka int
	fmt.Println("Input angka yang ingin anda kalikan, ketik \"sudah\" jika selesai")
	for angkastring != "sudah" {
		fmt.Println("Masukkan angka :")
		fmt.Scan(&angkastring)
		angka, _ = strconv.Atoi(angkastring)
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

func pemangkatan() {
	var angka, pangkat, gantipangkat float64
	pangkat = 2 //ceritanya menjadi default

	fmt.Println("Input angka yang ingin anda pangkatkan")
	fmt.Println("Masukkan angka :")
	fmt.Scan(&angka)
	fmt.Println("Masukkan pangkat :")
	fmt.Scan(&gantipangkat)
	change(&pangkat, gantipangkat)
	fmt.Println(angka, "^", pangkat, "=", math.Pow(angka, pangkat))
	main()
}

func change(original *float64, value float64) {
	*original = value
}

func pengakaran() {
	var angka float64
	fmt.Println("Input angka yang ingin anda akarkan")
	fmt.Println("Masukkan angka :")
	fmt.Scan(&angka)
	fmt.Println("Akar dari", angka, "=", math.Sqrt(angka))
	main()
}
