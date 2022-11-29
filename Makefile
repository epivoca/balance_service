postgres:
	docker run --name postgres12 --network service-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root balance_db

dropdb:
	docker exec -it postgres12 dropdb balance_db

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/balance_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/balance_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/epivoca/balance_service/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock