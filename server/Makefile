sqlc:
	sqlc generate

migrate-create:
	migrate create -ext sql -dir pkg/db/migrations

migrate-up:
	migrate -path "pkg/db/migrations" -database "postgres://postgres:postgres@localhost:5432/connect_chat?sslmode=disable" up

server:
	go run main.go

.PHONY: server sqlc migrate-create migrate-up