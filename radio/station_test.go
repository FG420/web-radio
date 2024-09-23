package radio

import (
	"log"
	"testing"
)

func TestFetchStations(t *testing.T) {
	st := FetchStations(StationsByCountry, "italy")

	for _, station := range st {
		log.Println(station)
		for _, tag := range station.Tags {
			log.Println(tag.GetValues())
		}
	}
}

func TestGetStationsByTag(t *testing.T) {
	GetStationsByTag("rock")
	// log.Println(st)
}
