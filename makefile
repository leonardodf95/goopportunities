.PHONY: default run build tests docs clean

APP_NAME=goopportunities

default: run-docs

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

tests:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

run-docs:
	@swag init
	@go run main.go
