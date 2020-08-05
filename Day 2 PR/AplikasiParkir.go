package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	menu()
}

//Mobil dengan struct
type Mobil struct {
	PlatNo int
	ParkID string
}

//Motor dengan struct
type Motor struct {
	PlatNo int
	ParkID string
}

func menu() {
	var input int
	fmt.Println("==========================")
	fmt.Println("Welcome to Parkleen System")
	fmt.Println("1. Parkir Masuk")
	fmt.Println("2. Parkir Keluar")
	fmt.Println("==========================")
	fmt.Println("Pilih menu :")
	fmt.Scan(&input)

	switch input {
	case 1:
		parkirMasuk()
	case 2:
		parkirKeluar()
	default:
		fmt.Println("Mohon input menu sesuai angka")
		main()
	}
}

func parkirKeluar() {
	var kendaraan, parkID string
	var platNo int
	fmt.Println("Input Tipe Kendaraan :")
	fmt.Scan(&kendaraan)

	//handle input kendaraan salah
	if kendaraan != "mobil" && kendaraan != "motor" {
		fmt.Println("Tipe kendaraan invalid")
		main()
	}

	fmt.Println("Input Nomor Plat :")
	fmt.Scan(&platNo)

	//ROOM FOR IMPROVEMENT : handle input platNo salah

	fmt.Println("Input ID Parkir :")
	fmt.Scan(&parkID)

	//ROOM FOR IMPROVEMENT : handle input parkID salah

	if kendaraan == "mobil" && parkID != "" {
		tempMobil := Mobil{PlatNo: platNo, ParkID: parkID}
		tempMobil.hitungParkirMobil()
	} else if kendaraan == "motor" && parkID != "" {
		tempMotor := Motor{PlatNo: platNo, ParkID: parkID}
		tempMotor.hitungParkirMotor()
	} else {
		fmt.Println("Tipe kendaraan invalid")
		main()
	}
}

func (mot Motor) hitungParkirMotor() {
	waktukeluar := time.Now()
	parseJam, _ := strconv.Atoi(mot.ParkID[4:6])
	parseMenit, _ := strconv.Atoi(mot.ParkID[6:8])
	parseDetik, _ := strconv.Atoi(mot.ParkID[8:10])
	waktumasuk := time.Date(
		2020, 8, 5, parseJam, parseMenit, parseDetik, 1, time.Local)
	diff := waktukeluar.Sub(waktumasuk)
	durasidetik, _ := time.ParseDuration(diff.String())

	fmt.Println("Anda parkir selama", int(durasidetik.Seconds()), "detik")
	biayaparkir := 3000 + ((int(durasidetik.Seconds()) - 1) * 2000)
	fmt.Println("Biaya parkir anda Rp", biayaparkir)
	main()
}

func (mob Mobil) hitungParkirMobil() {
	waktukeluar := time.Now()
	parseJam, _ := strconv.Atoi(mob.ParkID[4:6])
	parseMenit, _ := strconv.Atoi(mob.ParkID[6:8])
	parseDetik, _ := strconv.Atoi(mob.ParkID[8:10])
	waktumasuk := time.Date(
		2020, 8, 5, parseJam, parseMenit, parseDetik, 1, time.Local)
	diff := waktukeluar.Sub(waktumasuk)
	durasidetik, _ := time.ParseDuration(diff.String())

	fmt.Println("Anda parkir selama", int(durasidetik.Seconds()), "detik")
	biayaparkir := 5000 + ((int(durasidetik.Seconds()) - 1) * 3000)
	fmt.Println("Biaya parkir anda Rp", biayaparkir)
	main()
}

//GENERATE park id. Data Jam, Menit, Detik diambil dari time.Now(), lalu dijadikan string
func parkirMasuk() {
	fmt.Println("Generating your park id...")
	waktu := time.Now()
	var jam, menit, detik string
	waktustring := waktu.String()
	jam = waktustring[11:13]
	menit = waktustring[14:16]
	detik = waktustring[17:19]
	parkID := randomString(4) + jam + menit + detik
	fmt.Println(parkID)
	main()
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
