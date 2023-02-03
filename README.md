## Prerequisites
* In this tutorial postgresql used running locally.
* Database name is go_sql
* To help with data db migrations - used migrate tool.

* To create migrations folder
```
migrate create -ext sql -dir migrations create_users
```

* Following commands will help with up and down scripts
```
migrate -path migrations -database "postgres://localhost/go_sql?sslmode=disable" up
migrate -path migrations -database "postgres://localhost/go_sql?sslmode=disable" down
```

* Makefile used to keep build, etc scripts

