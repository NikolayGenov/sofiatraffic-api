# Sofia Traffic API

[![Build Status](https://travis-ci.org/NikolayGenov/sofiatraffic-api.svg?branch=master)](https://travis-ci.org/NikolayGenov/sofiatraffic-api)
[![codecov](https://codecov.io/gh/NikolayGenov/sofiatraffic-api/branch/master/graph/badge.svg)](https://codecov.io/gh/NikolayGenov/sofiatraffic-api)



## Usage

After that you should `go build` in the `cmd` directory, where `main.go` is located

In order to start the server to serve the API you need a certificate.
A quick and dirty way is by using this one liner:
```
openssl req -newkey rsa:2048 -nodes -keyout domain.key -x509 -days 365 -out domain.crt
```

Only then you can start the server on port 8080:
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
