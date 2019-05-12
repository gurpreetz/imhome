# I'm Home!

Turn on the lights when I come home. 

## How it works

* Get the time of the next [sunset](https://sunrise-sunset.org/api)
* Query [Mist](https://api-class.mist.com/) for presence of specific device (in this case, my phone)
* Create [applet](https://ifttt.com/create/if-send-ifttt-an-email-tagged?sid=12) on IFTTT which listens for tagged emails
* If sunset has elapsed, and phone has just connected to the netowrk
  * send an email to IFTTT with defined tag
  
## Requirements

* MIST Access Point
* Mist API Token
* IFTTT account
* Lat/Long for which sunset is to be calculated
* Smart Outlet/Device with IFTTT support
* App password in case using gmail account for emails

Place these in a config file : `config.json`
```
{
    "mistAPIToken" : "",
    "siteID" : "",
    "deviceMAC" : "",
    "coordinates" : {
        "latitude" : "37.733795",
        "longitude": "-122.446747"
    },
    "email" : {
        "from" : "YOUR IFTTT email address",
        "pass" : "",
        "smtp" : "smtp.gmail.com",
        "port" : "587",
        "subject" : "#home",
        "to" : "trigger@applet.ifttt.com"
    }    
}
```
