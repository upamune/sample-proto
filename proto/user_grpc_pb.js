// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_user_pb = require('../proto/user_pb.js');

function serialize_proto_CreateRequest(arg) {
  if (!(arg instanceof proto_user_pb.CreateRequest)) {
    throw new Error('Expected argument of type proto.CreateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_CreateRequest(buffer_arg) {
  return proto_user_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateResponse(arg) {
  if (!(arg instanceof proto_user_pb.CreateResponse)) {
    throw new Error('Expected argument of type proto.CreateResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_CreateResponse(buffer_arg) {
  return proto_user_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_DeleteRequest(arg) {
  if (!(arg instanceof proto_user_pb.DeleteRequest)) {
    throw new Error('Expected argument of type proto.DeleteRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_DeleteRequest(buffer_arg) {
  return proto_user_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_DeleteResponse(arg) {
  if (!(arg instanceof proto_user_pb.DeleteResponse)) {
    throw new Error('Expected argument of type proto.DeleteResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_DeleteResponse(buffer_arg) {
  return proto_user_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetAllRequest(arg) {
  if (!(arg instanceof proto_user_pb.GetAllRequest)) {
    throw new Error('Expected argument of type proto.GetAllRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_GetAllRequest(buffer_arg) {
  return proto_user_pb.GetAllRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetAllResponse(arg) {
  if (!(arg instanceof proto_user_pb.GetAllResponse)) {
    throw new Error('Expected argument of type proto.GetAllResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_GetAllResponse(buffer_arg) {
  return proto_user_pb.GetAllResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetRequest(arg) {
  if (!(arg instanceof proto_user_pb.GetRequest)) {
    throw new Error('Expected argument of type proto.GetRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_GetRequest(buffer_arg) {
  return proto_user_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetResponse(arg) {
  if (!(arg instanceof proto_user_pb.GetResponse)) {
    throw new Error('Expected argument of type proto.GetResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_GetResponse(buffer_arg) {
  return proto_user_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateRequest(arg) {
  if (!(arg instanceof proto_user_pb.UpdateRequest)) {
    throw new Error('Expected argument of type proto.UpdateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_UpdateRequest(buffer_arg) {
  return proto_user_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateResponse(arg) {
  if (!(arg instanceof proto_user_pb.UpdateResponse)) {
    throw new Error('Expected argument of type proto.UpdateResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_UpdateResponse(buffer_arg) {
  return proto_user_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserServiceService = exports.UserServiceService = {
  // CREATE
  create: {
    path: '/proto.UserService/Create',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_pb.CreateRequest,
    responseType: proto_user_pb.CreateResponse,
    requestSerialize: serialize_proto_CreateRequest,
    requestDeserialize: deserialize_proto_CreateRequest,
    responseSerialize: serialize_proto_CreateResponse,
    responseDeserialize: deserialize_proto_CreateResponse,
  },
  // READ
  get: {
    path: '/proto.UserService/Get',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_pb.GetRequest,
    responseType: proto_user_pb.GetResponse,
    requestSerialize: serialize_proto_GetRequest,
    requestDeserialize: deserialize_proto_GetRequest,
    responseSerialize: serialize_proto_GetResponse,
    responseDeserialize: deserialize_proto_GetResponse,
  },
  getAll: {
    path: '/proto.UserService/GetAll',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_pb.GetAllRequest,
    responseType: proto_user_pb.GetAllResponse,
    requestSerialize: serialize_proto_GetAllRequest,
    requestDeserialize: deserialize_proto_GetAllRequest,
    responseSerialize: serialize_proto_GetAllResponse,
    responseDeserialize: deserialize_proto_GetAllResponse,
  },
  // UPDATE
  update: {
    path: '/proto.UserService/Update',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_pb.UpdateRequest,
    responseType: proto_user_pb.UpdateResponse,
    requestSerialize: serialize_proto_UpdateRequest,
    requestDeserialize: deserialize_proto_UpdateRequest,
    responseSerialize: serialize_proto_UpdateResponse,
    responseDeserialize: deserialize_proto_UpdateResponse,
  },
  // DELETE
  delete: {
    path: '/proto.UserService/Delete',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_pb.DeleteRequest,
    responseType: proto_user_pb.DeleteResponse,
    requestSerialize: serialize_proto_DeleteRequest,
    requestDeserialize: deserialize_proto_DeleteRequest,
    responseSerialize: serialize_proto_DeleteResponse,
    responseDeserialize: deserialize_proto_DeleteResponse,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
