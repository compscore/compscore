BINARY_NAME=compscore
OUT_FILE=build/$(BINARY_NAME)

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%Y-%m-%d\ %H:%M)

LDFLAGS=-ldflags "-X 'main.gitCommit=${GIT_COMMIT}' -X 'main.gitBranch=${GIT_BRANCH}' -X 'main.buildDate=${BUILD_DATE}'"

.PHONY: build clean start status stop kill

build:
	go build $(LDFLAGS) -o $(OUT_FILE) main.go

clean:
	rm -f $(OUT_FILE)

start: 
	./$(OUT_FILE) engine start

status:  
	./$(OUT_FILE) engine status

stop:
	./$(OUT_FILE) engine stop

kill:
	./$(OUT_FILE) engine kill

