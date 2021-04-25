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
  ![SUSS home page](https://github.com/hojabri/suss/static/main/suss_first_page.png?raw=true)  
This product includes GeoLite2 data created by MaxMind, available from
https://www.maxmind.com