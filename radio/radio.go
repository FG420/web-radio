package radio

import (
	"gitlab.com/AgentNemo/goradios"
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

func GetStationsByCountry(country string) []*Station {
	getStations := goradios.FetchStations(goradios.StationsByCountry, country)

	var stations []*Station

	for i := 0; i < len(getStations); i++ {
		station := NewStation(getStations[i].StationUUID, getStations[i].Name, getStations[i].URL,
			getStations[i].Country, getStations[i].Tags, getStations[i].Favicon)
		stations = append(stations, station)
	}

	return stations
}
