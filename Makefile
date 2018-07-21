GOPATH=$(shell pwd)

all: test build

build:
	go build -o go-friend main.go

test:
	go test -cover ./...