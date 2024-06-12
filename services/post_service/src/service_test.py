from testcontainers.mongodb import MongoDbContainer

from grpc_stub import post_service_pb2_grpc
from grpc_stub.post_service_pb2 import Post
import grpc

import asyncio
import time
from threading import Thread

from db import get_database
from main import serve


def startService():
    asyncio.run(serve())

with MongoDbContainer("mongo:7.0.7") as mongo:
    # starting service
    client = mongo.get_connection_client()
    get_database(mongo.get_connection_url(), "coll", "db")
    Thread(target=startService, daemon=True).start()

    time.sleep(5)

    with grpc.insecure_channel("localhost:50051") as channel:
        stub = post_service_pb2_grpc.PostServiceStub(channel)
        post_id = stub.CreatePost(Post(post_id="id", user_id="id", post_title="title", post_text="text"))

    time.sleep(5)

    result = client.db.coll.find_one({"postId": post_id.id})

    assert result is not None
    assert result["userId"] == "id"
    assert result["postTitle"] == "title"
    assert result["postText"] == "text"
