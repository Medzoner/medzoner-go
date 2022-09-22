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
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go
	CGO_ENABLED=0 go build -o ./bin/migrate ./cmd/migrate/migrate.go

watch:
	~/go/bin/air -d -c .air.toml

start:
	go run ./cmd/app/main.go

db_run:
	docker-compose up -d --force-recreate

migrate:
	go run ./cmd/migrate/migrate.go

install_proto:
	sudo mv ~/Téléchargements/protoc-gen-grpc-web-1.3.1-linux-x86_64 /usr/bin/protoc-gen-grpc-web
	sudoo chmod +x /usr/local/bin/protoc-gen-grpc-web

gen_rpc:
	protoc --js_out=import_style=commonjs:./   --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./ ./calculator/calculator.proto
