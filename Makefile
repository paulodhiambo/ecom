createmigration:
	migrate create -ext sql -dir database/migration -seq init_schema

migrateup:
	migrate -path database/migration -database "postgresql://paulodhiambo:37119787@localhost:5432/ecom?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://paulodhiambo:37119787@localhost:5432/ecom?sslmode=disable" -verbose down
sqlcgenerate:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: createmigration migrateup migratedown sqlcgenerate test