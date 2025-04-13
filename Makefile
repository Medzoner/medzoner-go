.PHONY: githooks test_all build start migrate skaffold-run test trace wire lint staticcheck gosec

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

githooks:
	git config core.hooksPath .githooks

test_all:
	export DEBUG=true
	go test -v -cover -coverpkg=./... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go
	CGO_ENABLED=0 go build -o ./bin/migrate ./cmd/migrate/migrate.go

start:
	go run ./cmd/app/main.go

migrate:
	go run ./cmd/migrate/migrate.go

skaffold-run:
	skaffold dev --port-forward --platform=linux/arm64,linux/amd64 --insecure-registry=registry.medzoner.lan:5000

trace:
	go tool trace trace.out

wire:
	wire gen ./pkg/infra/dependency/.

lint:
	#curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.0
	#golangci-lint --version
	golangci-lint -v --config .golangci.v1.yml --issues-exit-code 1 run $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

lint-fix:
	#fieldalignment -fix -test=false ./...
	golangci-lint -v --config .golangci.v1.yml --issues-exit-code=1 run --fix $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

govet:
	go vet $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

gofmt:
	.github/gofmt.sh

staticcheck:
	staticcheck --debug.version
	staticcheck $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

gosec:
	gosec $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

ineffassign:
	ineffassign $(go list -e -f '{{.Dir}}' ./... | grep -v '/var/')

gocyclo:
	gocyclo -ignore "_test|Godeps|var|vendor/" .

run-qa: govet gofmt lint staticcheck gocyclo
	echo "QA passed"

#go install github.com/go-critic/go-critic/cmd/gocritic@latest
#gocritic check ./...

trivy:
	trivy image

k6:
	docker run --rm -i grafana/k6 run  --vus 10 --duration 30s  - <k6/test.js
