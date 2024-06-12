.PHONY: dev-server dev-tailwind dev-templ dev build-server build-tailwind build-templ build launch deploy

#-----------------------------------------------------
# DEV
#-----------------------------------------------------

dev:
	@make -j dev-server dev-templ dev-tailwind

dev-server:
	@cng -k -i '**/*.go' '**/*.templ' '**/*.html' 'views/public/**/*' -- go run ./server

dev-templ:
	@templ generate --watch

dev-tailwind:
	@make ARGS="--watch" build-tailwind

#-----------------------------------------------------
# BUILD
#-----------------------------------------------------

build:
	@make build-tailwind build-server build-templ

build-server:
	@go build -o bin/server ./server/main.go

build-templ:
	@templ generate

build-tailwind:
	@npx tailwindcss -m -i ./tailwind.css -o ./views/public/styles.css $(ARGS)

.DEFAULT_GOAL := dev  