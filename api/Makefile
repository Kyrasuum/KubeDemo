.PHONY: build docker
build:
	@go mod tidy
	@go build -o api cmd/main.go
	@cp -r internal/data ./data

docker:
	@docker build -t kubedemo-api -f docker/Dockerfile .
