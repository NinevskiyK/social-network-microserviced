from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ErrorEnum(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OK: _ClassVar[ErrorEnum]
    NO_SUCH_POST: _ClassVar[ErrorEnum]
    ACCESS_DENIED: _ClassVar[ErrorEnum]
OK: ErrorEnum
NO_SUCH_POST: ErrorEnum
ACCESS_DENIED: ErrorEnum

class Post(_message.Message):
    __slots__ = ("post_id", "user_id", "post_title", "post_text")
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    POST_TITLE_FIELD_NUMBER: _ClassVar[int]
    POST_TEXT_FIELD_NUMBER: _ClassVar[int]
    post_id: str
    user_id: str
    post_title: str
    post_text: str
    def __init__(self, post_id: _Optional[str] = ..., user_id: _Optional[str] = ..., post_title: _Optional[str] = ..., post_text: _Optional[str] = ...) -> None: ...

class Error(_message.Message):
    __slots__ = ("error",)
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: ErrorEnum
    def __init__(self, error: _Optional[_Union[ErrorEnum, str]] = ...) -> None: ...

class PostResponse(_message.Message):
    __slots__ = ("error", "post")
    ERROR_FIELD_NUMBER: _ClassVar[int]
    POST_FIELD_NUMBER: _ClassVar[int]
    error: ErrorEnum
    post: Post
    def __init__(self, error: _Optional[_Union[ErrorEnum, str]] = ..., post: _Optional[_Union[Post, _Mapping]] = ...) -> None: ...

class Pagination(_message.Message):
    __slots__ = ("offset", "limit")
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    offset: int
    limit: int
    def __init__(self, offset: _Optional[int] = ..., limit: _Optional[int] = ...) -> None: ...

class PostRequest(_message.Message):
    __slots__ = ("post_id", "requester_id")
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    REQUESTER_ID_FIELD_NUMBER: _ClassVar[int]
    post_id: str
    requester_id: str
    def __init__(self, post_id: _Optional[str] = ..., requester_id: _Optional[str] = ...) -> None: ...

class PaginatedPostRequest(_message.Message):
    __slots__ = ("requester_id", "pagination")
    REQUESTER_ID_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    requester_id: str
    pagination: Pagination
    def __init__(self, requester_id: _Optional[str] = ..., pagination: _Optional[_Union[Pagination, _Mapping]] = ...) -> None: ...

class PostId(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...
