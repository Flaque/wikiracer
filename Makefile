govendor:
	go get -u github.com/kardianos/govendor

run: 
	govendor sync
	go install
	@wikiracer

docker:
	govendor sync
	docker build -t wikiracer .
	docker run --publish 6060:8080 --name test --rm wikiracer
