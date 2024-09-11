build:
	go build -o bin/web-radio

run: build
	./bin/web-radio

test: 
	go test -v ./...