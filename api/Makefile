.PHONY: build docker
build:
	@go mod tidy
	@go build -o api cmd/main.go

docker:
	@docker build -t kubedemo-api -f docker/Dockerfile .
