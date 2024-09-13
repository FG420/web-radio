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

const (
	MAXPAGE = 9
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	var count = 0
	templ := template.Must(template.ParseFiles("views/home.html", "views/components/navbar.html", "views/components/footer.html"))
	templ.Execute(w, nil)

	country := req.FormValue("country")
	// page := req.FormValue("page")

	log.Println("country -> ", country)

	if country == "" {
		return
	} else {
		handleGetStationsByCountry(w, country, MAXPAGE)
		count++
	}

}

func handleGetStationsByCountry(w http.ResponseWriter, country string, max int) {

	templ := template.Must(template.ParseFiles("views/components/stations.html"))
	getStations := radio.GetStationsByCountry(country)

	log.Println(len(getStations) / max)

	var s []Station

	for i := 0; i < len(getStations)/max; i++ {
		ss := NewStation(getStations[i].StationUUID, getStations[i].Name, getStations[i].URL,
			getStations[i].Country, getStations[i].Tags, getStations[i].Favicon)

		s = append(s, *ss)
	}

	templ.ExecuteTemplate(w, "stations", s)
}

func main() {
	// mux := http.NewServeMux()

	http.HandleFunc("/", handleHome)
	// mux.HandleFunc("/{country}", handleGetStationsByCountry)

	// http.HandleFunc("/{country}", handleCiao)

	http.ListenAndServe(":8080", nil)

}
