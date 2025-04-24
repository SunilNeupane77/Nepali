.PHONY: build install test clean

build:
	go build -o bin/nepali cmd/nepali/main.go

install:
	go install ./cmd/nepali

test:
	go test ./...

clean:
	rm -rf bin/
	go clean 