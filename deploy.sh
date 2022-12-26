#!/bin/bash
make stop;

mkdir images;
mkdir images/lvl_1;
mkdir images/lvl_2;
mkdir images/lvl_3;

make build;
make up;

sleep 30;
make migrate_up;