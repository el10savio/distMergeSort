#!/usr/bin/env bash

docker ps -a | awk '$2 ~ /sort/ {print $1}' | xargs -I {} docker rm -f {}
docker network rm sort_network


