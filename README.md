# I'm Home!

Turn on the lights when I come home. 

## How it works

* Get the time of the next sunset
* Query MIST for presence of specific device (in this case, my phone)
* If sunset has elapse, and phone has just connected to the netowrk
  * send IFTTT to light
  
## Requirements

* MIST API Token
* Lat/Long for which sunset is to be calculated

Place these in a config file : `config.json`
```
{
    "mistAPIToken" : "",
    "siteID" : "",
    "deviceMAC" : "",
    "coordinates" : {
        "latitude" : "37.733795",
        "longitude": "-122.446747"
    }
}
```