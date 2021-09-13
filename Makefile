#!bin/sh
define USAGE
Commands:
    run			Run main.go
	build		Build binary
    start       Run built app
endef

run:
	go mod vendor
	docker-compose up

build:
	go mod vendor
	go build -o dist/user-service

start:
	./dist/user-service

clean:
	docker-compose down