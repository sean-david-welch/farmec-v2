.PHONY: build run test dev database client

include .env
export

REMOTE_HOST := $(REMOTE_SERVER_IP)

build:
	@go build -o bin/main

run: build 
	@./bin/main

test:
	@go test -v ./tests/...

dev:
	cd ../client && npm run embed && cd ../server && ENV=development PORT=8000 air -c .air.toml

database-down:
	scp -i ~/.ssh/farmec.pem -r ${REMOTE_HOST}:/home/seanwelch/server/bin/database/* ./database/

database-up:
	scp -i ~/.ssh/farmec.pem -r ./database/* ${REMOTE_HOST}:/home/seanwelch/server/bin/database/

sqlc:
	cd sql && sqlc generate

baseline:
	goose -dir sql/migrations/ create baseline_schema sql

schema:
	sqlite3 database/database.db '.schema' > sql/current_schema.sql

migrate-status:
	goose -dir sql/migrations/ sqlite3 database/database.db status

migrate-up:
	goose -dir sql/migrations/ sqlite3 database/database.db up

migrate-down:
	goose -dir sql/migrations/ sqlite3 database/database.db down

migrate-up-one:
	goose -dir sql/migrations/ sqlite3 database/database.db up-by-one

migrate-create:
	goose -dir sql/migrations/ create $(name) sql

migrate-reset:
	goose -dir sql/migrations/ sqlite3 database/database.db reset

migrate-fix:
	goose -dir sql/migrations/ sqlite3 database/database.db fix