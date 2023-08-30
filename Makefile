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
