BINARY_NAME=compscore
OUT_FILE=build/$(BINARY_NAME)

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%Y-%m-%d\ %H:%M)

LDFLAGS= -ldflags "-X 'main.gitCommit=${GIT_COMMIT}' -X 'main.gitBranch=${GIT_BRANCH}' -X 'main.buildDate=${BUILD_DATE}'"

.PHONY: build install

build:
	go build $(LDFLAGS) -o $(OUT_FILE) main.go

install: build
	sudo cp $(OUT_FILE) $(GOPATH)/bin/$(BINARY_NAME)

cloc:
	cloc --exclude-dir=assets,node_modules --exclude-ext=json .

generate:
	go generate .