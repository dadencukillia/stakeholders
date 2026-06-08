.PHONY: init_all generate_sqlc test_shared

shared/db/sqlc: shared/db/migrations shared/db/queries sqlc.yaml
	sqlc generate

# Shortcuts

init_all: generate_sqlc

generate_sqlc: shared/db/sqlc

# Tests

test_shared:
	go test ./shared/...
