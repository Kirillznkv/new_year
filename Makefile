DOCS=docker-compose.yml

MIGRATIONS_DIR = ./api/db
DB_URL = "host=localhost port=5432 user=kshanti password=wtrfP9397k19Xk dbname=new_year sslmode=disable"

.PHONY: build
build:
		docker-compose -f $(DOCS) build

.PHONY: up
up:
		docker-compose -f $(DOCS) up -d

.PHONY: down
down:
		docker-compose -f $(DOCS) down

.PHONY: stop
stop: down
		docker stop $$(docker ps -aq)
		docker rm $$(docker ps -aq)

.PHONY: migrate_up
migrate_up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

.PHONY: migrate_down
migrate_down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

.PHONY: fclean
fclean: stop
		sudo rm -rf ./pgdata ./images ./texts
		#docker volume rm tarot-cards-tgbot_postgres -f