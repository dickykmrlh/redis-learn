pkgs = $(shell go list ./...)

.PHONY: build

# go build command
build:
	@go build -v -o redislearn cmd/*.go

# go run command
run:
	make build
	@./redislearn

test:
	@echo "RUN TESTING..."
	@go test -v -cover -race $(pkgs)