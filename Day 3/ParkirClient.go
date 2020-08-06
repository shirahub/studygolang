package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	ParkID    string `json:"parkid"`
	Kendaraan string `json:"kendaraan"`
	PlatNo    string `json:"platno"`
}

func main() {
	http.HandleFunc("/getParkID", parkHandler)
	http.HandleFunc("/parkOut", outHandler)

	fmt.Printf("Starting client at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

func outHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//handler jika url path salah
	if r.URL.Path != "/parkOut" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	//handler jika method salah
	if r.Method != "POST" {
		http.Error(w, "400 method not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		var client = &http.Client{}
		//baca body request
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newReqToServer, err := http.NewRequest("POST", "http://localhost:8082/keluar", bytes.NewBuffer(reqBody))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := client.Do(newReqToServer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		//baca body response dari server
		serverResp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(serverResp)
		return
	}
}

func parkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//handler jika url path salah
	if r.URL.Path != "/getParkID" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	//handler jika method salah
	if r.Method != "GET" {
		http.Error(w, "400 method not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		resp, err := http.Get("http://localhost:8082/generateID")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		responseForPostman, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//buat response menjadi JSONObject
		generatedID := parkID{Status: resp.Status, ParkID: string(responseForPostman)}
		responsejs, err := json.Marshal(generatedID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(responsejs)
		return
	}

}
