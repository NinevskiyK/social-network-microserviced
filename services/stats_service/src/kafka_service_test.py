import time

import clickhouse_driver

from kafka import KafkaProducer, KafkaConsumer, TopicPartition, KafkaAdminClient
from kafka.admin import NewTopic

from testcontainers.clickhouse import ClickHouseContainer
from testcontainers.kafka import KafkaContainer

def connect_kafka(server, client):
    s = ""
    with open("../scripts/init_db/init_likes.sh", 'r') as f:
        for i, l in enumerate(f.readlines()):
            if i < 4 or i == 50:
                continue
            s += l
    s = s.replace('kafka:9092', server)
    s = s.replace(';', '')
    for e in s.split('\n\n'):
        print(e, end='\n---------------------\n')
        client.execute(e)


with ClickHouseContainer("clickhouse/clickhouse-server:21.8", dbname='default') as clickhouse:
    client = clickhouse_driver.Client.from_url(clickhouse.get_connection_url())
    with KafkaContainer() as kafka:
        topic_name = 'likes'
        bootstrap_server = kafka.get_bootstrap_server()

        admin = KafkaAdminClient(bootstrap_servers=[bootstrap_server])

        topic = NewTopic(name=topic_name,
                         num_partitions=1,
                         replication_factor=1)
        admin.create_topics([topic])

        print('created topic')

        connect_kafka(bootstrap_server, client)
        time.sleep(5)

        print('connected!')

        producer = KafkaProducer(bootstrap_servers=[bootstrap_server])
        producer.send(topic_name, b'{"post_id": "p_id", "author_id": "a_id", "user_id": "u_id"}')
        time.sleep(5)
        producer.close()

        print('sent data')

        consumer = KafkaConsumer(bootstrap_servers=[bootstrap_server], auto_offset_reset='earliest', consumer_timeout_ms=1000, group_id='test')
        consumer.subscribe([topic_name])

        print('subscribed')
        for msg in consumer:
            print(msg)
            break

        time.sleep(60)
        res = client.execute('SELECT * FROM likes')
        print(res, type(res))
