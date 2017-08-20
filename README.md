# Setup

You'll need go 1.8. To update your existing version, you can follow [these instructions](https://gist.github.com/nikhita/432436d570b89cab172dcf2894465753).

# Running tests

To run all tests in all packages, you can run `$ go test $(go list ./...)`. To exclude the integration tests (that make service calls and are therefore slower), run `$ go test $(go list ./... | grep -v integration_tests)`

# Dockerizing

Build:
```
docker build -t wikiracer .
```

Run: 
``` sh
$ docker run --publish 6060:8080 --name test --rm wikiracer
```

