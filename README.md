# Location of refuse containers in Essen

This script is a web scraper, getting the location of refuse containers in Essen, Germany, finding latitutes and longitudes using Google Geolocation API and writes to JSON file.

## Installation

* [install docker](https://docs.docker.com/engine/installation/)

## Usage

run **`docker-compose up --build`**

## Alternative way
If you have installed Go on your machine, simply download these two packages:
1. go get github.com/kellydunn/golang-geo
2. go get github.com/yhat/scrape
You may add your Google Geolocation API key in main.go. Than you are ready to run the script.

## Credits
This projects uses [golang-geo](https://github.com/kellydunn/golang-geo), licensed MIT and [scrape](https://github.com/yhat/scrape), licensed BSD-2. Thanks.
