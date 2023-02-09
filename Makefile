postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypassword -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_Bank

dropdb:
	docker exec -it postgres15 dropdb simple_Bank

migrate-status:
	sql-migrate status

migrateup:
	sql-migrate up

migratedown:
	sql-migrate down

sqlcinit:
	sqlc init

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate-status migrateup migratedown sqlcinit sqlc