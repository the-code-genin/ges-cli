test: test
	go test -v ./...

ges:
	@go build -o build/bin/ges ./cmd/ges/

all: ges