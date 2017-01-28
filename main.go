// Copyright (c) 2017 Nicolas Stuhlfauth
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"strings"

	"strconv"

	"github.com/nicostuhlfauth/geoosm"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Point Contains lat and lon
type Point struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

// Standort Stores information about location of refuse container
type Standort struct {
	Coordinates Point  `json:"coordinates"`
	Stadtteil   string `json:"stadtteil"`
	Strasse     string `json:"strasse"`
	Papier      bool   `json:"papier"`
	Altkleider  bool   `json:"altkleider"`
	Glas        bool   `json:"glas"`
}

func main() {
	url, err := http.Get("http://abfallkalender.ebe-essen.de/containerstandorte-essen.php")
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(url.Body)

	matcher := scrape.ByTag(atom.Span)

	timeStart := time.Now()
	log.Println("Job started")

	scrapingResult := scrape.FindAll(root, matcher)

	standorte := make(map[string]*Standort)
	mapID := 0
	currentStadtteil := ""

	for i := range scrapingResult {
		if i > 4 {
			if !(strings.Contains(scrape.Text(scrapingResult[i]), ",")) {
				currentStadtteil = scrape.Text(scrapingResult[i])
			} else {
				temp := new(Standort)
				temp.Stadtteil = currentStadtteil
				if strings.Contains(scrape.Text(scrapingResult[i]), "Altkleider") {
					temp.Altkleider = true
				}
				if strings.Contains(scrape.Text(scrapingResult[i]), "Papier") {
					temp.Papier = true
				}
				if strings.Contains(scrape.Text(scrapingResult[i]), "Glas") {
					temp.Glas = true
				}
				tempSplit := strings.SplitN(scrape.Text(scrapingResult[i]), ",", 2)
				temp.Strasse = tempSplit[0]
				address := temp.Strasse + ",+Essen,+Germany"
				address = strings.Replace(address, " ", "+", -1)

				osmdata, err := geoosm.NewOSMData().GetJSON(address)
				if err != nil {
					log.Fatal(err)
				}

				if len(osmdata) != 0 {
					temp.Coordinates.Latitude = osmdata[0].Lat
					temp.Coordinates.Longitude = osmdata[0].Lon
				}

				standorte[strconv.Itoa(mapID)] = temp
				mapID++
			}
		}
	}

	json, err := json.Marshal(standorte)
	if err != nil {
		panic("Error: Could not convert map to JSON data")
	}
	ioutil.WriteFile("./build/data.json", json, 0644)

	timeNeeded := time.Now().Sub(timeStart)

	log.Print("Successfully written data.json in ")
	log.Println(timeNeeded)
}
