// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stream_pb = require('./stream_pb.js');

function serialize_pb_FindStreamReq(arg) {
  if (!(arg instanceof stream_pb.FindStreamReq)) {
    throw new Error('Expected argument of type pb.FindStreamReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_FindStreamReq(buffer_arg) {
  return stream_pb.FindStreamReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_FindStreamResp(arg) {
  if (!(arg instanceof stream_pb.FindStreamResp)) {
    throw new Error('Expected argument of type pb.FindStreamResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_FindStreamResp(buffer_arg) {
  return stream_pb.FindStreamResp.deserializeBinary(new Uint8Array(buffer_arg));
}


var StreamServiceService = exports.StreamServiceService = {
  findStreamById: {
    path: '/pb.StreamService/FindStreamById',
    requestStream: false,
    responseStream: false,
    requestType: stream_pb.FindStreamReq,
    responseType: stream_pb.FindStreamResp,
    requestSerialize: serialize_pb_FindStreamReq,
    requestDeserialize: deserialize_pb_FindStreamReq,
    responseSerialize: serialize_pb_FindStreamResp,
    responseDeserialize: deserialize_pb_FindStreamResp,
  },
};

exports.StreamServiceClient = grpc.makeGenericClientConstructor(StreamServiceService);
