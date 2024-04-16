.PHONY: build go web docker
build: web go

go:
	@go mod tidy
	@go build -o web cmd/main.go

web:
	@$(MAKE) --no-print-directory -C internal/pages

docker:
	@docker build -t kubedemo-web -f docker/Dockerfile .
