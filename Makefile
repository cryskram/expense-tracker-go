DB_URL=postgres://postgres:postgres@localhost:5432/expense_tracker?sslmode=disable

run:
	go run ./cmd/api

dev:
	air

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

create-migration:
	migrate create -ext sql -dir migrations -seq $(NAME)