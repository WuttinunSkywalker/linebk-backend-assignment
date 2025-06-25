DATABASE_URL="mysql://root:secret@tcp(localhost:3306)/assignment"

mysql:
	@echo "Creating database..."
	docker run --name mysql8 -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=assignment mysql:8

swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/main.go

server: swagger
	@echo "Starting server..."
	go run cmd/main.go

migrateup: 
	@echo "Running migrations..."
	migrate -path migrations -database $(DATABASE_URL) -verbose up

migratedown: 
	@echo "Running migrations..."
	migrate -path migrations -database $(DATABASE_URL) -verbose down

mock:
	@echo "Generating Mocks"
	go clean -modcache
	mockery

test:
	@echo "Running tests..."
	go test ./...

stress-test:
	@echo "Running stress tests with k6..."
	k6 run k6/stress-test.js

.PHONY: server migrateup migratedown mysql mock test swagger stress-test