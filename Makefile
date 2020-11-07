test_behavior:
	docker-compose up -d --force-recreate
	go test godog_test.go
	docker-compose stop

test_unit:
	go test -v -cover -coverpkg=./... ./pkg/...

test_all:
	gotest -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build:
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go
	CGO_ENABLED=0 go build -o ./bin/migrate ./cmd/migrate/migrate.go

watch:
	~/go/bin/air -d -c .air.toml

start:
	go run ./cmd/app/main.go

db_run:
	docker-compose up -d --force-recreate

migrate:
	go run cmd/migrate/migrate.go
