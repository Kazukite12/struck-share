postgres:
	docker run --name postgres16 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Ujang123 -d postgres:16-alpine

createdb:
	docker exec -it postgres16  createdb --username=root --owner=root struck-share

dropdb:
	docker exec -it postgres16 dropdb struck-share

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:Ujang123@localhost:5431/struck-share?sslmode=disable" --verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:Ujang123@localhost:5431/struck-share?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v "E:\struck-share:/src" -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server