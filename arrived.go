package main

import (
	"fmt"
	"time"
)

// FreshConnectUptime - time less than which a device is considered to be a new connection
const FreshConnectUptime = 180 // seconds

func main() {
	cfg, err := cfgParser()
	if err != nil {
		return
	}

	solarData, err := GetSolarData(cfg.Coordinates)
	fmt.Printf("Sunset will be in %v\n", time.Until(solarData.Sunset))

	devTicker := time.NewTicker(1 * time.Minute)
	defer devTicker.Stop()

	sunTicker := time.NewTicker(12 * time.Hour)
	defer sunTicker.Stop()
	for {
		select {
		case <-devTicker.C:
			deviceInfo, err := GetDeviceInfo(cfg)
			if err == nil {
				if deviceInfo.UpTime < FreshConnectUptime {
					fmt.Printf("Device just connected\n")
					currentTime := time.Now()
					if currentTime.After(solarData.Sunset.Local()) {
						fmt.Printf("Do something\n")
					} else {
						fmt.Printf("Sun is still up; Do nothing\n")
					}
				}
				fmt.Printf("Found Device %s connected for %d seconds\n", deviceInfo.HostName, deviceInfo.UpTime)
			}
		case <-sunTicker.C:
			solarData, err = GetSolarData(cfg.Coordinates)
		}
	}
}
