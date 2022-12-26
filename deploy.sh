#!/bin/bash
make stop;

make build;
make up;

sleep 25;
make migrate_up;