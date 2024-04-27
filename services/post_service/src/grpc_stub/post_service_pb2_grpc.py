# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import grpc_stub.post_service_pb2 as post__service__pb2


class PostServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreatePost = channel.unary_unary(
                '/PostService/CreatePost',
                request_serializer=post__service__pb2.Post.SerializeToString,
                response_deserializer=post__service__pb2.PostId.FromString,
                )
        self.GetPost = channel.unary_unary(
                '/PostService/GetPost',
                request_serializer=post__service__pb2.PostRequest.SerializeToString,
                response_deserializer=post__service__pb2.PostResponse.FromString,
                )
        self.GetPaginatedPosts = channel.unary_stream(
                '/PostService/GetPaginatedPosts',
                request_serializer=post__service__pb2.PaginatedPostRequest.SerializeToString,
                response_deserializer=post__service__pb2.Post.FromString,
                )
        self.UpdatePost = channel.unary_unary(
                '/PostService/UpdatePost',
                request_serializer=post__service__pb2.Post.SerializeToString,
                response_deserializer=post__service__pb2.Error.FromString,
                )
        self.DeletePost = channel.unary_unary(
                '/PostService/DeletePost',
                request_serializer=post__service__pb2.PostRequest.SerializeToString,
                response_deserializer=post__service__pb2.Error.FromString,
                )


class PostServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreatePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPaginatedPosts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdatePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeletePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PostServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreatePost': grpc.unary_unary_rpc_method_handler(
                    servicer.CreatePost,
                    request_deserializer=post__service__pb2.Post.FromString,
                    response_serializer=post__service__pb2.PostId.SerializeToString,
            ),
            'GetPost': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPost,
                    request_deserializer=post__service__pb2.PostRequest.FromString,
                    response_serializer=post__service__pb2.PostResponse.SerializeToString,
            ),
            'GetPaginatedPosts': grpc.unary_stream_rpc_method_handler(
                    servicer.GetPaginatedPosts,
                    request_deserializer=post__service__pb2.PaginatedPostRequest.FromString,
                    response_serializer=post__service__pb2.Post.SerializeToString,
            ),
            'UpdatePost': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdatePost,
                    request_deserializer=post__service__pb2.Post.FromString,
                    response_serializer=post__service__pb2.Error.SerializeToString,
            ),
            'DeletePost': grpc.unary_unary_rpc_method_handler(
                    servicer.DeletePost,
                    request_deserializer=post__service__pb2.PostRequest.FromString,
                    response_serializer=post__service__pb2.Error.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'PostService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PostService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreatePost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/PostService/CreatePost',
            post__service__pb2.Post.SerializeToString,
            post__service__pb2.PostId.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/PostService/GetPost',
            post__service__pb2.PostRequest.SerializeToString,
            post__service__pb2.PostResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPaginatedPosts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/PostService/GetPaginatedPosts',
            post__service__pb2.PaginatedPostRequest.SerializeToString,
            post__service__pb2.Post.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdatePost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/PostService/UpdatePost',
            post__service__pb2.Post.SerializeToString,
            post__service__pb2.Error.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeletePost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/PostService/DeletePost',
            post__service__pb2.PostRequest.SerializeToString,
            post__service__pb2.Error.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)