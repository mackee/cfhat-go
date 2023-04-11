.PHONY: dev
dev:
	reflex -g '*.go' -s -- wrangler dev

.PHONY: build
build:
	mkdir -p dist
	GOOS=js GOARCH=wasm go build -o ./dist/app.wasm .

.PHONY: publish
publish:
	wrangler publish

.PHONY: init-db
init-db:
	wrangler d1 execute cfhat-chatdb --file=./schema.sql
