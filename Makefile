BINARY_NAME=compscore
OUT_FILE=build/$(BINARY_NAME)

.PHONY: build clean start status stop kill

build:
	go build -o $(OUT_FILE) main.go

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

