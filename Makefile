GIT_DESC := $(shell git describe --tags --always --dirty --match "v[0-9]*")
VERSION_TAG := $(patsubst v%,%,$(GIT_DESC))

GO=go
GO_BUILD_OPTS=-ldflags "-X main.versionTag=$(VERSION_TAG)"
GO_COMMON_OPTS=-ldflags "-X main.versionTag=$(VERSION_TAG)" -mod vendor 
GO_FMT=gofmt
GO_TEST=$(GO) test $(GO_COMMON_OPTS) -count=1
GO_BUILD_LINUX=GOOS=linux $(GO) build $(GO_COMMON_OPTS)
GO_SRC_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

BINS := build/bin/rewrite

all: $(BINS)

default: all

build/bin/%:
	$(GO_BUILD_LINUX) -o $@ ./cmd/$*

.PHONY: clean
clean:
	@rm -f $(BINS)
	@rm -rf $(realpath build)

.PHONY: versions
versions:
	@echo "$(CURDIR)"
	@echo "$(BINS)"
	@echo "VERSION_TAG: $(VERSION_TAG)"

