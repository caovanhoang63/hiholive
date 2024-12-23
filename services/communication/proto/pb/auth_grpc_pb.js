// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var auth_pb = require('./auth_pb.js');

function serialize_pb_IntrospectReq(arg) {
  if (!(arg instanceof auth_pb.IntrospectReq)) {
    throw new Error('Expected argument of type pb.IntrospectReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_IntrospectReq(buffer_arg) {
  return auth_pb.IntrospectReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_IntrospectResp(arg) {
  if (!(arg instanceof auth_pb.IntrospectResp)) {
    throw new Error('Expected argument of type pb.IntrospectResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_IntrospectResp(buffer_arg) {
  return auth_pb.IntrospectResp.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthServiceService = exports.AuthServiceService = {
  introspectToken: {
    path: '/pb.AuthService/IntrospectToken',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.IntrospectReq,
    responseType: auth_pb.IntrospectResp,
    requestSerialize: serialize_pb_IntrospectReq,
    requestDeserialize: deserialize_pb_IntrospectReq,
    responseSerialize: serialize_pb_IntrospectResp,
    responseDeserialize: deserialize_pb_IntrospectResp,
  },
};

exports.AuthServiceClient = grpc.makeGenericClientConstructor(AuthServiceService);
