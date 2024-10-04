import random
import time
import uuid
import clickhouse_driver
from testcontainers.clickhouse import ClickHouseContainer

import stats_service_pb2_grpc
from stats_service_pb2 import TopUsers, TopUser

import grpc
from threading import Thread

from main import StatsServiceServicer, serve_grpc

def poppulate_data(client):
    client.execute('CREATE TABLE likes ( \
            post_id UUID, \
            user_id UUID, \
            author_id UUID \
        ) Engine = MergeTree \
        ORDER BY post_id;')

    authors = ['30524bc5-236a-424a-964f-ddfe791d6b70',
               'de61b8fe-d413-4dc8-a1b2-9afb6e12b48e',
               '4f5b8f7b-3e8d-4072-83e4-9a9837928947',
               '4d284f6f-da85-4aae-ac00-375f2a67bd56']

    count_posts = [2, 3, 5, 3]
    count_likes = [[3, 4], [2, 2, 5], [1, 1, 1, 1, 1], [3, 4, 10]]
    # 7 9 5 17

    for a, cp, cl in zip(authors, count_posts, count_likes):
        author_id = a
        for p, l in zip(range(cp), cl):
            post_id = author_id[:-1] + chr(p)
            for _ in range(l):
                user_id = str(uuid.uuid4())
                for _ in range(random.randint(1, 5)):
                    client.execute(f"INSERT INTO likes (*) VALUES (\
                                   toUUID('{post_id}'), \
                                   toUUID('{user_id}'), \
                                   toUUID('{author_id}') \
                                        )")

with ClickHouseContainer("clickhouse/clickhouse-server:21.8", dbname='default') as clickhouse:
    client = clickhouse_driver.Client.from_url(clickhouse.get_connection_url())
    poppulate_data(client)
    service = StatsServiceServicer(host = clickhouse.get_container_host_ip(),
                                   port=clickhouse.get_exposed_port(8123),
                                   username=clickhouse.username,
                                   password=clickhouse.password)
    Thread(target=serve_grpc, args=(service, ), daemon=True).start()
    time.sleep(5)
    print('start request')
    with grpc.insecure_channel("localhost:44455") as channel:
        stub = stats_service_pb2_grpc.StatsServiceStub(channel)
        empty = stats_service_pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        result : TopUsers = stub.GetTopUsers(request=empty)

    wanted = TopUsers(
        users = [
            TopUser(userIds='4d284f6f-da85-4aae-ac00-375f2a67bd56', likesCount=17),
            TopUser(userIds='de61b8fe-d413-4dc8-a1b2-9afb6e12b48e', likesCount=9),
            TopUser(userIds='30524bc5-236a-424a-964f-ddfe791d6b70', likesCount=7),
        ]
    )

    assert wanted == result
