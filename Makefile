DB_URL=postgres://postgres:postgres@localhost:5432/go_notifications?sslmode=disable

migration:
	migrate create -ext sql -dir db/migrations

migrateup:
	migrate -path "db/migrations" -database "$(DB_URL)" up

migrateuplatest:
	migrate -path "db/migrations" -database "$(DB_URL)" up 1

migratedown:
	migrate -path "db/migrations" -database "$(DB_URL)" down

migratedownlast:
	migrate -path "db/migrations" -database "$(DB_URL)" down 1

sqlc:
	sqlc generate

start:
	go run cmd/main.go

evans:
	evans --host localhost --port 6000 -r repl

proto:
	rm -f rpcs/*.go
	protoc --proto_path=proto --go_out=rpcs --go_opt=paths=source_relative \
	--go-grpc_out=rpcs --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=rpcs  --grpc-gateway_opt paths=source_relative \
	proto/*.proto

.PHONY: migration migrateup migratedown migrateuplatest migratedownlast sqlc proto start evans
