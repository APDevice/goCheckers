BINNAME=goCheckers
FILES=main.go view.go controller.go

all: run

build:
	go build -o ./bin/$(BINNAME) $(FILES)

run:
	go run $(FILES)