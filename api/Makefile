#MIGRATIONS_DIR = ./db
#DB_URL = "host=postgres port=5432 user=kshanti password=wtrfP9397k19Xk dbname=new_year sslmode=disable"

.PHONY: build
build:
	go build -v ./cmd/api.go
	go build -v ./cmd/front.go

#.PHONY: db_up
#db_up:
#	docker run --rm --name my_postgres \
#	-e POSTGRES_PASSWORD=wtrfP9397k19Xk \
#	-e POSTGRES_USER=kshanti \
#	-e POSTGRES_DB=new_year \
#	-p 5432:5432 -d postgres
#
#.PHONY: db_down
#db_down:
#	docker stop my_postgres
#
#.PHONY: migrate_up
#migrate_up:
#	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up
#
#.PHONY: migrate_down
#migrate_down:
#	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

.DEFAULT_GOAL := build