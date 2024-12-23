// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var hls_pb = require('./hls_pb.js');

function serialize_pb_NewHlsStreamReq(arg) {
  if (!(arg instanceof hls_pb.NewHlsStreamReq)) {
    throw new Error('Expected argument of type pb.NewHlsStreamReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_NewHlsStreamReq(buffer_arg) {
  return hls_pb.NewHlsStreamReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_NewHlsStreamResp(arg) {
  if (!(arg instanceof hls_pb.NewHlsStreamResp)) {
    throw new Error('Expected argument of type pb.NewHlsStreamResp');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_NewHlsStreamResp(buffer_arg) {
  return hls_pb.NewHlsStreamResp.deserializeBinary(new Uint8Array(buffer_arg));
}


var HlsServiceService = exports.HlsServiceService = {
  newHlsStream: {
    path: '/pb.HlsService/NewHlsStream',
    requestStream: false,
    responseStream: false,
    requestType: hls_pb.NewHlsStreamReq,
    responseType: hls_pb.NewHlsStreamResp,
    requestSerialize: serialize_pb_NewHlsStreamReq,
    requestDeserialize: deserialize_pb_NewHlsStreamReq,
    responseSerialize: serialize_pb_NewHlsStreamResp,
    responseDeserialize: deserialize_pb_NewHlsStreamResp,
  },
};

exports.HlsServiceClient = grpc.makeGenericClientConstructor(HlsServiceService);
