package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/FG420/web-radio/radio"
)

// type Data struct {
// 	Station *radio.Station
// }

func handleHome(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))

	templ.Execute(w, nil)
}

func main() {

	stations := radio.GetStationsByCountry("italy")

	log.Println(stations)

	http.HandleFunc("/", handleHome)

	http.ListenAndServe(":8080", nil)

}
