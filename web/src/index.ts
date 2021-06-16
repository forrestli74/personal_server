import * as grpc from "./protobuf/command_grpc_web_pb";
import * as proto from "./protobuf/command_pb";
const tmp: any = {};
const c = new grpc.RoomServicePromiseClient("http://localhost:8000");
const r = new proto.ListRoomRequest();
c.listRoom(r, undefined).then(r => {
	debugger;
})

window["tmp"] = tmp;
