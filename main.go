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

	"github.com/kellydunn/golang-geo"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Standort Stores information about location of refuse container
type Standort struct {
	Coordinates *geo.Point `json:"coordinates"`
	Stadtteil   string     `json:"stadtteil"`
	Strasse     string     `json:"strasse"`
	Papier      bool       `json:"papier"`
	Altkleider  bool       `json:"altkleider"`
	Glas        bool       `json:"glas"`
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

	//geo.SetGoogleAPIKey("insert your key here")

	geocoder := new(geo.GoogleGeocoder)

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
				address := temp.Strasse + ", Essen, Germany"
				temp.Coordinates, err = geocoder.Geocode(address)
				if err != nil {
					log.Println(err)
				}
				log.Println(temp.Coordinates)
				standorte[strconv.Itoa(mapID)] = temp
				log.Println("Standort hinzugefügt")
				mapID++
			}
		}
	}

	json, err := json.Marshal(standorte)
	if err != nil {
		panic("Error: Could not convert map to JSON data")
	}
	ioutil.WriteFile("data.json", json, 0644)

	timeNeeded := time.Now().Sub(timeStart)

	log.Print("Successfully written data.json in ")
	log.Println(timeNeeded)
}
