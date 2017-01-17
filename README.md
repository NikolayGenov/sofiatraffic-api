# Sofia Traffic API

[![Build Status](https://travis-ci.org/NikolayGenov/sofiatraffic-api.svg?branch=master)](https://travis-ci.org/NikolayGenov/sofiatraffic-api)
[![codecov](https://codecov.io/gh/NikolayGenov/sofiatraffic-api/branch/master/graph/badge.svg)](https://codecov.io/gh/NikolayGenov/sofiatraffic-api)

## Goal
The goal is to be able to provide more useful and usable interface (compared with the one as of end of 2016) to the service provided by Sofia Urban Mobility Center for the Sofia's Public Transportation
The API will be a modern API using the Open API standard. It will use swagger spec, which can be used to auto generate clients that will consume the API

## Planned functionality
The API will have main functionality to get, stations, lines and ETA (Estimated Time of Arrival) of the desired vehicle.


* The user of the API should be able to search for stations using names, Urban Mobility Center codes, and other.

* The user should be able to get all relevant information about given line - start station , final station, list of all stations, and soft schedule for that line at given days.

* The user should be able to get ETA of all vehicles serving a given station.

* An administrator should be able to change the information about the soft schedules, stations, lines, for example adding a new line or deleting a station.

The API is planned to work with `json` as format and to try to follow RESTful principles.

## Usage
(TODO: Can that be a library in some way? )

After that you should `go build` in the `cmd` directory, where `main.go` is located

In order to start the server to serve the API you need a certificate.
A quick and dirty way is by using this one liner:
```
openssl req -newkey rsa:2048 -nodes -keyout domain.key -x509 -days 365 -out domain.crt
```

Only then you can start the server on port `8080`:
```
./sofia-traffic-server --tls-key domain.key --tls-certificate domain.crt --tls-port 8080
```
## When making changes
You can update the auto generated code by executing the following code in the project main directory:
```
swagger generate server -f ./swagger.yml -A SofiaTraffic
```

## Documentation
Currently the documentation of the API can be found on `https://localhost:8080/v1/docs`
only when the server is running
