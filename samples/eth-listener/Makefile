OUTPUT_DIR=bin
APP_BINARY=eth-listener

.PHONY: clean, build, test, run

clean:
	rm -fr ${OUTPUT_DIR}

build:
	go build -o ${OUTPUT_DIR}/${APP_BINARY} cli/main.go

run:
	go run cli/main.go