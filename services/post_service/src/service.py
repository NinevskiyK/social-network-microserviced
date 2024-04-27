import uuid
from typing import AsyncIterator
import logging

from grpc_stub.post_service_pb2_grpc import PostServiceServicer
from grpc_stub.post_service_pb2 import Post, PostRequest, PaginatedPostRequest, PostResponse, PostId, Error, ErrorEnum

import db

class PostService(PostServiceServicer):
    def __init__(self):
        super().__init__()
        self.logger = logging.getLogger(__file__ + '.' + __name__)

    async def CreatePost(self, post: Post, context) -> PostId:
        self.logger.debug('Create post: %s', str(post).replace('\n', '; '))
        post.post_id = str(uuid.uuid4())
        _id = await db.insert_post(post)
        return PostId(id=_id)

    async def GetPost(self, request: PostRequest, context) -> PostResponse:
        post = await db.get_post(request.post_id)
        if post is None:
            return PostResponse(error=ErrorEnum.NO_SUCH_POST)
        if post.user_id != request.requester_id:
            return PostResponse(error=ErrorEnum.ACCESS_DENIED)
        return PostResponse(error=ErrorEnum.OK, post=post)

    async def GetPaginatedPosts(self, request: PaginatedPostRequest, context) -> AsyncIterator[Post]:
        async for post in db.get_paginated(request.requester_id, request.pagination):
            yield post

    async def UpdatePost(self, new_post: Post, context) -> Error:
        self.logger.debug('Update post: %s', str(new_post).replace('\n', '; '))
        post = await db.get_post_and_id(new_post.post_id)
        if post is None:
            return Error(error=ErrorEnum.NO_SUCH_POST)
        post, _id = post
        if post.user_id != new_post.user_id:
            return Error(error=ErrorEnum.ACCESS_DENIED)
        await db.update_post(_id, new_post)
        return Error(error=ErrorEnum.OK)

    async def DeletePost(self, request: PostRequest, context) -> Error:
        post = await db.get_post_and_id(request.post_id)
        if post is None:
            return Error(error=ErrorEnum.NO_SUCH_POST)
        post, _id = post
        if post.user_id != request.requester_id:
            return Error(error=ErrorEnum.ACCESS_DENIED)
        await db.delete_post(_id)
        return Error(error=ErrorEnum.OK)
