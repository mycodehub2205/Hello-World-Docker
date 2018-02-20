# Golang Hello World Demo
!["AWS Build Status"](https://codebuild.eu-central-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiLy9HTkdEUFNBUWNFWnpsNloyR1IvSUhwd3J0bkJBOGtyQ245MDBtYjMvZWdjTDNISjZia01aVkJ0ZS9peU9jOG1Ca2p1ZUdHMGZTY1Q1VG1NQ25UTTZjPSIsIml2UGFyYW1ldGVyU3BlYyI6InQrV2NSazNtNUNEM1dSTFYiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)
## Test

```
$ go test
```

## Build

```
$ go build cmd/main.go
```

## Build docker image

```
$ go build cmd/main.go && docker build -t synoa/helloworld-go .
```
