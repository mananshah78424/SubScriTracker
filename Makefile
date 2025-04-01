
.PHONY: create-migration
create-migration:
	@echo "---> Creating a new migration"
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a migration name. Use: make create-migration name=your_migration_name"; \
		exit 1; \
	fi
	@migrate create -ext sql -dir ./cmd/migrations/ $(name)
# make create-migration name=add_subscriptions_table

.PHONY: db-migrate
db-migrate:
	@echo "Running migrations..."
	@go run cmd/migrations/main.go up

.PHONY: db-rollback
db-rollback:
	@echo "Rolling back migrations..."
	@go run cmd/migrations/main.go down
