.PHONY: githooks test_all build start migrate skaffold-run test trace wire lint staticcheck gosec

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

githooks:
	git config core.hooksPath .githooks

test_all:
	go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
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
	skaffold dev --port-forward

trace:
	go tool trace trace.out

wire:
	wire gen ./pkg/infra/dependency/.

lint:
	#curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.0
	#golangci-lint --version
	golangci-lint --issues-exit-code 1 run  ./...

govet:
	go vet ./...

gofmt:
	.github/gofmt.sh
	#gofmt -s -w .

staticcheck:
	staticcheck ./...

gosec:
	gosec ./...

run-qa: govet gofmt lint staticcheck gosec
	echo "QA passed"
