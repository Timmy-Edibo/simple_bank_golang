postgres:
	docker run -d --name postgres15 -p 5432:5432 -e POSTGRES_USER=secret -e POSTGRES_PASSWORD=root -d postgres:15-alpine
# docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=secret simple_bank
# docker exec -it postgres15 createdb --username=root --owner=root simple_bank

pgshell:
	docker exec -it postgres15 psql -U postgres simple_bank

dropdb:
# docker exec -it postgres15 dropdb -U root simple_bank
	docker exec -it postgres15 dropdb -U postgres simple_bank

init_migrate:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrateup:
	migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
# 172.17.131.145
test:
	go test -v -cover ./...
sqlc:
	sqlc generate
.PHONY: createdb dropdb postgres pgshell migratedown migrateup init_migrate sqlc test
