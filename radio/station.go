package radio

import (
	json2 "encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	StationsURL = "https://de1.api.radio-browser.info/json/stations"
)

type (
	Data struct {
		ChangeUUID         string `json:"changeuuid"`
		StationUUID        string `json:"stationuuid"`
		Name               string `json:"name"`
		URL                string `json:"url"`
		URLResolved        string `json:"url_resolved"`
		Homepage           string `json:"homepage"`
		Favicon            string `json:"favicon"`
		Tags               string `json:"tags"`
		Country            string `json:"country"`
		CountryCode        string `json:"countrycode"`
		State              string `json:"state"`
		Language           string `json:"language"`
		Votes              string `json:"votes"`
		LastChangeTime     string `json:"lastchangetime"`
		Codec              string `json:"codec"`
		Bitrate            int    `json:"bitrate"`
		HLS                bool   `json:"hls"`
		LastCheckOk        bool   `json:"lastcheckok"`
		LastCheckTime      string `json:"lastchecktime"`
		LastCheckOkTime    string `json:"lastcheckoktime"`
		LastLocalCheckTime string `json:"lastlocalchecktime"`
		ClickTimestamp     string `json:"clicktimestamp"`
		ClickCount         int    `json:"clickcount"`
		ClickTrend         int    `json:"clicktrend"`
	}

	Station struct {
		ID      string
		Name    string
		Url     string
		Country string
		Favicon string
		Tags    []Tag
	}
)

func NewStation(id, name, url, country, favicon string, tags string) *Station {
	var ts []Tag

	aT := strings.Split(tags, ",")
	for _, value := range aT {
		newT := NewTag(value)
		ts = append(ts, *newT)
	}

	return &Station{
		ID:      id,
		Name:    name,
		Url:     url,
		Country: country,
		Favicon: favicon,
		Tags:    ts,
	}
}

func GetStationUrl(country, id string) string {
	stations := FetchStations(StationsByCountry, country)
	var url = ""

	for _, station := range stations {
		if station.ID != id {
			continue
		}
		url = station.Url
	}
	return url
}

func GetStationsByTag(tag string) []Station {
	getStations := FetchStations(StationsByTag, tag)
	return getStations
}

func FetchAllStations() []Station {
	res := Post(StationsURL, "", nil)
	return UnmarshalStations(res)
}

func FetchStations(by StationsBy, term string) []Station {
	res := Post(GenerateStationsURL(by, term), "", nil)
	return UnmarshalStations(res)
}

func FetchAllStationsDetailed(order StationsOrder, reverse bool, offset uint,
	limit uint, hideBroken bool) []Station {
	q := make(map[string]string)
	q["order"] = string(order)
	q["reverse"] = strconv.FormatBool(reverse)
	if offset > 0 {
		q["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit > 0 {
		q["limit"] = fmt.Sprintf("%d", limit)
	}
	q["hidebroken"] = strconv.FormatBool(hideBroken)
	res := Post(StationsURL, "", q)
	return UnmarshalStations(res)
}

func FetchStationsDetailed(by StationsBy, term string, order StationsOrder,
	reverse bool, offset uint, limit uint, hideBroken bool) []Station {
	q := make(map[string]string)
	q["order"] = string(order)
	q["reverse"] = strconv.FormatBool(reverse)
	if offset > 0 {
		q["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit > 0 {
		q["limit"] = fmt.Sprintf("%d", limit)
	}
	q["hidebroken"] = strconv.FormatBool(hideBroken)
	res := Post(GenerateStationsURL(by, term), "", q)
	return UnmarshalStations(res)
}

func UnmarshalStations(json string) []Station {
	var data []Data
	var stations []Station
	json2.Unmarshal([]byte(json), &data)

	for _, s := range data {
		if s.Tags == "" {
			s.Tags = "No tags available!"
		}

		addS := NewStation(s.StationUUID, s.Name, s.URL, s.Country, s.Favicon, s.Tags)
		stations = append(stations, *addS)
	}

	return stations
}

func GenerateStationsURL(by StationsBy, term string) string {
	return StationsURL + "/" + fmt.Sprintf("%s/%s", by, term)
}
