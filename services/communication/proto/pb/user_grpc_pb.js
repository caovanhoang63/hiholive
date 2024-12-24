// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var user_pb = require('./user_pb.js');
var image_pb = require('./image_pb.js');

function serialize_pb_CreateUserReq(arg) {
  if (!(arg instanceof user_pb.CreateUserReq)) {
    throw new Error('Expected argument of type pb.CreateUserReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_CreateUserReq(buffer_arg) {
  return user_pb.CreateUserReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_GetUserByIdReq(arg) {
  if (!(arg instanceof user_pb.GetUserByIdReq)) {
    throw new Error('Expected argument of type pb.GetUserByIdReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GetUserByIdReq(buffer_arg) {
  return user_pb.GetUserByIdReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_GetUserRoleReps(arg) {
  if (!(arg instanceof user_pb.GetUserRoleReps)) {
    throw new Error('Expected argument of type pb.GetUserRoleReps');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GetUserRoleReps(buffer_arg) {
  return user_pb.GetUserRoleReps.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_GetUserRoleReq(arg) {
  if (!(arg instanceof user_pb.GetUserRoleReq)) {
    throw new Error('Expected argument of type pb.GetUserRoleReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GetUserRoleReq(buffer_arg) {
  return user_pb.GetUserRoleReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_GetUsersByIdsReq(arg) {
  if (!(arg instanceof user_pb.GetUsersByIdsReq)) {
    throw new Error('Expected argument of type pb.GetUsersByIdsReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GetUsersByIdsReq(buffer_arg) {
  return user_pb.GetUsersByIdsReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_NewUserIdResp(arg) {
  if (!(arg instanceof user_pb.NewUserIdResp)) {
    throw new Error('Expected argument of type pb.NewUserIdResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_NewUserIdResp(buffer_arg) {
  return user_pb.NewUserIdResp.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_PublicUserInfoResp(arg) {
  if (!(arg instanceof user_pb.PublicUserInfoResp)) {
    throw new Error('Expected argument of type pb.PublicUserInfoResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_PublicUserInfoResp(buffer_arg) {
  return user_pb.PublicUserInfoResp.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_PublicUsersInfoResp(arg) {
  if (!(arg instanceof user_pb.PublicUsersInfoResp)) {
    throw new Error('Expected argument of type pb.PublicUsersInfoResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_PublicUsersInfoResp(buffer_arg) {
  return user_pb.PublicUsersInfoResp.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserServiceService = exports.UserServiceService = {
  createUser: {
    path: '/pb.UserService/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.CreateUserReq,
    responseType: user_pb.NewUserIdResp,
    requestSerialize: serialize_pb_CreateUserReq,
    requestDeserialize: deserialize_pb_CreateUserReq,
    responseSerialize: serialize_pb_NewUserIdResp,
    responseDeserialize: deserialize_pb_NewUserIdResp,
  },
  getUserRole: {
    path: '/pb.UserService/GetUserRole',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUserRoleReq,
    responseType: user_pb.GetUserRoleReps,
    requestSerialize: serialize_pb_GetUserRoleReq,
    requestDeserialize: deserialize_pb_GetUserRoleReq,
    responseSerialize: serialize_pb_GetUserRoleReps,
    responseDeserialize: deserialize_pb_GetUserRoleReps,
  },
  getUserById: {
    path: '/pb.UserService/GetUserById',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUserByIdReq,
    responseType: user_pb.PublicUserInfoResp,
    requestSerialize: serialize_pb_GetUserByIdReq,
    requestDeserialize: deserialize_pb_GetUserByIdReq,
    responseSerialize: serialize_pb_PublicUserInfoResp,
    responseDeserialize: deserialize_pb_PublicUserInfoResp,
  },
  getUsersByIds: {
    path: '/pb.UserService/GetUsersByIds',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUsersByIdsReq,
    responseType: user_pb.PublicUsersInfoResp,
    requestSerialize: serialize_pb_GetUsersByIdsReq,
    requestDeserialize: deserialize_pb_GetUsersByIdsReq,
    responseSerialize: serialize_pb_PublicUsersInfoResp,
    responseDeserialize: deserialize_pb_PublicUsersInfoResp,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
