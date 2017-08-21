# 🚗 Wikiracer

## Running 

### Setup

You'll need go 1.8. To update your existing version, you can follow [these instructions](https://gist.github.com/nikhita/432436d570b89cab172dcf2894465753).

### Running a build

Building and running should just be 

```sh
$ git clone https://github.com/Flaque/wikiracer.git
$ cd wikiracer
$ make run
```

Then you can hit endpoints from `localhost:8080`.

### Running tests

To run all tests in all packages, you can run:

```
$ go test $(go list ./... | grep -v vendor)`
``` 

To exclude the integration tests (and load tests) (that make service calls and are therefore slower), run:
```
$ go test $(go list ./... | grep -v integration_tests | grep -v load_tests | grep -v vendor)
```

### Dockerizing

If you have docker installed, you should be able to dockerize the project with:

```
make docker
```

Then you can hit endpoints from `localhost:6060`. 

### Testing out an endpoint

To try out a search, try the connections between `Dog` and `Airplane`:

If running locally without docker: [http://localhost:8080/search/Cat/Airplane](http://localhost:8080/search/Cat/Airplane)

If running locally with docker: [http://localhost:6060/search/Dog/Airplane](http://localhost:6060/search/Cat/Airplane)

# Technical Overview

This wikiracer uses a concurrent **breadth first search (BFS)** in order to find a path as fast as possible. It keeps track of nodes in a priority queue where the priority is defined as the current depth of the node. That way, even if nodes get added out of order (because one goroutine finishes quicker than another), they'll still be searched in roughly equivilant "rows" of the search tree. If we didn't do this, then the search speed would be wildly different each time since some might end up looking more like a depth first search (which doesn't work all that well). 

When all goes well, the search can take roughly 1 to 3 seconds. We also **cache previous wikipedia nodes, as well as previous requests**, so if we search through nodes we've already seen or search the same thing twice, it should be under a second. 

## Logging 
The project uses structured logging in JSON format via [zap](https://github.com/uber-go/zap). This let's us easily query our logs and easily ship them to other services in the future. 

## Packages
There are currently 5 packages. 

### wikimedia
The wikimedia package let's us interact with the wikmedia query API. Note that this is different than the REST API which will return HTML data.

### search
The search package is what runs our concurrent BFS through wikimedia. 

### load_tests
Load tests is a WIP package that may eventually let us run some load tests. At the moment it's not perfect since it relies on there already being an instance of our API running. 

### integration_tests
Integration tests are currently in their own seperate package to make it easier to skip over them if you'd just like to run simple unit tests. These tests will query the api and therefore take a bit to complete.

### tracer
This package is a simple little utility to log the time a function takes.
