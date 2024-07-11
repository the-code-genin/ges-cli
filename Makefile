.PHONY: test
test:
	go test -v ./...

.PHONY: ges
ges:
	go build -o build/bin/ges ./cmd/ges

.PHONY: build
build: ges
