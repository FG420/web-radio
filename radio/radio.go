package radio

import (
	json2 "encoding/json"

	"gitlab.com/AgentNemo/goradios"
)

const (
	CountriesURL     = "https://de1.api.radio-browser.info/json/countries"
	CountriesCodeURL = "https://de1.api.radio-browser.info/json/countrycodes"
)

type Country struct {
	Name         string `json:"name"`
	ISO          string `json:"iso_3166_1"`
	StationCount int    `json:"stationcount"`
}

func GetCountriesNames() []Country {
	res := Post(CountriesURL, "", nil)
	return UnmarshalCountries(res)
}

func GetStationsByCountry(country string) []goradios.Station {
	return goradios.FetchStations(goradios.StationsByCountry, country)
}

// func GetCountriesDetailed()

func UnmarshalCountries(json string) []Country {
	var countries []Country
	json2.Unmarshal([]byte(json), &countries)
	return countries
}
