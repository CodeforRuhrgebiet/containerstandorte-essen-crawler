# Location of refuse containers in Essen

This script is a web scraper, getting the location of refuse containers in Essen, Germany, finding latitutes and longitudes using OpenStreetMap data and writes to JSON file.

## Data Source

If you want to learn more about the data we used, [here](http://abfallkalender.ebe-essen.de/containerstandorte-essen.php) you'll find the website we scraped.

## Installation

* [install docker](https://docs.docker.com/engine/installation/)

## Usage

run **`docker-compose up --build`**

## Alternative way
If you have installed Go on your machine, you are ready to go. Just download the required packages using **`go get`**

## LICENSE
MIT

## Data license
The OpenStreetMap data is provided under the terms of [Open Database License (ODbL)](http://opendatacommons.org/licenses/odbl).

## Credits
This project uses [geoosm](https://github.com/nicostuhlfauth/geoosm), licensed MIT and [scrape](https://github.com/yhat/scrape), licensed BSD-2. Thanks.
