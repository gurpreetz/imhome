package main

import (
	"fmt"
	"time"
)

// FreshConnectUptime - time less than which a device is considered to be a new connection
const FreshConnectUptime = 180 // seconds

// DeviceReconnectActionTime - device must be absent for at least this long before any
// action will be taken
const DeviceReconnectActionTime = 30 * time.Minute

func main() {
	cfg, err := cfgParser()
	if err != nil {
		return
	}

	var deviceLastSeenTime time.Time

	solarData, err := GetSolarData(cfg.Coordinates)
	fmt.Printf("Sunset will be in %v\n", time.Until(solarData.Sunset))
	fmt.Printf("Sunrise will be in %v\n", time.Until(solarData.Sunrise))

	devTicker := time.NewTicker(1 * time.Minute)
	defer devTicker.Stop()

	sunTicker := time.NewTicker(12 * time.Hour)
	defer sunTicker.Stop()
	for {
		select {
		case <-devTicker.C:
			deviceInfo, err := GetDeviceInfo(cfg)
			if err == nil {
				if time.Since(deviceLastSeenTime) > DeviceReconnectActionTime {
					fmt.Printf("%v :: Device last seen at %v \n", time.Now(), deviceLastSeenTime)
					if deviceInfo.UpTime < FreshConnectUptime {
						fmt.Printf("%v :: Device just connected after being away\n", time.Now())
						currentTime := time.Now()
						if currentTime.After(solarData.Sunset.Local()) {
							fmt.Printf("%v :: Do something\n", time.Now())
							sendMail(cfg.Email)
						} else {
							fmt.Printf("%v :: Sun is still up; Do nothing\n", time.Now())
						}
					}
				}
				deviceLastSeenTime = time.Now()
				fmt.Printf("Found Device %s connected for %d seconds\n", deviceInfo.HostName, deviceInfo.UpTime)
			}
		case <-sunTicker.C:
			solarData, err = GetSolarData(cfg.Coordinates)
			fmt.Printf("Sunset will be in %v\n", time.Until(solarData.Sunset))
			fmt.Printf("Sunrise will be in %v\n", time.Until(solarData.Sunrise))
		}
	}
}
