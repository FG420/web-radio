package main

import (
	"log"

	"github.com/FG420/web-radio/radio"
)

// type Data struct {
// 	Station *radio.Station
// }

func main() {

	stations := radio.GetStationsByCountry("italy")

	log.Println(stations)
}
