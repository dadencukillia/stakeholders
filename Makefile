.PHONY: install_tools gen_all gen_sqlc test_shared

shared/db/sqlc: shared/db/migrations shared/db/queries sqlc.yaml
	sqlc generate

# Environment

install_tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

format:
	gofmt -w ./shared

# Generators

gen_all: gen_sqlc

gen_sqlc: shared/db/sqlc

# Tests

test_shared:
	go test -v ./shared/...
