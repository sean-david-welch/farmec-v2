.PHONY: build run test dev database

include .env
export

REMOTE_HOST := $(REMOTE_SERVER_IP)

build:
	@go build -o bin/main

run: build 
	@./bin/main

test:
	@go test -v ./tests/...

dev: client
    ENV=development PORT=8000 air -c .air.toml

database-down:
	scp -i ~/.ssh/farmec.pem -r ${REMOTE_HOST}:/home/seanwelch/server/bin/database/* ./database/

database-up:
	scp -i ~/.ssh/farmec.pem -r ./database/* ${REMOTE_HOST}:/home/seanwelch/server/bin/database/

templ:
	templ generate

sqlc:
	cd sql && sqlc generate

.PHONY: client

client:
	cd ../client && npm run embed

baseline:
	goose -dir sql/migrations/ create baseline_schema sql

schema:
	sqlite3 database/database.db '.schema' > sql/current_schema.sql