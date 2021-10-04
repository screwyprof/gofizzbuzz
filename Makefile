# This repo's root import path.
PKG := github.com/screwyprof/gofizzbuzz

GO = gotip

## DO NOT EDIT BELLOW THIS LINE
GO_FILES = $(shell find . -name "*.go" | grep -v vendor | uniq)

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
MAKE_COLOR=\033[36m%-20s\033[0m

all: test ## run tests

fmt: ## format code
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	$(GO) fmt ./...
	#@gofumpt -s -l -w $(GO_FILES)
	#@gci -w -local $(PKG) $(GO_FILES)

#lint: install-tools ## run linters for current changes
#	@echo "$(OK_COLOR)==> Linting current changes$(NO_COLOR)"
#	@tools/bin/golangci-lint run ./...

#lint-all: install-tools ## run linters
#	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
#	@tools/bin/golangci-lint run ./... --new-from-rev=""

test: ## run tests
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	$(GO) test -v -race -timeout=10s -count=1 ./...

test-bench: install-tools ## run benchmarks
	@echo "$(OK_COLOR)==> Running benchmarks$(NO_COLOR)"
	@$(GO) test -run=^$  -bench=. -benchmem | tools/bin/prettybench | grep -v -e github.com

test-cover-txt: ## run tests with code coverage and show plain report in console
	@echo "$(OK_COLOR)==> Running tests with coverage $(NO_COLOR)"
	@$(GO) test -v -race -timeout=10s -count=1 -cover -coverpkg=./... -coverprofile=coverage.tmp ./...
	@cat coverage.tmp | grep -v fbtest > coverage.txt
	@go tool cover -func coverage.txt

test-cover-html: ## run tests with code coverage and show html report
	@echo "$(OK_COLOR)==> Running tests with coverage $(NO_COLOR)"
	@$(GO) test -v -timeout=10s -count=1 -cover -coverpkg=./... -coverprofile=coverage.tmp ./...
	@cat coverage.tmp | grep -v fbtest > coverage.txt
	@go tool cover -html=coverage.txt

install-tools: ## install-tools
ifeq (,$(wildcard ./tools/bin/prettybench))
	@echo "$(OK_COLOR)--> Installing tools from tools/tools.go$(NO_COLOR)"
	@export GOBIN=$$PWD/tools/bin; export PATH=$$GOBIN:$$PATH; cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % $(GO) install %
endif

clean: ## cleans-up artifacts
	@echo "$(OK_COLOR)==> Cleaning up$(NO_COLOR)"
	rm -rf ./coverage.txt
	rm -rf ./coverage.tmp

help: ## show this help
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) %s\n", $$1, $$2}'

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: all fmt lint lint-all test test-bench test-cover-txt test-cover-html install-tools clean help
