SHELL := /bin/bash

build:
	go vet ./src/
	go build -o ./bin/load_balancer ./src/
