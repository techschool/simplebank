DB_URL=postgresql://root:password@localhost:5432/db_simple_bank?sslmode=disable

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

.PHONY: migrateup1
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

.PHONY: sqlc
sqlc:
	sqlc generate