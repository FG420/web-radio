package radio

import (
	"gitlab.com/AgentNemo/goradios"
)

func GetStationsByCountry(country string) []goradios.Station {
	return goradios.FetchStations(goradios.StationsByCountry, country)
}
