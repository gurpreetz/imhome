package main

import (
	"fmt"
	"time"
)

// freshConnectUptime - time less than which a device is considered to be a new connection
const freshConnectUptime = 180 // seconds

// deviceReconnectActionTime - device must be absent for at least this long before any
// action will be taken
const deviceReconnectActionTime = 30 * time.Minute

func getDeviceStatus(deviceInfo *DeviceInfo, deviceLastSeenTime time.Time, sunsetTime time.Time) (bool, time.Time) {
	if time.Since(deviceLastSeenTime) > deviceReconnectActionTime {
		fmt.Printf("%v :: Device last seen at %v \n", time.Now(), deviceLastSeenTime)
		if deviceInfo.UpTime < freshConnectUptime {
			fmt.Printf("%v :: Device just connected after being away\n", time.Now())
			currentTime := time.Now()
			if currentTime.After(sunsetTime.Local()) {
				fmt.Printf("%v :: Do something\n", time.Now())
				return false, deviceLastSeenTime
			}
			fmt.Printf("%v :: Sun is still up; Do nothing\n", time.Now())
		}
	}
	deviceLastSeenTime = time.Now()
	//fmt.Printf("Found Device %s connected for %d seconds\n", deviceInfo.HostName, deviceInfo.UpTime)
	return false, deviceLastSeenTime
}

func main() {
	cfg, err := cfgParser()
	if err != nil {
		return
	}

	var deviceLastSeenTime time.Time
	var deviceConnected bool

	devTicker := time.NewTicker(30 * time.Second)
	defer devTicker.Stop()

	sunTicker := time.NewTicker(6 * time.Hour)
	defer sunTicker.Stop()

	solarData, err := GetSolarData(cfg.Coordinates)

	for {
		select {
		case <-devTicker.C:
			deviceInfo, err := GetDeviceInfo(cfg)
			if err == nil {
				deviceConnected, deviceLastSeenTime = getDeviceStatus(deviceInfo, deviceLastSeenTime, solarData.Sunset)
				if deviceConnected == true {
					sendMail(cfg.Email)
				}
			}
		case <-sunTicker.C:
			solarData, err = GetSolarData(cfg.Coordinates)
		}
	}
}
