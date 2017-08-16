# Setup

You'll need go 1.8. To update your existing version, you can follow [these instructions](https://gist.github.com/nikhita/432436d570b89cab172dcf2894465753).

# Running tests

To run all tests in all packages, you can run `$ go test $(go list ./... | grep -v /vendor/)`.