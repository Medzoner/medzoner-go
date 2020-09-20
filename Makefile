test_behavior:
	docker-compose up -d --force-recreate
	go test godog_test.go
	docker-compose stop

test_unit:
	go test -v -cover -coverpkg=./... ./pkg/...

build:
	go build -o bin/app ./cmd/app/app.go

start:
	go run ./cmd/app/app.go

db_run:
	docker-compose up -d --force-recreate

migrate:
	go run cmd/migrate/migrate.go
