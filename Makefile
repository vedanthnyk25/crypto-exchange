build:
	mkdir -p bin
	go build -o bin/exchange

run: build
	./bin/exchange

test:
	go test -v ./...
