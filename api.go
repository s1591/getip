package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const url string = "https://ipwho.is"

type ipHolder struct {
	IpVersion     string     `json:"type"`
	IpAddress     string     `json:"ip"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	Country       string     `json:"country"`
	CountryCode   string     `json:"country_code"`
	RegionCode    string     `json:"region_code"`
	Capital       string     `json:"capital"`
	ZipCode       string     `json:"postal"`
	CityName      string     `json:"city"`
	RegionName    string     `json:"region"`
	Continent     string     `json:"continent"`
	ContinentCode string     `json:"continent_code"`
	CallingCode   string     `json:"calling_code"`
	TimeZone      timeZone   `json:"timeZone"`
	Connection    connection `json:"connection"`
	Flag          flag       `json:"flag"`
}

type flag struct {
	Emoji string `json:"emoji"`
}

type connection struct {
	Isp    string `json:"isp"`
	Domain string `json:"domain"`
}

type timeZone struct {
	Id    string `json:"id"`
	Abbr  string `json:"abbr"`
	IsDST bool   `json:"is_dst"`
	UTC   string `json:"utc"`
}

func getIp() ipHolder {

	var ipDetails ipHolder
	http.DefaultClient.Timeout = 15 * time.Second

	resp, err := http.Get(url)
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

func dataString(data ipHolder) (ui string) {

	ui += fmt.Sprintf("\n%s Ip(%s): %s\n", data.Flag.Emoji, data.IpVersion, data.IpAddress)
	ui += fmt.Sprintf("Country: %s(%s), %s(%s)\n", data.Country, data.CountryCode, data.Continent, data.ContinentCode)
	ui += fmt.Sprintf("Region: %s(%s), %s\n", data.RegionName, data.RegionCode, data.CityName)
	ui += fmt.Sprintf("Latitude: %v, Longitude: %v\n", data.Latitude, data.Longitude)
	ui += fmt.Sprintf("Postal: %s, Calling Code: %s\n", data.ZipCode, data.CallingCode)
	ui += fmt.Sprintf("Capital: %s\n", data.Capital)
	ui += fmt.Sprintf("TimeZone: %s, %s, %s. DST: %v\n", data.TimeZone.Id, data.TimeZone.Abbr, data.TimeZone.UTC, data.TimeZone.IsDST)
	ui += fmt.Sprintf("Connection: %s, %s\n", data.Connection.Isp, data.Connection.Domain)

	return
}
