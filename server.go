package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ClientInfo struct {
	Platform string  `json:"platform"`
	Language string  `json:"language"`
	IP       string  `json:"IP"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

const (
	reset = "\033[0m"
	blue  = "\033[34m"
	green = "\033[32m"
	yellow = "\033[33m"
	cyan = "\033[36m"
	bold = "\033[1m"
)

func printClientInfo(info ClientInfo) {
	log.Println(bold + cyan + "=== Incoming Client Info ===" + reset)
	log.Printf("%sPlatform:%s %s\n", green, reset, info.Platform)
	log.Printf("%sLanguage:%s %s\n", green, reset, info.Language)
	log.Printf("%sIP Address:%s %s\n", green, reset, info.IP)
	log.Printf("%sLatitude:%s %.6f\n", green, reset, info.Lat)
	log.Printf("%sLongitude:%s %.6f\n", green, reset, info.Lon)
	log.Println(bold + cyan + "=============================" + reset)
}


func reciever(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var info ClientInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Bad JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Print the info to the terminal in a nice format
	printClientInfo(info)

}

func serveFrontEnd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
		return
	}
	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)

}

func main() {

	http.HandleFunc("/", serveFrontEnd)
	http.HandleFunc("/send", reciever)
	log.Println(bold + yellow + "🚀 Waiting for client data on http://localhost:6969/send ..." + reset)
	http.ListenAndServe(":6969", nil)
}
