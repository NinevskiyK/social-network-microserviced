#!/usr/bin/env bash
docker network create social-network-microserviced_grpc
docker compose up --build --detach
docker run --rm --network stats_service_clickhouse_db --network stats_service_kafka $(docker build -q -f TestDockerfile .)
docker compose down
