package radio

import (
	"testing"
)

func TestFetchStations(t *testing.T) {
	FetchStations(StationsByCountry, "italy")

	// for _, station := range st {
	// 	log.Println(station)
	// 	for _, tag := range station.Tags {
	// 		log.Println(tag.GetValues())
	// 	}
	// }
}

func TestGetStationsByTag(t *testing.T) {
	GetStationsByTag("top 100")
	// log.Println(st)
}
