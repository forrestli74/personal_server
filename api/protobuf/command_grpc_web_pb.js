/**
 * @fileoverview gRPC-Web generated client stub for tmp
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.tmp = require('./command_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tmp.RoomServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tmp.RoomServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tmp.ListRoomRequest,
 *   !proto.tmp.ListRoomResponse>}
 */
const methodDescriptor_RoomService_ListRoom = new grpc.web.MethodDescriptor(
  '/tmp.RoomService/ListRoom',
  grpc.web.MethodType.UNARY,
  proto.tmp.ListRoomRequest,
  proto.tmp.ListRoomResponse,
  /**
   * @param {!proto.tmp.ListRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.ListRoomResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tmp.ListRoomRequest,
 *   !proto.tmp.ListRoomResponse>}
 */
const methodInfo_RoomService_ListRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tmp.ListRoomResponse,
  /**
   * @param {!proto.tmp.ListRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.ListRoomResponse.deserializeBinary
);


/**
 * @param {!proto.tmp.ListRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tmp.ListRoomResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tmp.ListRoomResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tmp.RoomServiceClient.prototype.listRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tmp.RoomService/ListRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_ListRoom,
      callback);
};


/**
 * @param {!proto.tmp.ListRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tmp.ListRoomResponse>}
 *     Promise that resolves to the response
 */
proto.tmp.RoomServicePromiseClient.prototype.listRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tmp.RoomService/ListRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_ListRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tmp.CreateRoomRequest,
 *   !proto.tmp.CreateRoomResponse>}
 */
const methodDescriptor_RoomService_CreateRoom = new grpc.web.MethodDescriptor(
  '/tmp.RoomService/CreateRoom',
  grpc.web.MethodType.UNARY,
  proto.tmp.CreateRoomRequest,
  proto.tmp.CreateRoomResponse,
  /**
   * @param {!proto.tmp.CreateRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.CreateRoomResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tmp.CreateRoomRequest,
 *   !proto.tmp.CreateRoomResponse>}
 */
const methodInfo_RoomService_CreateRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tmp.CreateRoomResponse,
  /**
   * @param {!proto.tmp.CreateRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.CreateRoomResponse.deserializeBinary
);


/**
 * @param {!proto.tmp.CreateRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tmp.CreateRoomResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tmp.CreateRoomResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tmp.RoomServiceClient.prototype.createRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tmp.RoomService/CreateRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_CreateRoom,
      callback);
};


/**
 * @param {!proto.tmp.CreateRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tmp.CreateRoomResponse>}
 *     Promise that resolves to the response
 */
proto.tmp.RoomServicePromiseClient.prototype.createRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tmp.RoomService/CreateRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_CreateRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tmp.DeleteRoomRequest,
 *   !proto.tmp.DeleteRoomResponse>}
 */
const methodDescriptor_RoomService_DeleteRoom = new grpc.web.MethodDescriptor(
  '/tmp.RoomService/DeleteRoom',
  grpc.web.MethodType.UNARY,
  proto.tmp.DeleteRoomRequest,
  proto.tmp.DeleteRoomResponse,
  /**
   * @param {!proto.tmp.DeleteRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.DeleteRoomResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tmp.DeleteRoomRequest,
 *   !proto.tmp.DeleteRoomResponse>}
 */
const methodInfo_RoomService_DeleteRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tmp.DeleteRoomResponse,
  /**
   * @param {!proto.tmp.DeleteRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tmp.DeleteRoomResponse.deserializeBinary
);


/**
 * @param {!proto.tmp.DeleteRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tmp.DeleteRoomResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tmp.DeleteRoomResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tmp.RoomServiceClient.prototype.deleteRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tmp.RoomService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_DeleteRoom,
      callback);
};


/**
 * @param {!proto.tmp.DeleteRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tmp.DeleteRoomResponse>}
 *     Promise that resolves to the response
 */
proto.tmp.RoomServicePromiseClient.prototype.deleteRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tmp.RoomService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_RoomService_DeleteRoom);
};


module.exports = proto.tmp;

