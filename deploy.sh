#!/bin/bash

sudo -s
curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh;
curl -fsSL get.docker.com -o get-docker.sh; sh get-docker.sh;

#curl -L "https://github.com/docker/compose/releases/download/1.26.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose;

#make stop;
#
#mkdir images;
#mkdir images/lvl_1;
#mkdir images/lvl_2;
#mkdir images/lvl_3;
#
#make build;
#make up;
#
#sleep 5;
#make migrate_up;