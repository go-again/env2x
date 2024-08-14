.PHONY: fmt build release mod $(PLATFORMS) shasum clean
GOFMT_FILES ?= $$(find . -name '?*.go' -maxdepth 2)
NAME := env2x
PLATFORMS ?= darwin/amd64 darwin/arm64 linux/amd64 linux/arm64
VERSION ?= $(shell git describe --tags --always)
VER ?= $(shell echo $(VERSION)|sed "s/^v\([0-9.]*\).*/\1/")

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

BASE := $(NAME)
RELEASE_DIR := ./release

default: build

fmt:
	go fmt ./...

mod:
	go mod download
	go mod tidy

build: fmt
	go build -trimpath -ldflags="-s -w"
	ln -sf env2x env2json
	ln -sf env2x env2yaml
	ln -sf env2x env2env

release: clean $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -trimpath -ldflags="-s -w" \
	    -o "$(RELEASE_DIR)/$(BASE)-$(os)-$(arch)/$(NAME)"

	cp README.md LICENSE.txt $(RELEASE_DIR)/$(BASE)-$(os)-$(arch)
	cd $(RELEASE_DIR)/$(BASE)-$(os)-$(arch)/ && \
		ln -sf env2x env2json && \
		ln -sf env2x env2yaml && \
		ln -sf env2x env2env && \
		tar czf ../$(BASE)-$(os)-$(arch).tgz .
	rm -rf $(RELEASE_DIR)/$(BASE)-$(os)-$(arch)

clean:
	rm -rf release/
