build:
	docker-compose up --build
migup:
	migrate -database postgres://postgres:12345@localhost:6080/models?sslmode=disable -path ./migrations up
migdown:
	migrate -database postgres://postgres:12345@localhost:6080/models?sslmode=disable -path ./migrations down
down:
	docker-compose down
