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

dev:
	ENV=development PORT=8000 air

database-down:
	scp -i ~/.ssh/farmec.pem -r ${REMOTE_HOST}:/home/seanwelch/server/bin/database/* ./database/

database-up:
	scp -i ~/.ssh/farmec.pem -r ./database/* ${REMOTE_HOST}:/home/seanwelch/server/bin/database/

templ:
	templ generate

sqlc:
	cd sql && sqlc generate

js:
	@cp ./static/js/*.js ./public/static/js/
	@cp ./static/fonts/* ./public/static/fonts/

CSS_FILES := $(wildcard static/styles/*.css)
CSS_OUTPUT := $(patsubst static/styles/%.css,public/static/styles/%.css,$(CSS_FILES))

tailwind: $(CSS_OUTPUT)

public/static/styles/%.css: static/styles/%.css
	@mkdir -p public/static/styles
	@tailwindcss -i $< -o $@
