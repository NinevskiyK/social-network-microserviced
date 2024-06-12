from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TopUser(_message.Message):
    __slots__ = ("userIds", "likesCount")
    USERIDS_FIELD_NUMBER: _ClassVar[int]
    LIKESCOUNT_FIELD_NUMBER: _ClassVar[int]
    userIds: str
    likesCount: int
    def __init__(self, userIds: _Optional[str] = ..., likesCount: _Optional[int] = ...) -> None: ...

class TopUsers(_message.Message):
    __slots__ = ("users",)
    USERS_FIELD_NUMBER: _ClassVar[int]
    users: _containers.RepeatedCompositeFieldContainer[TopUser]
    def __init__(self, users: _Optional[_Iterable[_Union[TopUser, _Mapping]]] = ...) -> None: ...

class Id(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class TopPost(_message.Message):
    __slots__ = ("postId", "authorId", "count")
    POSTID_FIELD_NUMBER: _ClassVar[int]
    AUTHORID_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    postId: str
    authorId: str
    count: int
    def __init__(self, postId: _Optional[str] = ..., authorId: _Optional[str] = ..., count: _Optional[int] = ...) -> None: ...

class Posts(_message.Message):
    __slots__ = ("posts",)
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[TopPost]
    def __init__(self, posts: _Optional[_Iterable[_Union[TopPost, _Mapping]]] = ...) -> None: ...

class Type(_message.Message):
    __slots__ = ("isViews",)
    ISVIEWS_FIELD_NUMBER: _ClassVar[int]
    isViews: bool
    def __init__(self, isViews: bool = ...) -> None: ...

class Count(_message.Message):
    __slots__ = ("likesCount", "viewsCount")
    LIKESCOUNT_FIELD_NUMBER: _ClassVar[int]
    VIEWSCOUNT_FIELD_NUMBER: _ClassVar[int]
    likesCount: int
    viewsCount: int
    def __init__(self, likesCount: _Optional[int] = ..., viewsCount: _Optional[int] = ...) -> None: ...

class StatsRequest(_message.Message):
    __slots__ = ("type", "id")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    type: Type
    id: str
    def __init__(self, type: _Optional[_Union[Type, _Mapping]] = ..., id: _Optional[str] = ...) -> None: ...
