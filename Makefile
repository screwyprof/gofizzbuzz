# This repo's root import path.
PKG := github.com/screwyprof/gofizzbuzz

## DO NOT EDIT BELLOW THIS LINE
GO_FILES = $(shell find . -name "*.go" | grep -v vendor | uniq)

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
MAKE_COLOR=\033[36m%-20s\033[0m

all:  lint test ## build application, run linters and tests

fmt: ## format code
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	@gofumpt -s -l -w $(GO_FILES)
	@gci -w -local $(PKG) $(GO_FILES)

lint: ## run linters for current changes
	@echo "$(OK_COLOR)==> Linting current changes$(NO_COLOR)"
	golangci-lint run ./...

lint-all: ## run linters
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
	golangci-lint run ./... --new-from-rev=""

test: ## run tests
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	go test -v -race -bench -benchmem -timeout=10s -count=1 ./...

test-cover-txt: ## run tests with code coverage and show plain report in console
	@echo "$(OK_COLOR)==> Running tests with coverage $(NO_COLOR)"
	go test -v -race -bench -benchmem -timeout=10s -count=1 -cover -coverpkg=./... -coverprofile=coverage.tmp ./...
	@cat coverage.tmp | grep -v -e *_gen.go -e backend-server.go -e  backend-types.go -e internal -e securosyswrapper -e pkg/web3signer/client -e pkg/eth2/spec -e deposittest -e handlertest > coverage.txt
	@go tool cover -func coverage.txt

test-cover-html: ## run tests with code coverage and show html report
	@echo "$(OK_COLOR)==> Running tests with coverage $(NO_COLOR)"
	@go test -v -race -bench -benchmem -timeout=10s -count=1 -cover -coverpkg=./... -coverprofile=coverage.tmp ./...
	@cat coverage.tmp | grep -v -e *_gen.go -e backend-server.go -e  backend-types.go -e internal -e securosyswrapper -e pkg/web3signer/client -e pkg/eth2/spec -e deposittest -e handlertest > coverage.txt
	@go tool cover -html=coverage.txt

clean: ## cleans-up artifacts
	@echo "$(OK_COLOR)==> Cleaning up$(NO_COLOR)"
	@rm -rf ./coverage.txt
	@rm -rf ./coverage.tmp
	@rm -rf ./build/yucatan

help: ## show this help
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) %s\n", $$1, $$2}'

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: all fmt test test-cover-txt test-cover-html lint lint-all help
