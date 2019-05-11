package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config - config to drive this program
type Config struct {
	MistAPIToken string      `json:"mistAPIToken"`
	SiteID       string      `json:"siteID"`
	DeviceMac    string      `json:"deviceMAC"`
	Coordinates  Coordinates `json:"coordinates"`
}

// Coordinates represent geographic coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

func cfgParser() (Config, error) {
	var cfg Config
	cfgFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer cfgFile.Close()

	bytes, _ := ioutil.ReadAll(cfgFile)

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Printf("Have you constructed a valid `config.json` ? : %s\n", err)
		return cfg, err
	}
	return cfg, nil

}
