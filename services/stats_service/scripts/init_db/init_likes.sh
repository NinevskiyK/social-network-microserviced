#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
    CREATE TABLE likes (
        post_id String,
        user_id String
    ) Engine = MergeTree
    ORDER BY post_id;

    CREATE TABLE likes_queue (
        post_id String,
        user_id String
    )
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'kafka:9092',
        kafka_topic_list = 'likes',
        kafka_group_name = 'likes_stats_service',
        kafka_format = 'JSON',
        kafka_max_block_size = 1048576;

    CREATE MATERIALIZED VIEW likes_queue_mv TO likes AS
    SELECT post_id, user_id
    FROM likes_queue;
EOSQL
