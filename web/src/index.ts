import * as grpc from "./protobuf/command_grpc_web_pb";
import * as proto from "./protobuf/command_pb";
import $ from 'jquery';
import tmp from "./test.template.ejs";
console.log(tmp({a:1}));
const c = new grpc.RoomServicePromiseClient("http://localhost:8000");
const r = new proto.ListRoomRequest();
$("#create-button").on("click",(_) => {
	const s: string = ($("#create-input").val()) as string;
	const r = new proto.CreateRoomRequest();
	c.createRoom(r, {a: "b"})
});



/*
c.listRoom(r, undefined).then(r => {
	debugger;
})
*/
