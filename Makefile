DOCS=docker-compose.yml
#HOME=/Users/kshanti/Desktop/tarot-cards-tgbot

#all: create_dir build up

all: build up

#.PHONY: create_dir
#create_dir:
#		mkdir -p $(HOME)/pgdata

.PHONY: build
build:
		docker-compose -f $(DOCS) build

.PHONY: up
up:
		docker-compose -f $(DOCS) up -d

.PHONY: stop
stop:
		docker stop $$(docker ps -aq)
		docker rm $$(docker ps -aq)

#.PHONY: fclean
#fclean: stop
#		sudo rm -rf $(HOME)/pgdata
#		docker volume rm tarot-cards-tgbot_postgres -f