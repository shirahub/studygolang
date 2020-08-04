package main

import (
	"fmt"
	"time"
)

func main() {
	//KALENDER PER TAHUN. GANTI TAHUN DI VARIABEL tahun
	tahun := 1994
	jumlahhari := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	namabulan := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	haripertamadarisetiapbulan := [12]int{}
	empty := "-- "
	var haripertama int

	//cek tahun kabisat
	if (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0 {
		jumlahhari[1] = 29
	}

	//loop untuk menentukan hari pertama dari setiap bulan. Menggunakan time
	for d := 0; d < 12; d++ {
		hari := time.Date(tahun, time.Month(d+1), 1, 1, 1, 1, 1, time.UTC)

		switch hari.Weekday() {
		case 1:
			haripertama = 1
		case 2:
			haripertama = 2
		case 3:
			haripertama = 3
		case 4:
			haripertama = 4
		case 5:
			haripertama = 5
		case 6:
			haripertama = 6
		case 7:
			haripertama = 7
		}
		haripertamadarisetiapbulan[d] = haripertama
	}

	fmt.Println("    KALENDER", tahun)

	//loop template per bulan
	for i := 0; i < 12; i++ {

		//agar ==== di penulisannya bulannya rapi aja
		fmt.Print("======" + namabulan[i])
		sisaspasi := 14 - len(namabulan[i])
		for a := 0; a < sisaspasi; a++ {
			fmt.Print("=")
		}
		fmt.Println("")
		fmt.Println("S  S  R  K  J  S  M ")

		//hitung start angka di hari apa, lalu isi hari kosong
		if haripertamadarisetiapbulan[i] > 0 {
			for k := 1; k < haripertamadarisetiapbulan[i]; k++ {
				fmt.Print(empty)
			}
		}

		//loop tanggal untuk isi kalender
		for j, h := 1, haripertamadarisetiapbulan[i]; j <= jumlahhari[i]; j, h = j+1, h+1 {
			if j < 10 {
				fmt.Print(j, "  ")
			} else {
				fmt.Print(j, " ")
			}

			//enter new line jika pembagian 7
			if h%7 == 0 {
				fmt.Print("\n")
			}
		}

		//hitung sisa hari, lalu isi sisa hari
		if haripertamadarisetiapbulan[i]+jumlahhari[i] <= 35 {
			sisahari := 35 - (haripertamadarisetiapbulan[i] + jumlahhari[i])
			for m := 0; m <= sisahari; m++ {
				fmt.Print(empty)
			}
		} else if haripertamadarisetiapbulan[i]+jumlahhari[i] > 36 {
			sisahari := 42 - (haripertamadarisetiapbulan[i] + jumlahhari[i])
			for m := 0; m <= sisahari; m++ {
				fmt.Print(empty)
			}
		}

		fmt.Println("")
	}

}
