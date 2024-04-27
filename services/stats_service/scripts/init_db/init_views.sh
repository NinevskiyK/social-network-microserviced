#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
    CREATE TABLE views (
        post_id String,
        user_id String
    ) Engine = MergeTree
    ORDER BY post_id;

    CREATE TABLE views_queue (
        post_id String,
        user_id String
    )
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'kafka',
        kafka_topic_list = 'views',
        kafka_group_name = 'views_stats_service',
        kafka_format = 'JSON',
        kafka_max_block_size = 1048576;

    CREATE MATERIALIZED VIEW views_queue_mv TO views AS
    SELECT post_id, user_id
    FROM views_queue;
EOSQL