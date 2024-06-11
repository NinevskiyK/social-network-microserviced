#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
    CREATE TABLE likes (
        post_id UUID,
        user_id UUID,
        author_id UUID
    ) Engine = MergeTree
    ORDER BY post_id;

    CREATE TABLE likes_and_authors (
        author_id UUID,
        post_id UUID,
        user_id UUID
    ) Engine = MergeTree
    ORDER BY author_id;

    create table likes_queue (
        post_id UUID,
        user_id UUID,
        author_id UUID
    )
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'kafka:9092',
        kafka_topic_list = 'likes',
        kafka_group_name = 'likes_stats_service',
        kafka_format = 'JSON',
        kafka_max_block_size = 1048576;

    create table likes_and_authors_queue (
        post_id UUID,
        user_id UUID,
        author_id UUID
    )
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'kafka:9092',
        kafka_topic_list = 'likes',
        kafka_group_name = 'likes_stats_service1',
        kafka_format = 'JSON',
        kafka_max_block_size = 1048576;

    CREATE MATERIALIZED VIEW likes_queue_mv TO likes AS
    SELECT post_id, user_id, author_id
    FROM likes_queue;

    CREATE MATERIALIZED VIEW likes_and_authors_queue_mv TO likes_and_authors AS
    SELECT post_id, user_id, author_id
    FROM likes_and_authors_queue;
EOSQL
