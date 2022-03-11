build:
	docker-compose up --build
mig:
	migrate -database postgres://postgres:12345@localhost:6080/models?sslmode=disable -path ./migrations up
down:
	docker-compose down
