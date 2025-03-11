postgres:
	docker run --name docker_postgres -e POSTGRES_PASSWORD=12345 -p 5678:5432 -d postgres
createdb:
	docker exec -it docker_postgres createdb --username=postgres --owner=postgres swiss_bank
dropdb:
	docker exec -it docker_postgres dropdb swiss_bank
migrate-up:
	migrate -path db/migration -database "postgresql://postgres:12345@localhost:5678/swiss_bank?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://postgres:12345@localhost:5678/swiss_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc test