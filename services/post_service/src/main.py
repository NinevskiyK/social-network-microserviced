import asyncio
import logging
import grpc

import grpc_stub.post_service_pb2_grpc as pb2_grpc
from service import PostService
import db

async def serve() -> None:
    server = grpc.aio.server()
    pb2_grpc.add_PostServiceServicer_to_server(
        PostService(), server
    )
    server.add_insecure_port("[::]:50051")
    logging.info("started!")
    await server.start()
    await server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    db.get_database()
    asyncio.get_event_loop().run_until_complete(serve())
