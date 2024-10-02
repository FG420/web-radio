package radio

import (
	"log"
	"testing"
)

func TestFetchStations(t *testing.T) {
	FetchStations(StationsByCountry, "italy")
}

func TestGetStationsByTag(t *testing.T) {
	GetStationsByTag("24-hour punk")
	// log.Println(st)
}

func TestFetchAllStations(t *testing.T) {
	allS := FetchAllStations()

	for _, s := range allS {
		log.Println(s.Tags)
	}
}
