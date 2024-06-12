import asyncio
import db
import service

from grpc_stub.post_service_pb2 import Post

import unittest
from unittest.mock import AsyncMock

class TestService(unittest.TestCase):
    def __init__(self, *args):
        super().__init__(*args)
        self.post_service = service.PostService()

    def test_CreatePost(self):
        post : Post = Post(user_id = 'user_id', post_title='post_title', post_text='post_text')
        db.insert_post = AsyncMock(return_value="id")

        asyncio.run(self.post_service.CreatePost(post, None))

        db.insert_post.assert_called_once_with(post)
        self.assertEqual(post.post_title, 'post_title')
        self.assertEqual(post.user_id, 'user_id')
        self.assertEqual(post.post_text, 'post_text')

    def test_UpdatePost(self):
        post : Post = Post(user_id = 'user_id', post_title='post_title', post_text='post_text')
        db.get_post_and_id = AsyncMock(return_value=(post, "id"))
        db.update_post = AsyncMock()

        asyncio.run(self.post_service.UpdatePost(post, None))

        db.get_post_and_id.assert_called_once()
        db.update_post.assert_called_once_with("id", post)

if __name__ == '__main__':
    unittest.main()