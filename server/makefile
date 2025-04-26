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

dev: client
    ENV=development PORT=8000 air -c .air.toml

database-down:
	scp -i ~/.ssh/farmec.pem -r ${REMOTE_HOST}:/home/seanwelch/server/bin/database/* ./database/

database-up:
	scp -i ~/.ssh/farmec.pem -r ./database/* ${REMOTE_HOST}:/home/seanwelch/server/bin/database/

sqlc:
	cd sql && sqlc generate

client:
	cd ../client && npm run embed

baseline:
	goose -dir sql/migrations/ create baseline_schema sql

schema:
	sqlite3 database/database.db '.schema' > sql/current_schema.sql

# Check migration status
migrate-status:
	goose -dir sql/migrations/ sqlite3 database/database.db status

# Apply all migrations
migrate-up:
	goose -dir sql/migrations/ sqlite3 database/database.db up

# Roll back most recent migration
migrate-down:
	goose -dir sql/migrations/ sqlite3 database/database.db down

# Apply one migration forward
migrate-up-one:
	goose -dir sql/migrations/ sqlite3 database/database.db up-by-one

# Create a new migration
migrate-create:
	goose -dir sql/migrations/ create $(name) sql

# Reset database (down all migrations then up all)
migrate-reset:
	goose -dir sql/migrations/ sqlite3 database/database.db reset

# Fix version conflicts (if needed)
migrate-fix:
	goose -dir sql/migrations/ sqlite3 database/database.db fix