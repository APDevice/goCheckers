BINNAME=goCheckers

all: run

build:
	go build -o ./bin/$(BINNAME) ./...

run:
	go run ./...