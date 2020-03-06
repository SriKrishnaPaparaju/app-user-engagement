install:
	go install ./...

clean:
	go clean ./...

test:
	go test ./... -cover

build:
	go mod tidy
	go mod vendor
	go build ./...
