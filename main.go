package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/FG420/web-radio/radio"
)

type (
	Station struct {
		ID      string
		Name    string
		Url     string
		Country string
		Favicon string
		Tags    string
	}

	Country struct {
		Name         string
		ISO          string
		StationCount int
	}
)

func NewStation(id, name, url, country, favicon string, tags string) *Station {
	return &Station{
		ID:      id,
		Name:    name,
		Url:     url,
		Country: country,
		Favicon: favicon,
		Tags:    tags,
	}
}

func NewCountry(name, iso string, stations int) *Country {
	return &Country{
		Name:         name,
		ISO:          iso,
		StationCount: stations,
	}
}

// func SortStationByTag(tag string)[]*Station {return }

func handleHome(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("views/home.html", "views/components/head.html",
		"views/components/navbar.html", "views/components/footer.html"))

	c := radio.GetCountriesNames()
	var countries []Country
	countrySeen := make(map[string]bool)

	for _, country := range c {
		if country.Name == "" {
			continue
		}

		if _, exists := countrySeen[country.Name]; exists {
			continue
		}

		countrySeen[country.Name] = true
		addC := NewCountry(country.Name, country.ISO, country.StationCount)
		countries = append(countries, *addC)
	}

	templ.Execute(w, countries)
}

func handleGetStationsByCountry(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	country := ""
	if len(parts) > 1 {
		country = parts[1]
	}

	templ := template.Must(template.ParseFiles("views/components/head.html", "views/components/stations.html"))

	getStations := radio.FetchStations(radio.StationsByCountry, country)

	log.Println(getStations)

	var s []Station

	for _, station := range getStations {
		addS := NewStation(station.StationUUID, station.Name, station.URL,
			station.Country, station.Favicon, station.Tags)
		// log.Print(len(addS.Tags))
		log.Print(addS.Tags)
		s = append(s, *addS)

	}

	templ.ExecuteTemplate(w, "stations", s)
}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/{country}", handleGetStationsByCountry)

	http.ListenAndServe(":8080", nil)

}
