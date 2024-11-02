package radio

import (
	"testing"
)

func TestFetchStations(t *testing.T) {
	FetchStations(StationsByCountry, "italy")
}

func TestGetStationsByTag(t *testing.T) {
	GetStationsByTag("24-hour punk")
}

func TestFetchAllStations(t *testing.T) {
	FetchAllStations()

	// for _, s := range allS {
	// 	log.Println(s.Tags)
	// }
}
