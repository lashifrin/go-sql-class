local postgres db is go_sql
migrate create -ext sql -dir migrations create_users
migrate -path migrations -database "postgres://localhost/go_sql?sslmode=disable" up
migrate -path migrations -database "postgres://localhost/go_sql?sslmode=disable" down

using GNU Makefile to build

