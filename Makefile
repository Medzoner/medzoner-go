test_behavior:
	docker-compose up -d --force-recreate
	go test godog_test.go
	docker-compose stop

test_unit:
	go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test_all:
	go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
<<<<<<< HEAD
	go build -o bin/app ./cmd/app/app.go
	go build -o bin/app ./cmd/migrate/migrate.go
=======
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go
	CGO_ENABLED=0 go build -o ./bin/migrate ./cmd/migrate/migrate.go

watch:
	~/go/bin/air -d -c .air.toml
>>>>>>> 54472a5b7c0d65228f7e2f9f710ea6f3096ac00f

start:
	go run ./cmd/app/main.go

db_run:
	docker-compose up -d --force-recreate

migrate:
	go run cmd/migrate/migrate.go
