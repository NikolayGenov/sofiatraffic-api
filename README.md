# sofiatraffic-api

[![Build Status](https://travis-ci.org/NikolayGenov/sofiatraffic-api.svg?branch=master)](https://travis-ci.org/NikolayGenov/sofiatraffic-api)
[![Coverage Status](https://coveralls.io/repos/github/NikolayGenov/sofiatraffic-api/badge.svg?branch=master)](https://coveralls.io/github/NikolayGenov/sofiatraffic-api?branch=master)

## Usage
You can generate/update the auto generated code by executing:
```
swagger generate server -f ./swagger.yml -A SofiaTraffic
```
in the project main directory

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

## Documentation
Currently the documentation of the API can be found on `https://localhost:8080/v1/docs`
only when the server is running
