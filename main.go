package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/FG420/web-radio/radio"
)

// type Data struct {
// 	Stations []*Station
// }

type Station struct {
	ID      string
	Name    string
	Url     string
	Country string
	Tags    string
	Favicon string
}

func NewStation(id, name, url, country, tags, favicon string) *Station {
	return &Station{
		ID:      id,
		Name:    name,
		Url:     url,
		Country: country,
		Tags:    tags,
		Favicon: favicon,
	}
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	// templ := template.Must(template.ParseFiles("views/index.html"))
	// templ.Execute(w, nil)

	templ := template.Must(template.ParseFiles("views/data.html"))
	getStations := radio.GetStationsByCountry("italy")

	var s *Station
	for i := 0; i < len(getStations); i++ {
		ss := NewStation(getStations[i].StationUUID, getStations[i].Name, getStations[i].URL,
			getStations[i].Country, getStations[i].Tags, getStations[i].Favicon)
		s = ss
	}

	log.Println(s)

	templ.ExecuteTemplate(w, "stationData", s)
}

func main() {

	http.HandleFunc("/", handleHome)

	http.ListenAndServe(":8080", nil)

}
