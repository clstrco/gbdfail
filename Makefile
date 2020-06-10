.PHONY: clean test

default: build/gbdfaild
GOFILES := $(shell find . | grep -v vendor | grep go$$ )

build/gbdfaild: pkged.go $(GOFILES)
	go build -o $@ ./cmd/gbdfaild

pkged.go: build/pkgr $(shell find s -type f)
	./build/pkgr -include /s

build/pkgr: vendor
	go build -o build/pkgr ./vendor/github.com/markbates/pkger/cmd/pkger

vendor:
	go mod vendor

clean:
	rm -rf build pkged.go vendor

test: pkged.go
	go test
