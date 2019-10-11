exe = main.go
cmd = edm
TRAVIS_TAG ?= "0.0.0"

all: deps

deps:
	@echo INFO: Installing dependencies
	@go mod vendor

run:
	@go run $(exe) ${ARGS}

install:
	@echo INFO: Installing edm
	@go install

test:
	@echo INFO: Running tests
	@go test ./...

test-coverage:
	@echo INFO: Generate coverage
	@go test -coverprofile c.out ./...

release-amd64:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-linux-amd64 $(exe)
	@echo INFO: Finished Building $@

release-arm:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-linux-arm $(exe)
	@echo INFO: Finished Building $@

release-386:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-linux-386 $(exe)
	@echo INFO: Finished Building $@

release-mac-amd64:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-mac-amd64 $(exe)
	@echo INFO: Finished Building $@

release-windows-amd64:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-windows-amd64.exe $(exe)
	@echo INFO: Finished Building $@

release-windows-386:
	@echo INFO: Building $@
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o dist/$(cmd)-windows-386.exe $(exe)
	@echo INFO: Finished Building $@
    	
release: release-amd64 release-arm release-386 release-mac-amd64 release-windows-386 release-windows-amd64