package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "Day4/protofolder"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

//parkID
type parkID struct {
	Status string
	ParkID string
}
type error struct {
	Status  string
	Message string
}

//Parker dengan struct
type Parker struct {
	Parkid    string `json:"parkid"`
	Kendaraan string `json:"kendaraan"`
	Platno    string `json:"platno"`
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	menu()
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
		menu()
	}
}

func parkirMasuk() {
	// Set up a connection to the server.
	user := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getPark, err := user.GetParkID(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Your Park Id: %s", getPark.GetParkid())
	menu()
}

//ConnectServer untuk connect ke proto
func ConnectServer() pb.ConnectClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	return pb.NewConnectClient(conn)
}

func parkirKeluar() {
	// Set up a connection to the server.
	user := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var kendaraan, parkID, platNo string
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

	parker := pb.Parker{Parkid: parkID, Kendaraan: kendaraan, Platno: platNo}

	getFee, err := user.GetParkFee(ctx, &parker)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	prettyPrint(getFee)
	log.Printf(prettyPrint(getFee))
	menu()
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
