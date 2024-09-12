package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/FG420/web-radio/radio"
)

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
	templ := template.Must(template.ParseFiles("views/home.html"))
	templ.Execute(w, nil)

	country := req.FormValue("country")

	log.Println("home -> ", country)
}

func handleGetStationsByCountry(w http.ResponseWriter, req *http.Request) {

	country := req.PathValue("country")

	log.Println("stations -> ", country)

	templ := template.Must(template.ParseFiles("views/stations.html"))
	getStations := radio.GetStationsByCountry(country)

	var s []Station

	for i := 0; i < len(getStations); i++ {
		ss := NewStation(getStations[i].StationUUID, getStations[i].Name, getStations[i].URL,
			getStations[i].Country, getStations[i].Tags, getStations[i].Favicon)

		s = append(s, *ss)
	}

	log.Println("\n", s)

	templ.ExecuteTemplate(w, "stations", s)
}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/:country", handleGetStationsByCountry)

	http.ListenAndServe(":8080", nil)

}
