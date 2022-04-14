VERSION = $(shell git describe --tags)
BUILD_TIME = $(shell date +%FT%T)
GIT_COMMIT = $(shell expr substr `git log --pretty=format:'%H' -n 1` 1 7)
PACKAGE = github.com/xuender/gopy
CONST_PACKAGE = ${PACKAGE}/app

default: lint clean build go

test:
	@go run cmd/server/main.go -alsologtostderr

lint:
	golangci-lint run

dist/maps: proto wire
	go build -ldflags \
  "-X '${CONST_PACKAGE}.Version=${VERSION}' \
  -X '${CONST_PACKAGE}.BuildTime=${BUILD_TIME}' \
	-X '${CONST_PACKAGE}.GitCommit=${GIT_COMMIT}'" \
  -o dist/maps cmd/server/main.go
