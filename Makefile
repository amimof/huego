MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GOVERSION=$(shell go version | awk -F\go '{print $$3}' | awk '{print $$1}')
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))
BUILDPATH ?= $(BIN)/$(shell basename $(MODULE))
SRC_FILES=find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -not -path "./.cache/*" -print0 | xargs -0 
TBIN		 = $(CURDIR)/test/bin
GO			 = go
TIMEOUT  = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m➜\033[0m")

export GO111MODULE=on
export CGO_ENABLED=0

# Tools

$(TBIN):
	@mkdir -p $@
$(TBIN)/%: | $(TBIN) ; $(info $(M) building $(PACKAGE))
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(TBIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = $(TBIN)/golint
$(BIN)/golint: PACKAGE=golang.org/x/lint/golint

GOCYCLO = $(TBIN)/gocyclo
$(TBIN)/gocyclo: PACKAGE=github.com/fzipp/gocyclo/cmd/gocyclo

INEFFASSIGN = $(TBIN)/ineffassign
$(TBIN)/ineffassign: PACKAGE=github.com/gordonklaus/ineffassign

MISSPELL = $(TBIN)/misspell
$(TBIN)/misspell: PACKAGE=github.com/client9/misspell/cmd/misspell

GOLINT = $(TBIN)/golint
$(TBIN)/golint: PACKAGE=golang.org/x/lint/golint

GOCOV = $(TBIN)/gocov
$(TBIN)/gocov: PACKAGE=github.com/axw/gocov/...

# Tests

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint) @ ## Runs the golint command
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: gocyclo
gocyclo: | $(GOCYCLO) ; $(info $(M) running gocyclo) @ ## Calculates cyclomatic complexities of functions in Go source code
	$Q $(GOCYCLO) -over 25 .

.PHONY: ineffassign
ineffassign: | $(INEFFASSIGN) ; $(info $(M) running ineffassign) @ ## Detects ineffectual assignments in Go code
	$Q $(INEFFASSIGN) .

.PHONY: misspell
misspell: | $(MISSPELL) ; $(info $(M) running misspell) @ ## Finds commonly misspelled English words
	$Q $(MISSPELL) .

.PHONY: test
test: ; $(info $(M) running go test) @ ## Runs unit tests
	$Q $(GO) test ${PKGS}

.PHONY: fmt
fmt: ; $(info $(M) running gofmt) @ ## Formats Go code
	$Q $(GO) fmt $(PKGS)

.PHONY: vet
vet: ; $(info $(M) running go vet) @ ## Examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string
	$Q $(GO) vet $(PKGS)

.PHONY: race
race: ; $(info $(M) running go race) @ ## Runs tests with data race detection
	$Q CGO_ENABLED=1 $(GO) test -race -short $(PKGS)

.PHONY: benchmark
benchmark: ; $(info $(M) running go benchmark test) @ ## Benchmark tests to examine performance
	$Q $(GO) test -run=__absolutelynothing__ -bench=. $(PKGS)

.PHONY: coverage
coverage: ; $(info $(M) running go coverage) @ ## Runs tests and generates code coverage report at ./test/coverage.out
	$Q mkdir -p $(CURDIR)/test/
	$Q CGO_ENABLED=1 $(GO) test -race -coverprofile="$(CURDIR)/test/coverage.out" $(PKGS)

.PHONY: coverreport
coverreport: coverage; $(info $(M) running go tool cover) @ ## Runs code coverage and opens the report in a browser
	$Q $(GO) tool cover -html="$(CURDIR)/test/coverage.out"

.PHONY: checkfmt
checkfmt: ; $(info $(M) running checkfmt) @ ## Checks if code is formatted with go fmt and errors out if not
	@test "$(shell $(SRC_FILES) gofmt -l)" = "" \
    || { echo "Code not formatted, please run 'make fmt'"; exit 2; }

# Misc

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m∙ %s:\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:	## Print version information
	@echo App: $(VERSION)
	@echo Go: $(GOVERSION)
	@echo Commit: $(COMMIT)
	@echo Branch: $(BRANCH)

.PHONY: clean
clean: ; $(info $(M) cleaning)	@ ## Cleanup everything
	@rm -rfv $(TBIN)
	@rm -rfv $(CURDIR)/test