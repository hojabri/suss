# SUSS
Suspicious user session system

This service, detects suspicious user logins, based on login location and, the time spent from the last login.
It assumes the maximum speed of possible movement of 500 Miles per hour (configurable in the config files) if the location of use is changed while login.
Location is extracted by looking up in a Geo DB (from Maxmind).

## Instructions

1. Make sure you have Go installed ([download](https://golang.org/dl/)).
2. To build an executable binary file, type `make build` in a command line terminal. This will build and create `susswebservice` file.
3. To build and run the docker package container for this application, type in the terminal: ``make docker-up``
4. To run unit tests, type in the terminal: ``make test``

## Configuration:
- There are two versions of configurations file by default: `dev` and `production`, but you can create as many configuration file with any name. To use a specific config file before running the application, specify `MODE` environment variable. for example `MODE=dev`, `MODE=production` or `MODE=your_config_file`. This will tell the application to look for `dev.yaml`, `production.yaml` or `your_config_file.yaml` file and reads configurations from that file.
- Configuration files are located under ``configs`` directory.
- Format of the configuration file is YAML type.

#### SUSPICIOUS_THRESHOLD:

The maximum speed of supposed a use can move in the earth between two logins. 

#### DETECTION_MODE:

It is used how to determine a suspicious login according to distance and speed.
Detection mode either Optimistic, Normal or Pessimistic.

- _Optimistic_:   will deduct accuracy radius from distance
- _Normal_:       ignore accuracy radius
- _Pessimistic_:  will add accuracy radius to distance

#### GEO_CITY_DB:
The file name of GEO CITY DB. (Please download the latest version from https://dev.maxmind.com/geoip/geoip2/geolite2/)
It located under the ``geodb`` folder.

#### WEBSERVICE:

This contains some key/value configuration of the webservice.
- PORT: port number of webservice
- DOMAIN: domain name. it is used for get a valid ssl certification from letencrypt website (https://letsencrypt.org/). To use this feature, we need to host the webservice in a registered domain.
- ENABLE_AUTOCERT: it indicates whether we want to enable auto certificate from letsencrypt site or not. Don't make it enable, unless you host the webservice in a registered domain.

#### DB:
- GORM_LOG_LEVEL: the log level of GORM (ORM of the db), if set to `info` it will print all SQL queries, while running. if set to `error`, it only prints GORM related errors in the log.

## How to run
1- Build and run locally:
- Build the application (read instructions part)
- In the terminal type:
```
MODE=dev ./susswebservice
```

Output:
```
...
INFO[2021-04-25T13:25:46.980009+03:00] PORT:5000                                    
INFO[2021-04-25T13:25:46.980032+03:00] ENABLE_AUTOCERT:false                        
INFO[2021-04-25T13:25:46.980037+03:00] DOMAIN:localhost                             
INFO[2021-04-25T13:25:46.980786+03:00] Serving 'SUSS - 1.0.0'                       
INFO[2021-04-25T13:25:46.980799+03:00] Insecure HTTP                                

 ┌───────────────────────────────────────────────────┐ 
 │                    Fiber v2.8.0                   │ 
 │               http://127.0.0.1:5000               │ 
 │       (bound on host 0.0.0.0 and port 5000)       │ 
 │                                                   │ 
 │ Handlers ............. 9  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID .............. 2516 │ 
 └───────────────────────────────────────────────────┘ 
```

2- Run in a docker container:

- Build and run docker container (`make docker-up`)
- Docker will run in the background and, you can then use the webservice.

### Consuming the webservice

- To check if the webservice is running, call this utl in a browser (http://127.0.0.1:5000).
_Note: Make sure the port 5000 is already free in the hosting environment._
  ![SUSS home page](https://raw.githubusercontent.com/hojabri/suss/main/static/suus_first_page.png)  
  
This is the swagger file documentation of the webservice.
In this page you can also download the swagger file.

- To test the webservice in Postman tools, there already a ready to use Postman collection file, named `postman_collection.json` under the `openapi` folder. Simply import it to postman, and call the webservice endpoints.

- At the moment, the only endpoint implemented is:
```POST http://localhost:5000/v1/event```
  
- Request body format:
```json
{
  "username": "Omid",
  "unix_timestamp": 1619268172,
  "event_uuid": "fa05a9bf-bf5c-4db2-88f7-dd764d2cf1f7",
  "ip_address": "125.219.64.174"
}
```

- Response body format:
```json
{
    "currentGeo": {
        "lat": 35.1865,
        "lon": 138.6628,
        "radius": 200
    },
    "travelToCurrentGeoSuspicious": false,
    "travelFromCurrentGeoSuspicious": false,
    "precedingIpAccess": {
        "lat": 34.7732,
        "lon": 113.722,
        "radius": 1000,
        "speed": 38.07041358517798,
        "ip": "125.219.64.174",
        "timestamp": 1619281807
    },
    "subsequentIpAccess": {
        "lat": 0,
        "lon": 0,
        "radius": 0,
        "speed": 0,
        "ip": "",
        "timestamp": 0
    }
}
```

- After calling the endpoint, with POST method, it checks the ``event_uuid`` value, and the application assumes this value as it own primary key in the events table. So for creating a new login event, you should pass a unique uuid, otherwise it will use the previously saved record for that event.
- To check a previously saved event, among the preceding or subsequent event, you can use the uuid for that event which is already saved.
- The maximum allowed speed which the application judges the suspicious login can be configured in the config files, but the default value is 500 Miles per hour.

## Authentication

This webservice can check for authentication, based on API_KEY and API_PASSWORD fields in the request POST call. The
To enable authentication, you can fill SERVER_API_KEY and SERVER_API_PASSWORD environment variables, before running the application.
For example:

```
 SERVER_API_KEY="my_api_key" SERVER_API_PASSWORD="my_api_password"  MODE=dev ./susswebservice
```

Then add two headers `API_KEY` and `API_PASSWORD` in the POST method:
![SUSS home page](https://raw.githubusercontent.com/hojabri/suss/main/static/suss_authentication.png)



This product includes GeoLite2 data created by MaxMind, available from
https://www.maxmind.com