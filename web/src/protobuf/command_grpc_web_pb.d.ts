import * as grpcWeb from 'grpc-web';

import * as protobuf_command_pb from '../protobuf/command_pb';


export class RoomServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listRoom(
    request: protobuf_command_pb.ListRoomRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: protobuf_command_pb.ListRoomResponse) => void
  ): grpcWeb.ClientReadableStream<protobuf_command_pb.ListRoomResponse>;

  createRoom(
    request: protobuf_command_pb.CreateRoomRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: protobuf_command_pb.CreateRoomResponse) => void
  ): grpcWeb.ClientReadableStream<protobuf_command_pb.CreateRoomResponse>;

  deleteRoom(
    request: protobuf_command_pb.DeleteRoomRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: protobuf_command_pb.DeleteRoomResponse) => void
  ): grpcWeb.ClientReadableStream<protobuf_command_pb.DeleteRoomResponse>;

}

export class RoomServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listRoom(
    request: protobuf_command_pb.ListRoomRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<protobuf_command_pb.ListRoomResponse>;

  createRoom(
    request: protobuf_command_pb.CreateRoomRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<protobuf_command_pb.CreateRoomResponse>;

  deleteRoom(
    request: protobuf_command_pb.DeleteRoomRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<protobuf_command_pb.DeleteRoomResponse>;

}

