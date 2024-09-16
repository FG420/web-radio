package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/FG420/web-radio/radio"
)

type (
	Station struct {
		ID      string
		Name    string
		Url     string
		Country string
		Tags    string
		Favicon string
	}

	Country struct {
		Name         string
		ISO          string
		StationCount int
	}
)

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

func NewCountry(name, iso string, stations int) *Country {
	return &Country{
		Name:         name,
		ISO:          iso,
		StationCount: stations,
	}
}

const (
	MAXPAGE = 9
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("views/home.html",
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

	country := req.FormValue("country")

	log.Println("country -> ", country)
}

func handleGetStationsByCountry(w http.ResponseWriter, country string, max int) {

	templ := template.Must(template.ParseFiles("views/components/stations.html"))
	getStations := radio.GetStationsByCountry(country)

	log.Println(len(getStations) / max)

	var s []Station

	for _, station := range getStations {
		addS := NewStation(station.StationUUID, station.Name, station.URL,
			station.Country, station.Tags, station.Favicon)
		s = append(s, *addS)
	}

	templ.ExecuteTemplate(w, "stations", s)
}

// func LoadMoreHandler(w http.ResponseWriter, req *http.Request) {
// 	page := req.URL.Query().Get("page")

// 	intPage, err := strconv.Atoi(page)
// 	if err != nil {
// 		http.Error(w, "Invalid page number", http.StatusBadRequest)
// 		return
// 	}

// 	intPage++
// }

func main() {
	// mux := http.NewServeMux()

	http.HandleFunc("/", handleHome)
	// mux.HandleFunc("/{country}", handleGetStationsByCountry)

	// http.HandleFunc("/{country}", handleCiao)

	http.ListenAndServe(":8080", nil)

}
