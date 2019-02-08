all: test

ci: dependencies test

doc:
	@echo "http://localhost:6060/pkg/github.com/boreq/guinea/"
	godoc -http=:6060

test:
	go test ./...

test-verbose:
	go test -v ./...

test-short:
	go test -short ./...

dependencies:
	go get -t ./...

.PHONY: all ci doc test test-verbose test-short dependencies
