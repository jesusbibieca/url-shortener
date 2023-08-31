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

createpostgres:
	@docker run --name postgres-url-shortener -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createredis:
	@docker run --name redis-url-shortener -p 6379:6379 -d redis:alpine

dockerup:
	@docker-compose up -d

dockerdown:
	@docker-compose down

createdb: 
	@docker exec -it url-shortener-postgres-1 createdb --username=root --owner=root url-shortener

dropdb:
	@docker exec -it url-shortener-postgres-1  dropdb url-shortener

migrateup:
	@migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/url-shortener?sslmode=disable" -verbose up

migratedown:
	@migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/url-shortener?sslmode=disable" -verbose down

sqlc:
	@echo "Generating sqlc"
	@sqlc generate
	@echo "Database queries generated successfully"

