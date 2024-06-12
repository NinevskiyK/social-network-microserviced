import unittest
from unittest.mock import MagicMock
import uuid
import clickhouse_connect

from stats_service_pb2 import Id, Count, Type, Posts, TopUsers, TopUser, TopPost

import main

class ClientMock:
    pass

class StatsMock:
    def __init__(self):
        self.client = ClientMock

class QueryResultMock:
    def __init__(self, res):
        self.result_rows = res

class Test(unittest.TestCase):
    def __init__(self, *args):
        super().__init__(*args)
        main.StatsServiceServicer.__init__ = StatsMock.__init__
        self.service = main.StatsServiceServicer()

    def test_GetTopUsers(self):
        self.service.client.query = MagicMock(return_value=QueryResultMock([
            (
                uuid.UUID(str('0d22c60b-ccf0-4f6b-86f7-04d39781752a')),
                5
            )]))

        users = self.service.GetTopUsers(None, None)

        self.service.client.query.assert_called_once()

        self.assertEqual(users, TopUsers(users=[TopUser(userIds='0d22c60b-ccf0-4f6b-86f7-04d39781752a', likesCount=5)]))

    def test_GetStats(self):
        self.service.client.query = MagicMock()
        self.service.client.query.side_effect = [QueryResultMock([(2, )]), QueryResultMock([(3, )])]

        posts = self.service.GetStats(Id(id="id"), None)

        self.service.client.query.assert_called()

        self.assertEqual(posts, Count(likesCount=2, viewsCount=3))

if __name__ == '__main__':
    unittest.main()