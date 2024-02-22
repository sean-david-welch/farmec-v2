build:
	@go build -o bin/main

run: build 
	@./bin/main

test:
	@go test -v ./tests/...

dev:
	ENV=development PORT=80 air
