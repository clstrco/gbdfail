.PHONY: clean test

default: build/gbdfaild
GOFILES := $(shell find . | grep -v vendor | grep go$$ )

build/gbdfaild: static.go $(GOFILES)
	@go build -o $@ ./cmd/gbdfaild

static.go: build/go-bindata $(shell find s -type f)
	./build/go-bindata -o $@ -fs -pkg gbdfail -prefix "s/" s/...

build/go-bindata: vendor
	go build -o build/go-bindata ./vendor/github.com/go-bindata/go-bindata/go-bindata

vendor:
	go mod vendor

clean:
	rm -rf build static.go vendor

test: static.go
	go test
