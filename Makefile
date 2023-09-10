GO_FILES = $(shell find . '(' -path '*/.*' -o -path './vendor' ')' -prune -o -name '*.go' -print | cut -b3-)
GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./...)

dod: build test fmt lint

build:
	go build -v ./...

test:
	go test -v ./...

fmt:
	gofmt -s -w ${GO_FILES}
	gofumpt -l -w ${GO_FILES}
	goimports -w ${GO_PATHS}

lint:
	goreportcard-cli -v
	golangci-lint run --config=.golangci.yml ./...

install:
	bash install.sh