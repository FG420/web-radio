package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/FG420/web-radio/radio"
)

type (
	Component string

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

const (
	HEAD_URL   Component = "views/components/head.html"
	NAVBAR_URL Component = "views/components/navbar.html"
	FOOTER_URL Component = "views/components/footer.html"
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

func createTemplate(templatePath string) *template.Template {
	templ := template.Must(template.ParseFiles(templatePath, string(HEAD_URL),
		string(NAVBAR_URL), string(FOOTER_URL)))

	return templ
}

func handleHome(w http.ResponseWriter, req *http.Request) {
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

	createTemplate("views/home.html").Execute(w, countries)
}

func handleGetStationsByCountry(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	country := ""
	if len(parts) > 1 {
		country = parts[1]
	}

	getStations := radio.FetchStations(radio.StationsByCountry, country)
	var s []Station

	for _, station := range getStations {

		addS := NewStation(station.StationUUID, station.Name, station.URL,
			station.Country, station.Favicon, station.Tags)

		s = append(s, *addS)

	}

	createTemplate("views/components/stations.html").Execute(w, s)
}

func handleStationUrl(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	country, ID := "", ""

	if len(parts) > 1 {
		country = parts[1]
	}
	if len(parts) > 2 {
		ID = parts[2]
	}

	log.Println("country ->", country)
	log.Println("ID ->", ID)

	url := radio.GetStationUrl(country, ID)
	log.Print("url ->", url)

	fmt.Fprintf(w,
		`<div class="audio-controls">
			<audio controls>
				<source src="%s" type="audio/mp3">
				<source src="%s" type="audio/wav">
			</audio>
    	</div>`,
		url, url)
}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/{country}", handleGetStationsByCountry)
	http.HandleFunc("/{country}/{name}", handleStationUrl)

	http.ListenAndServe(":8080", nil)

}
