# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: post_service.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12post_service.proto\"O\n\x04Post\x12\x0f\n\x07post_id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x12\n\npost_title\x18\x03 \x01(\t\x12\x11\n\tpost_text\x18\x04 \x01(\t\"\"\n\x05\x45rror\x12\x19\n\x05\x65rror\x18\x01 \x01(\x0e\x32\n.ErrorEnum\">\n\x0cPostResponse\x12\x19\n\x05\x65rror\x18\x01 \x01(\x0e\x32\n.ErrorEnum\x12\x13\n\x04post\x18\x02 \x01(\x0b\x32\x05.Post\"+\n\nPagination\x12\x0e\n\x06offset\x18\x01 \x01(\x04\x12\r\n\x05limit\x18\x02 \x01(\x04\"4\n\x0bPostRequest\x12\x0f\n\x07post_id\x18\x01 \x01(\t\x12\x14\n\x0crequester_id\x18\x02 \x01(\t\"`\n\x14PaginatedPostRequest\x12\x14\n\x0crequester_id\x18\x01 \x01(\t\x12\x11\n\ttarget_id\x18\x03 \x01(\t\x12\x1f\n\npagination\x18\x02 \x01(\x0b\x32\x0b.Pagination\"\x14\n\x06PostId\x12\n\n\x02id\x18\x01 \x01(\t*8\n\tErrorEnum\x12\x06\n\x02OK\x10\x00\x12\x10\n\x0cNO_SUCH_POST\x10\x01\x12\x11\n\rACCESS_DENIED\x10\x02\x32\xd3\x01\n\x0bPostService\x12\x1e\n\nCreatePost\x12\x05.Post\x1a\x07.PostId\"\x00\x12(\n\x07GetPost\x12\x0c.PostRequest\x1a\r.PostResponse\"\x00\x12\x35\n\x11GetPaginatedPosts\x12\x15.PaginatedPostRequest\x1a\x05.Post\"\x00\x30\x01\x12\x1d\n\nUpdatePost\x12\x05.Post\x1a\x06.Error\"\x00\x12$\n\nDeletePost\x12\x0c.PostRequest\x1a\x06.Error\"\x00\x42\x10Z\x0e./post_serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'post_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\016./post_service'
  _globals['_ERRORENUM']._serialized_start=422
  _globals['_ERRORENUM']._serialized_end=478
  _globals['_POST']._serialized_start=22
  _globals['_POST']._serialized_end=101
  _globals['_ERROR']._serialized_start=103
  _globals['_ERROR']._serialized_end=137
  _globals['_POSTRESPONSE']._serialized_start=139
  _globals['_POSTRESPONSE']._serialized_end=201
  _globals['_PAGINATION']._serialized_start=203
  _globals['_PAGINATION']._serialized_end=246
  _globals['_POSTREQUEST']._serialized_start=248
  _globals['_POSTREQUEST']._serialized_end=300
  _globals['_PAGINATEDPOSTREQUEST']._serialized_start=302
  _globals['_PAGINATEDPOSTREQUEST']._serialized_end=398
  _globals['_POSTID']._serialized_start=400
  _globals['_POSTID']._serialized_end=420
  _globals['_POSTSERVICE']._serialized_start=481
  _globals['_POSTSERVICE']._serialized_end=692
# @@protoc_insertion_point(module_scope)