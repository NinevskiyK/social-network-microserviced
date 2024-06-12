from concurrent import futures
from threading import Thread
import time
from flask import Flask

import grpc
from stats_service_pb2 import Id, Count, Type, Posts, TopUsers, TopUser, TopPost
import stats_service_pb2_grpc
import clickhouse_connect

from dotenv import load_dotenv
import os

load_dotenv()

app = Flask(__name__)

@app.route("/")
def healthcheck():
    return "<p>Alive!</p>"

class StatsServiceServicer(stats_service_pb2_grpc.StatsServiceServicer):
    def __init__(self, host='clickhouse', port=8123, username=None, password=None):
        time.sleep(5)

        if username is None:
            username = os.environ['CLICKHOUSE_USER']
        if password is None:
            password = os.environ['CLICKHOUSE_PASSWORD']

        self.client = clickhouse_connect.get_client(host=host, port=port, username=username, password=password, connect_timeout=500)

    def GetStats(self, request: Id, context) -> Count:
        likes = self.client.query(f'SELECT countDistinct(user_id) AS count FROM likes WHERE post_id=\'{request.id}\'')
        views = self.client.query(f'SELECT countDistinct(user_id) AS count FROM views WHERE post_id=\'{request.id}\'')
        return Count(likesCount=likes.result_rows[0][0], viewsCount=views.result_rows[0][0])

    def GetTopPosts(self, request: Type, context) -> Posts:
        table = 'views' if request.isViews else 'likes'
        result = self.client.query(f'SELECT author_id, post_id, countDistinct(user_id) AS count FROM {table} GROUP BY author_id, post_id ORDER BY count DESC LIMIT 5')
        return Posts(posts=[TopPost(postId=str(post_id), authorId=str(author_id), count=int(c)) for author_id, post_id, c in result.result_rows])

    def GetTopUsers(self, request, context) -> TopUsers:
        result = self.client.query('SELECT author_id, sum(user_count) AS count FROM ( SELECT b.author_id, a.user_count FROM ( SELECT post_id, countDistinct(user_id) as user_count FROM likes GROUP BY post_id ) AS a INNER JOIN ( SELECT post_id, author_id FROM likes GROUP BY post_id, author_id ) AS b ON a.post_id = b.post_id ) GROUP BY author_id ORDER BY count DESC LIMIT 3;')
        return TopUsers(users=[TopUser(userIds=str(id), likesCount=int(count)) for id, count in result.result_rows])

def serve_grpc(service):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    stats_service_pb2_grpc.add_StatsServiceServicer_to_server(
        service, server
    )
    server.add_insecure_port("[::]:44455")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    thread = Thread(target=serve_grpc, args=(StatsServiceServicer(), ))
    thread.start()
    app.run("0.0.0.0", "44445")
    thread.join()