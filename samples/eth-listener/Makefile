OUTPUT_DIR=bin
APP_BINARY=eth-listener

# For _WS_URL & _API_KEY
include config.properties

.PHONY: clean, build, test, run

clean:
	rm -fr ${OUTPUT_DIR}

build:
	go build -o ${OUTPUT_DIR}/${APP_BINARY} cli/main.go

run:
	export WS_URL=${_WS_URL}; export API_KEY=${_API_KEY}; go run cli/main.go