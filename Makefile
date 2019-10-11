exe = main.go
cmd = edm

all: deps

deps:
	@echo INFO: Installing dependencies
	@go mod vendor
build:
	@echo INFO: Building dist/$(cmd)
	@go build -o dist/$(cmd) $(exe)

run:
	@go run $(exe) ${ARGS}