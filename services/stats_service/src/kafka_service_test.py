import time
import os

import clickhouse_connect

from dotenv import load_dotenv
load_dotenv()

from kafka import KafkaProducer, KafkaConsumer
from kafka.admin import NewTopic

time.sleep(10)

topic_name = 'likes'
bootstrap_server = 'kafka:9092'
username = os.environ['CLICKHOUSE_USER']
password = os.environ['CLICKHOUSE_PASSWORD']
client = clickhouse_connect.get_client(host='clickhouse', username=username, password=password, connect_timeout=500)

producer = KafkaProducer(bootstrap_servers=[bootstrap_server])
producer.send(topic_name, b'{"post_id": "3c27bd38-7818-4ad9-b89b-38b21580a73f", "author_id": "3c27bd38-7818-4ad9-b89b-38b21580a73f", "user_id": "3c27bd38-7818-4ad9-b89b-38b21580a73f"}')
time.sleep(5)
producer.close()

print('sent data')

time.sleep(5)
res = client.query('SELECT * FROM likes')
print(res.result_rows)
assert str(res.result_rows[0][0]) == '3c27bd38-7818-4ad9-b89b-38b21580a73f'
assert str(res.result_rows[0][1]) == '3c27bd38-7818-4ad9-b89b-38b21580a73f'
assert str(res.result_rows[0][2]) == '3c27bd38-7818-4ad9-b89b-38b21580a73f'
