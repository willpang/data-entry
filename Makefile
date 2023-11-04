run:
	go run main.go

build:
	go build -o bin/main main.go

test:
	go test ./...

clean:
	rm -f bin/main
