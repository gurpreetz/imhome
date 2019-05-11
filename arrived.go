package main

import (
	"fmt"
	"time"
)

func main() {
	cfg, err := cfgParser()
	if err != nil {
		return
	}

	solarData, err := GetSolarData(cfg.Coordinates)
	nextSunsetIn := time.Until(solarData.Sunset)
	fmt.Printf("Sunset will be in %v\n", nextSunsetIn)

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	var devicePresent bool
	devicePresent = false
	deviceInfo, err := GetDeviceInfo(cfg)
	if err == nil {
		devicePresent = true
	}
	for {
		select {
		case <-ticker.C:
			deviceInfo, err = GetDeviceInfo(cfg)
			if err == nil {
				devicePresent = true
				fmt.Printf("Found Device %s connected for %d seconds\n", deviceInfo.HostName, deviceInfo.UpTime)
			}
			fmt.Println(devicePresent)
		}
	}
}
