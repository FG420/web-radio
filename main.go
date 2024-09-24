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

func handleHomePage(w http.ResponseWriter, req *http.Request) {
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

func handleStationsPage(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	country := ""
	if len(parts) > 1 {
		country = parts[1]
	}

	stations := radio.FetchStations(radio.StationsByCountry, country)
	log.Println(stations)
	for _, s := range stations {
		log.Println("all tags: -> ", s.Tags)
		for _, tag := range s.Tags {
			log.Println("tags value : -> ", tag.Name)
		}
	}

	createTemplate("views/components/stations.html").Execute(w, stations)
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
	url := radio.GetStationUrl(country, ID)

	fmt.Fprintf(w,
		`<div class="audio-controls">
			<audio controls>
				<source src="%s" type="audio/mp3">
				<source src="%s" type="audio/wav">
			</audio>
    	</div>`,
		url, url)
}

func handleTagsPage(w http.ResponseWriter, req *http.Request) {
	t := radio.GetTags()
	createTemplate("views/components/tags.html").Execute(w, t)
}

func handleSelectedTag(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	tag := ""

	if len(parts) > 1 {
		tag = parts[1]
	}

	t := radio.GetStationsByTag(tag)

	log.Print(t)
	log.Println("length ->", len(t))

	fmt.Fprint(w, t)

}

func main() {
	http.HandleFunc("/", handleHomePage)

	http.HandleFunc("/tags", handleTagsPage)
	http.HandleFunc("/tags/{tag}", handleSelectedTag)

	http.HandleFunc("/{country}", handleStationsPage)
	http.HandleFunc("/{country}/{name}", handleStationUrl)

	http.ListenAndServe(":8080", nil)

}
