DB_URL=postgres://postgres:admin@localhost/order_assignment?sslmode=disable
MIGRATIONS_DIR=database/migrations

run:
	@go run main.go

.PHONY: up

up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up
down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down
