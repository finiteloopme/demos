
.PHONY: init, run, build

init:
	go mod tidy

run:
	go run ./cli/main.go

build:
	go build -o ./build/api-data-gen ./cli/main.go

test:
	go test -v --cover ./...