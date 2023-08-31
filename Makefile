dev: 
	@go run main.go

dev-watch:
	@reflex -r '\.go$$' -s -- sh -c '$(MAKE) dev'

test:
	@go test -v ./...

test-watch:
	@reflex -r '\.go$$' -s -- sh -c '$(MAKE) test'

build:
	@go build -o bin/ ./...

postgres:
	@docker run --name postgres-url-shortener -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb: 
	@docker exec -it postgres-url-shortener createdb --username=root --owner=root url-shortener

dropdb:
	@docker exec -it postgres-url-shortener dropdb url-shortener

migrateup:
	@migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/url-shortener?sslmode=disable" -verbose up

migratedown:
	@migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/url-shortener?sslmode=disable" -verbose down

sqlc:
	@echo "Generating sqlc"
	@sqlc generate
	@echo "Database queries generated successfully"

