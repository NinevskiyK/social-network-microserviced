from typing import Optional, AsyncIterable, Dict, Tuple
import os
import uuid

import motor.motor_asyncio
from google.protobuf.json_format import ParseDict, MessageToDict

from grpc_stub.post_service_pb2 import Post, Error, ErrorEnum, Pagination


coll = None
def get_database():
    user = os.getenv("MONGO_USER")
    password = os.getenv("MONGO_PASSWORD")
    host = os.getenv("MONGO_HOST")
    port = os.getenv("MONGO_PORT")
    database = os.getenv("MONGO_DATABASE")
    collection = os.getenv("MONGO_COLLECTION")

    connection_string = f"mongodb://{user}:{password}@{host}:{port}"
    client = motor.motor_asyncio.AsyncIOMotorClient(connection_string)
    global coll
    coll = client.get_database(database).get_collection(collection)

async def insert_post(post: Post):
    post = MessageToDict(post)
    res = await coll.insert_one(post)
    return post['postId']

async def get_post_and_id(post_id: str) -> Optional[Tuple[Post, str]]:
    d : Dict = await coll.find_one({'postId': post_id})
    if d is None:
        return None
    _id = d.pop('_id')
    return ParseDict(d, Post()), _id

async def get_post(post_id: str) -> Optional[Post]:
    post = await get_post_and_id(post_id)
    if post is None:
        return None
    post, _id = post
    return post

async def update_post(_id, post: Post):
    await coll.replace_one({"_id": _id}, MessageToDict(post))

async def delete_post(_id):
    await coll.delete_one({"_id": _id})

async def get_paginated(user_id: str, pagination: Pagination) -> AsyncIterable[Post]:
    cursor =  coll.find({'userId': user_id})\
                            .sort("postId")\
                            .skip(pagination.offset)
    if pagination.limit != 0:
        cursor = cursor.limit(pagination.limit)
    async for post in cursor:
        post.pop('_id')
        yield ParseDict(post, Post())
