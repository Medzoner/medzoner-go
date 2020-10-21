# Medzoner.com
[![Build Status](https://travis-ci.com/Medzoner/medzoner-go.svg?token=USx1h5scpzCMKrhJnEzv&branch=master)](https://travis-ci.com/github/Medzoner/medzoner-go)
[![Coverage Status](https://coveralls.io/repos/github/Medzoner/medzoner-go/badge.svg?branch=master)](https://coveralls.io/github/Medzoner/medzoner-go?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)
[![Go report](https://goreportcard.com/badge/github.com/Medzoner/medzoner-go)](https://goreportcard.com/report/github.com/Medzoner/medzoner-go)

## Short Description
My website https://www.medzoner.com rewrite in golang from php (https://github.com/Medzoner/medzoner.com)

## Configuration

Configuration can be specified in .env or exported as environnement variable:

| Name  | Value |
| ------------- | ------------- |
| ENV  | string  |
| DEBUG  | bool  |
| OPTIONS  | []string{} (option1,option2....)  |
| DATABASE_DSN  | string (ex: root:changeme@tcp(database:3306))  |
| DATABASE_NAME  | string  |
| DATABASE_DRIVER  | string  |
| MAILER_USER  | string  |
| MAILER_PASSWORD  | string  |

## Build
```
go build -o bin/app ./cmd/app.go
go build -o bin/migrate ./cmd/migrate.go
```

## Run
```
./bin/migrate
./bin/app
```

## Test

### UnitTest
```
go test -v -cover -coverpkg=./... ./pkg/...
```

### BehaviorTest
Database (mysql) must be up and configured before run.
```
go test godog_test.go
```

## Developed & Maintained by
[Mehdi Youb](https://github.com/Medzoner) 

## License 
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)
