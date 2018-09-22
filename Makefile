COVERAGE_PROFILE ?= coverage.out
PKGS := ./...

default: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags '-w -s' -o ./bin/app .

.PHONY: clean
clean:
	rm -rf ./bin ./vendor

.PHONY: install
install:
	dep ensure

.PHONY: lint
lint:
	gometalinter --vendor --tests $(PKGS)

.PHONY: setup
setup:
	go get -u -v \
	  github.com/alecthomas/gometalinter \
	  github.com/golang/dep/cmd/dep \
	  github.com/codegangsta/gin
	gometalinter --install

.PHONY: test
test:
	go test ./... -coverprofile $(COVERAGE_PROFILE)
