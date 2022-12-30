#!/bin/bash

./api & ./front

wait -n

exit $?