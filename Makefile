.PHONY: generate lint

generate:
	buf generate

lint:
	buf lint
