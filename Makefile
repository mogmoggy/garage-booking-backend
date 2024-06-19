postgres:
	docker run --name postgres_garage -p 5430:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres_garage createdb --username=postgres --owner=postgres safekab_garage

dropdb:
	docker exec -it postgres_garage dropdb safekab_garage

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5430/safekab_garage?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5430/safekab_garage?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5430/safekab_garage?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5430/safekab_garage?sslmode=disable" -verbose up 1

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go example/simplebank/db/sqlc Store

api:
	docker run --name safekab_go --network safekab-network -p 8080:8080 -e DB_SOURCE="postgresql://postgres:secret@postgres_garage:5430/safekab_garage?sslmode=disable" -d safekab_garage:latest

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1 api