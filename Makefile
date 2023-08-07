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

# proto:
# 	protoc --proto_path=proto proto/*.proto  --go_out=. --go-grpc_out=.

proto:
	rm -f rpc/*.go
	protoc --proto_path=proto --go_out=rpc --go_opt=paths=source_relative \
	--go-grpc_out=rpc --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=rpc  --grpc-gateway_opt paths=source_relative \
	proto/*.proto

.PHONY: migration migrateup migratedown migrateuplatest migratedownlast sqlc proto
