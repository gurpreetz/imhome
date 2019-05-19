package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SunriseSunsetResults represents the results struct from sunrise-sunset.org
type SunriseSunsetResults struct {
	Sunrise                   time.Time `json:"sunrise"`
	Sunset                    time.Time `json:"sunset"`
	SolarNoon                 time.Time `json:"solar_noon"`
	DayLength                 int       `json:"day_length"`
	CivilTwilightBegin        time.Time `json:"civil_twilight_begin"`
	CivilTwilightEnd          time.Time `json:"civil_twilight_end"`
	NauticalTwilightBegin     time.Time `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       time.Time `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin time.Time `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   time.Time `json:"astronomical_twilight_end"`
}

// SunriseSunsetResponseContainer represents the response container from sunrise-sunset.org
type SunriseSunsetResponseContainer struct {
	Results SunriseSunsetResults `json:"results"`
	Status  string               `json:"status"`
}

// GetSolarData - get the sunset time
func GetSolarData(coordinates Coordinates) (*SunriseSunsetResults, error) {
	date := time.Now()

	solarDataURL := fmt.Sprintf("https://api.sunrise-sunset.org/json?formatted=0&lat=%f&lng=%f&date=%s",
		coordinates.Latitude, coordinates.Longitude, date.Format("2006-01-02"))

	sunriseSunsetRsp, err := http.Get(solarDataURL)
	if err != nil {
		return nil, err
	}
	if sunriseSunsetRsp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected Sunrise-Sunset Response Code: %d", sunriseSunsetRsp.StatusCode)
	}

	body, err := ioutil.ReadAll(sunriseSunsetRsp.Body)
	if err != nil {
		return nil, err
	}
	container := &SunriseSunsetResponseContainer{}
	err = json.Unmarshal(body, container)
	if err != nil {
		return nil, err
	}
	if container.Status != "OK" {
		return nil, fmt.Errorf("Invalid Sunrise-Sunset Response Status: %s", container.Status)
	}
	fmt.Printf("*** Queried at %v *** \n", time.Now())
	fmt.Printf("Sunset will be in %v\n", time.Until(container.Results.Sunset))
	fmt.Printf("Sunrise will be in %v\n", time.Until(container.Results.Sunrise))

	return &container.Results, nil
}
