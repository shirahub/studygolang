package main

import (
	pb "Day4/protofolder"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//Mobil dengan struct
type Mobil struct {
	PlatNo string
	ParkID string
}

//Motor dengan struct
type Motor struct {
	PlatNo string
	ParkID string
}

//Nota dengan struct
type Nota struct {
	Masuk     string `json:"masuk"`
	Keluar    string `json:"keluar"`
	Kendaraan string `json:"kendaraan"`
	PlatNo    string `json:"platno"`
	Biaya     int32  `json:"biaya"`
}

//Parker dengan struct
type Parker struct {
	ParkID    string `json:"parkid"`
	Kendaraan string `json:"kendaraan"`
	PlatNo    string `json:"platno"`
}

// server is used to implement parkir.ConnectServer.
type server struct {
	pb.UnimplementedConnectServer
}

func (s *server) ConnectToServer(ctx context.Context, void *empty.Empty) (*pb.ServerReply, error) {
	return &pb.ServerReply{Message: "Client is connected to Server"}, nil
}

func (s *server) GetParkID(ctx context.Context, void *empty.Empty) (*pb.Parkid, error) {
	generatedID := parkirMasuk()
	return &pb.Parkid{Parkid: generatedID}, nil
}
func (s *server) GetParkFee(ctx context.Context, parker *pb.Parker) (*pb.Nota, error) {
	biayaparkir, waktumasuk, waktukeluar := parkirKeluar(parker.GetKendaraan(), parker.GetPlatno(), parker.GetParkid())
	var result = pb.Nota{Masuk: waktumasuk, Keluar: waktukeluar, Kendaraan: parker.GetKendaraan(), Platno: parker.GetPlatno(), Biaya: biayaparkir}
	return &result, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterConnectServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func generateID(w http.ResponseWriter, r *http.Request) {
	generatedID := parkirMasuk()
	fmt.Fprintf(w, generatedID)
}

func parkirKeluar(kendaraan string, platNo string, parkID string) (int32, string, string) {

	if kendaraan == "mobil" && parkID != "" {
		tempMobil := Mobil{PlatNo: platNo, ParkID: parkID}
		return tempMobil.hitungParkirMobil()
	} else if kendaraan == "motor" && parkID != "" {
		tempMotor := Motor{PlatNo: platNo, ParkID: parkID}
		return tempMotor.hitungParkirMotor()
	} else {
		fmt.Println("Tipe kendaraan invalid")
		return 0, "input invalid", "input invalid"
	}
}

func (mot Motor) hitungParkirMotor() (int32, string, string) {
	waktukeluar := time.Now()
	parseJam, _ := strconv.Atoi(mot.ParkID[4:6])
	parseMenit, _ := strconv.Atoi(mot.ParkID[6:8])
	parseDetik, _ := strconv.Atoi(mot.ParkID[8:10])

	//Ganti waktu parkir sesuai hari
	waktumasuk := time.Date(
		2020, 8, 7, parseJam, parseMenit, parseDetik, 1, time.Local)
	diff := waktukeluar.Sub(waktumasuk)
	durasidetik, _ := time.ParseDuration(diff.String())
	fmt.Println("Anda parkir selama", int(durasidetik.Seconds()), "detik")
	biayaparkir := 3000 + ((int(durasidetik.Seconds()) - 1) * 2000)
	fmt.Println("Biaya parkir anda Rp", biayaparkir)
	return int32(biayaparkir), waktumasuk.String(), waktukeluar.String()
}

func (mob Mobil) hitungParkirMobil() (int32, string, string) {
	waktukeluar := time.Now()
	parseJam, _ := strconv.Atoi(mob.ParkID[4:6])
	parseMenit, _ := strconv.Atoi(mob.ParkID[6:8])
	parseDetik, _ := strconv.Atoi(mob.ParkID[8:10])

	//Ganti waktu parkir sesuai hari
	waktumasuk := time.Date(
		2020, 8, 7, parseJam, parseMenit, parseDetik, 1, time.Local)
	diff := waktukeluar.Sub(waktumasuk)
	durasidetik, _ := time.ParseDuration(diff.String())

	fmt.Println("Anda parkir selama", int(durasidetik.Seconds()), "detik")
	biayaparkir := 5000 + ((int(durasidetik.Seconds()) - 1) * 3000)
	fmt.Println("Biaya parkir anda Rp", biayaparkir)
	return int32(biayaparkir), waktumasuk.String(), waktukeluar.String()
}

//GENERATE park id. Data Jam, Menit, Detik diambil dari time.Now(), lalu dijadikan string
func parkirMasuk() string {
	fmt.Println("Generating your park id...")
	waktu := time.Now()
	var jam, menit, detik string
	waktustring := waktu.String()
	jam = waktustring[11:13]
	menit = waktustring[14:16]
	detik = waktustring[17:19]
	parkID := randomString(4) + jam + menit + detik
	fmt.Println(parkID)
	return parkID
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
