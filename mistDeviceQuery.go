package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceInfo - expected return from API
type DeviceInfo struct {
	Manufacturer string `json:"manufacture"`
	HostName     string `json:"hostname"`
	UpTime       int    `json:"uptime"`
}

func doRequest(req *http.Request, apiToken string) ([]byte, error) {
	token := fmt.Sprintf("Token " + apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 404 == resp.StatusCode {
		return nil, errors.New("Device not found")
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// GetDeviceInfo - query the REST API
func GetDeviceInfo(cfg Config) (*DeviceInfo, error) {
	url := fmt.Sprintf("https://api.mistsys.com/api/v1/sites/%s/stats/clients/%s", cfg.SiteID, cfg.DeviceMac)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := doRequest(req, cfg.MistAPIToken)
	if err != nil {
		return nil, err
	}
	var data DeviceInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
