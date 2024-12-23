// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var user_pb = require('./user_pb.js');

function serialize_pb_CreateUserReq(arg) {
  if (!(arg instanceof user_pb.CreateUserReq)) {
    throw new Error('Expected argument of type pb.CreateUserReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_CreateUserReq(buffer_arg) {
  return user_pb.CreateUserReq.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_pb_NewUserIdResp(arg) {
  if (!(arg instanceof user_pb.NewUserIdResp)) {
    throw new Error('Expected argument of type pb.NewUserIdResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_NewUserIdResp(buffer_arg) {
  return user_pb.NewUserIdResp.deserializeBinary(new Uint8Array(buffer_arg));
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
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
