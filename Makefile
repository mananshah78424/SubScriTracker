.PHONY: create-migrate
create-migrate:
	@echo "Creating new migration..."
	@go run cmd/migrations/main.go create $(name)

.PHONY: db-migrate
db-migrate:
	@echo "Running migrations..."
	@go run cmd/migrations/main.go up

.PHONY: db-rollback
db-rollback:
	@echo "Rolling back migrations..."
	@go run cmd/migrations/main.go down
