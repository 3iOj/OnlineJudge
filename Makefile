postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root 3iOj
dropdb:
	docker exec -it postgres15 dropdb 3iOj
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/3iOj?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/3iOj?sslmode=disable" --verbose down
sqlc-gen:
	docker run --rm -v $${pwd}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server  