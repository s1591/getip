package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var URL string = "https://freeipapi.com/api/json"

type ipHolder struct {
	IpVersion     int     `json:"ipVersion"`
	IpAddress     string  `json:"ipAddress"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Country       string  `json:"countryName"`
	CountryCode   string  `json:"countryCode"`
	TimeZone      string  `json:"timeZone"`
	ZipCode       string  `json:"zipCode"`
	CityName      string  `json:"cityName"`
	RegionName    string  `json:"regionName"`
	IsProxy       bool    `json:"isProxy"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
}

func getIp() ipHolder {

	var ipDetails ipHolder
	http.DefaultClient.Timeout = 15 * time.Second

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &ipDetails)
	if err != nil {
		panic(err)
	}

	return ipDetails

}
