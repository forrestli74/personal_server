/*
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
*/
import * as pbjs from "protobufjs";
import commandProto from './protobuf/command.proto';
import * as grpcjs from '@grpc/grpc-js';
import * as grpcweb from 'grpc-web';
import {serviceToDefinition} from './tmp';
import {methodToGrpc2, ProtobufMessage3} from './tmp2';
import * as commandGrpcWeb from './protobuf/command_grpc_web_pb';
import * as commandGrpc from './protobuf/command_pb';
console.log(commandProto);
console.log(commandGrpcWeb);
function tmp1() {
	const client = new commandGrpcWeb.RoomServiceClient('http://localhost:8000');
	client.listRoom(new commandGrpc.ListRoomRequest(), {}, (err: grpcweb.Error, response: commandGrpc.ListRoomResponse) => {
		console.log(response);
	});
}
tmp1();
/*
pbjs.load(commandProto).then(root => {
	const s = root.lookupService("RoomService");
	const d = serviceToDefinition(s);
	const s2 = grpcjs.makeGenericClientConstructor(d, s.name);
	debugger;
});
*/

pbjs.load(commandProto).then(root => {
	const s = root.lookupService("RoomService");
	const t = root.lookupType("ListRoomRequest");
	window['t'] = t;
	const client = methodToGrpc2(s.methods['ListRoom'], 'http://localhost:8000', 'tmp.RoomService');
	client.start();
	client.send(t.create() as ProtobufMessage3);
	client.finishSend();
	debugger;
});




/*
c.listRoom(r, undefined).then(r => {
	debugger;
})
*/
