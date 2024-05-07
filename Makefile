githooks:
	git config core.hooksPath .githooks

test_behavior:
	go test godog_test.go

test_unit:
	go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test_all:
	go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go
	CGO_ENABLED=0 go build -o ./bin/migrate ./cmd/migrate/migrate.go

watch:
	~/go/bin/air -d -c .air.toml

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
	golangci-lint --issues-exit-code 1 run  ./...

staticcheck:
	staticcheck ./...

gosec:
	gosec ./...
