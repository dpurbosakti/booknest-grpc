DB_URL=postgresql://root:mokopass@localhost:5432/booknest_grpc?sslmode=disable

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## createdb: create database
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root booknest_grpc

## new_migration: init sql migration file with name as parameter
new_migration:
	migrate create -ext sql -dir internal/db/migration -seq $(name)

## migrateup: migrate all up schema sql
migrateup:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

## migrateup1: migrate up schema sql by 1
migrateup1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up 1

## migratedown: migrate all down schema sql
migratedown:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

## migratedown1: migrate down schema sql by 1
migratedown1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down 1


## postgresrun: start docker postgres
postgresrun:
	docker start postgres14

## redisrun: start docker redis 
redisrun:
	docker start redis

## sqlc: generate repository code from query
sqlc:
	sqlc generate

## proto: generate go code from proto
proto:
	rm -f internal/pb/*.go
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
    --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative \
    internal/proto/*.proto

## server: run server
server:
	go run cmd/app/main.go

.PHONY: help createdb new_migration migrateup migrateup1 migratedown migratedown1 postgresrun redisrun proto server sqlc