.PHONY: init_all gen_sqlc test_shared

shared/db/sqlc: shared/db/migrations shared/db/queries sqlc.yaml
	sqlc generate

# Shortcuts

init_all: gen_sqlc

gen_sqlc: shared/db/sqlc

# Tests

test_shared:
	go test -v ./shared/...
