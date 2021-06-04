

all: gen fmt test install

deps:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint

clean:
	go clean ./...

install:
	go install ./...

test:
	go test ./...

gen:
	go generate ./...

fmt:
	find . -name '*.go' -exec goimports -l -w {} \;

lint:
	golint ./...

.PHONY: all clean install test gen fmt lint

