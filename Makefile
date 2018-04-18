PKGS := $(shell go list ./... | grep -v /vendor)
BINARY := server
OS ?= linux
ARCH ?= amd64
VERSION ?= latest

.PHONY: dep
dep: ## Get the dependencies
	go get -v -d ./...

.PHONY: test
test: dep
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install > /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor --errors --deadline=2m --fast

.PHONY: docker
docker:
ifeq ($(REGISTRY),)
	docker build --build-arg VERSION=$(VERSION) -t $(BINARY):$(VERSION) -f Dockerfile .
else
	docker build --build-arg VERSION=$(VERSION) -t $(REGISTRY)/$(BINARY):$(VERSION) -f Dockerfile .
	docker push $(REGISTRY)/$(BINARY):$(VERSION)
endif

.PHONY: bin
bin: dep
	mkdir -p build
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -v -ldflags "-X main.version=$(VERSION)" -a -installsuffix cgo -o build/$(BINARY)-$(OS)-$(ARCH)-v$(VERSION) github.com/Kaitsh/herrnhuter-daily

.DEFAULT_GOAL := bin
