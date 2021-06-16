import * as jspb from 'google-protobuf'



export class Commands extends jspb.Message {
  getCommandsList(): Array<Command>;
  setCommandsList(value: Array<Command>): Commands;
  clearCommandsList(): Commands;
  addCommands(value?: Command, index?: number): Command;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Commands.AsObject;
  static toObject(includeInstance: boolean, msg: Commands): Commands.AsObject;
  static serializeBinaryToWriter(message: Commands, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Commands;
  static deserializeBinaryFromReader(message: Commands, reader: jspb.BinaryReader): Commands;
}

export namespace Commands {
  export type AsObject = {
    commandsList: Array<Command.AsObject>,
  }
}

export class Command extends jspb.Message {
  getTickCommand(): TickCommand | undefined;
  setTickCommand(value?: TickCommand): Command;
  hasTickCommand(): boolean;
  clearTickCommand(): Command;

  getIdCommand(): IdCommand | undefined;
  setIdCommand(value?: IdCommand): Command;
  hasIdCommand(): boolean;
  clearIdCommand(): Command;

  getWriterCommand(): WriterCommand | undefined;
  setWriterCommand(value?: WriterCommand): Command;
  hasWriterCommand(): boolean;
  clearWriterCommand(): Command;

  getCommandCase(): Command.CommandCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Command.AsObject;
  static toObject(includeInstance: boolean, msg: Command): Command.AsObject;
  static serializeBinaryToWriter(message: Command, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Command;
  static deserializeBinaryFromReader(message: Command, reader: jspb.BinaryReader): Command;
}

export namespace Command {
  export type AsObject = {
    tickCommand?: TickCommand.AsObject,
    idCommand?: IdCommand.AsObject,
    writerCommand?: WriterCommand.AsObject,
  }

  export enum CommandCase { 
    COMMAND_NOT_SET = 0,
    TICK_COMMAND = 1,
    ID_COMMAND = 2,
    WRITER_COMMAND = 3,
  }
}

export class TickCommand extends jspb.Message {
  getRandomSeed(): Uint8Array | string;
  getRandomSeed_asU8(): Uint8Array;
  getRandomSeed_asB64(): string;
  setRandomSeed(value: Uint8Array | string): TickCommand;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TickCommand.AsObject;
  static toObject(includeInstance: boolean, msg: TickCommand): TickCommand.AsObject;
  static serializeBinaryToWriter(message: TickCommand, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TickCommand;
  static deserializeBinaryFromReader(message: TickCommand, reader: jspb.BinaryReader): TickCommand;
}

export namespace TickCommand {
  export type AsObject = {
    randomSeed: Uint8Array | string,
  }
}

export class IdCommand extends jspb.Message {
  getOldId(): string;
  setOldId(value: string): IdCommand;

  getNewId(): string;
  setNewId(value: string): IdCommand;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IdCommand.AsObject;
  static toObject(includeInstance: boolean, msg: IdCommand): IdCommand.AsObject;
  static serializeBinaryToWriter(message: IdCommand, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IdCommand;
  static deserializeBinaryFromReader(message: IdCommand, reader: jspb.BinaryReader): IdCommand;
}

export namespace IdCommand {
  export type AsObject = {
    oldId: string,
    newId: string,
  }
}

export class WriterCommand extends jspb.Message {
  getId(): string;
  setId(value: string): WriterCommand;

  getCommand(): Uint8Array | string;
  getCommand_asU8(): Uint8Array;
  getCommand_asB64(): string;
  setCommand(value: Uint8Array | string): WriterCommand;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WriterCommand.AsObject;
  static toObject(includeInstance: boolean, msg: WriterCommand): WriterCommand.AsObject;
  static serializeBinaryToWriter(message: WriterCommand, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WriterCommand;
  static deserializeBinaryFromReader(message: WriterCommand, reader: jspb.BinaryReader): WriterCommand;
}

export namespace WriterCommand {
  export type AsObject = {
    id: string,
    command: Uint8Array | string,
  }
}

export class StandAloneRoomServerSetting extends jspb.Message {
  getPort(): number;
  setPort(value: number): StandAloneRoomServerSetting;

  getRoomSetting(): RoomSetting | undefined;
  setRoomSetting(value?: RoomSetting): StandAloneRoomServerSetting;
  hasRoomSetting(): boolean;
  clearRoomSetting(): StandAloneRoomServerSetting;

  getBackupFile(): string;
  setBackupFile(value: string): StandAloneRoomServerSetting;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StandAloneRoomServerSetting.AsObject;
  static toObject(includeInstance: boolean, msg: StandAloneRoomServerSetting): StandAloneRoomServerSetting.AsObject;
  static serializeBinaryToWriter(message: StandAloneRoomServerSetting, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StandAloneRoomServerSetting;
  static deserializeBinaryFromReader(message: StandAloneRoomServerSetting, reader: jspb.BinaryReader): StandAloneRoomServerSetting;
}

export namespace StandAloneRoomServerSetting {
  export type AsObject = {
    port: number,
    roomSetting?: RoomSetting.AsObject,
    backupFile: string,
  }
}

export class CreateRoomRequest extends jspb.Message {
  getRoomId(): string;
  setRoomId(value: string): CreateRoomRequest;

  getRoomSetting(): RoomSetting | undefined;
  setRoomSetting(value?: RoomSetting): CreateRoomRequest;
  hasRoomSetting(): boolean;
  clearRoomSetting(): CreateRoomRequest;

  getShortDescription(): string;
  setShortDescription(value: string): CreateRoomRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoomRequest): CreateRoomRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoomRequest;
  static deserializeBinaryFromReader(message: CreateRoomRequest, reader: jspb.BinaryReader): CreateRoomRequest;
}

export namespace CreateRoomRequest {
  export type AsObject = {
    roomId: string,
    roomSetting?: RoomSetting.AsObject,
    shortDescription: string,
  }
}

export class CreateRoomResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoomResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoomResponse): CreateRoomResponse.AsObject;
  static serializeBinaryToWriter(message: CreateRoomResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoomResponse;
  static deserializeBinaryFromReader(message: CreateRoomResponse, reader: jspb.BinaryReader): CreateRoomResponse;
}

export namespace CreateRoomResponse {
  export type AsObject = {
  }
}

export class DeleteRoomRequest extends jspb.Message {
  getRoomId(): string;
  setRoomId(value: string): DeleteRoomRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoomRequest): DeleteRoomRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoomRequest;
  static deserializeBinaryFromReader(message: DeleteRoomRequest, reader: jspb.BinaryReader): DeleteRoomRequest;
}

export namespace DeleteRoomRequest {
  export type AsObject = {
    roomId: string,
  }
}

export class DeleteRoomResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoomResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoomResponse): DeleteRoomResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteRoomResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoomResponse;
  static deserializeBinaryFromReader(message: DeleteRoomResponse, reader: jspb.BinaryReader): DeleteRoomResponse;
}

export namespace DeleteRoomResponse {
  export type AsObject = {
  }
}

export class ListRoomRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomRequest): ListRoomRequest.AsObject;
  static serializeBinaryToWriter(message: ListRoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomRequest;
  static deserializeBinaryFromReader(message: ListRoomRequest, reader: jspb.BinaryReader): ListRoomRequest;
}

export namespace ListRoomRequest {
  export type AsObject = {
  }
}

export class ListRoomResponse extends jspb.Message {
  getRoomsList(): Array<RoomSummary>;
  setRoomsList(value: Array<RoomSummary>): ListRoomResponse;
  clearRoomsList(): ListRoomResponse;
  addRooms(value?: RoomSummary, index?: number): RoomSummary;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomResponse): ListRoomResponse.AsObject;
  static serializeBinaryToWriter(message: ListRoomResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomResponse;
  static deserializeBinaryFromReader(message: ListRoomResponse, reader: jspb.BinaryReader): ListRoomResponse;
}

export namespace ListRoomResponse {
  export type AsObject = {
    roomsList: Array<RoomSummary.AsObject>,
  }
}

export class RoomSummary extends jspb.Message {
  getId(): string;
  setId(value: string): RoomSummary;

  getSetting(): RoomSetting | undefined;
  setSetting(value?: RoomSetting): RoomSummary;
  hasSetting(): boolean;
  clearSetting(): RoomSummary;

  getShortDescription(): string;
  setShortDescription(value: string): RoomSummary;

  getHistorySummary(): RoomSummary.HistorySummary | undefined;
  setHistorySummary(value?: RoomSummary.HistorySummary): RoomSummary;
  hasHistorySummary(): boolean;
  clearHistorySummary(): RoomSummary;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomSummary.AsObject;
  static toObject(includeInstance: boolean, msg: RoomSummary): RoomSummary.AsObject;
  static serializeBinaryToWriter(message: RoomSummary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomSummary;
  static deserializeBinaryFromReader(message: RoomSummary, reader: jspb.BinaryReader): RoomSummary;
}

export namespace RoomSummary {
  export type AsObject = {
    id: string,
    setting?: RoomSetting.AsObject,
    shortDescription: string,
    historySummary?: RoomSummary.HistorySummary.AsObject,
  }

  export class HistorySummary extends jspb.Message {
    getActiveConnectionSize(): number;
    setActiveConnectionSize(value: number): HistorySummary;

    getCommandSize(): number;
    setCommandSize(value: number): HistorySummary;

    getCommandMemorySize(): number;
    setCommandMemorySize(value: number): HistorySummary;

    getTickSize(): number;
    setTickSize(value: number): HistorySummary;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HistorySummary.AsObject;
    static toObject(includeInstance: boolean, msg: HistorySummary): HistorySummary.AsObject;
    static serializeBinaryToWriter(message: HistorySummary, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HistorySummary;
    static deserializeBinaryFromReader(message: HistorySummary, reader: jspb.BinaryReader): HistorySummary;
  }

  export namespace HistorySummary {
    export type AsObject = {
      activeConnectionSize: number,
      commandSize: number,
      commandMemorySize: number,
      tickSize: number,
    }
  }

}

export class DebugRequest extends jspb.Message {
  getB(): Uint8Array | string;
  getB_asU8(): Uint8Array;
  getB_asB64(): string;
  setB(value: Uint8Array | string): DebugRequest;

  getS(): string;
  setS(value: string): DebugRequest;

  getI32(): number;
  setI32(value: number): DebugRequest;

  getI64(): number;
  setI64(value: number): DebugRequest;

  getOi32(): number;
  setOi32(value: number): DebugRequest;

  getOi64(): number;
  setOi64(value: number): DebugRequest;

  getOoCase(): DebugRequest.OoCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DebugRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DebugRequest): DebugRequest.AsObject;
  static serializeBinaryToWriter(message: DebugRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DebugRequest;
  static deserializeBinaryFromReader(message: DebugRequest, reader: jspb.BinaryReader): DebugRequest;
}

export namespace DebugRequest {
  export type AsObject = {
    b: Uint8Array | string,
    s: string,
    i32: number,
    i64: number,
    oi32: number,
    oi64: number,
  }

  export enum OoCase { 
    OO_NOT_SET = 0,
    OI32 = 5,
    OI64 = 6,
  }
}

export class DebugResponse extends jspb.Message {
  getB(): Uint8Array | string;
  getB_asU8(): Uint8Array;
  getB_asB64(): string;
  setB(value: Uint8Array | string): DebugResponse;

  getS(): string;
  setS(value: string): DebugResponse;

  getI32(): number;
  setI32(value: number): DebugResponse;

  getI64(): number;
  setI64(value: number): DebugResponse;

  getOi32(): number;
  setOi32(value: number): DebugResponse;

  getOi64(): number;
  setOi64(value: number): DebugResponse;

  getOoCase(): DebugResponse.OoCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DebugResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DebugResponse): DebugResponse.AsObject;
  static serializeBinaryToWriter(message: DebugResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DebugResponse;
  static deserializeBinaryFromReader(message: DebugResponse, reader: jspb.BinaryReader): DebugResponse;
}

export namespace DebugResponse {
  export type AsObject = {
    b: Uint8Array | string,
    s: string,
    i32: number,
    i64: number,
    oi32: number,
    oi64: number,
  }

  export enum OoCase { 
    OO_NOT_SET = 0,
    OI32 = 5,
    OI64 = 6,
  }
}

export class RoomSetting extends jspb.Message {
  getTick(): TickSetting | undefined;
  setTick(value?: TickSetting): RoomSetting;
  hasTick(): boolean;
  clearTick(): RoomSetting;

  getEndOfLife(): EndOfLifeSetting | undefined;
  setEndOfLife(value?: EndOfLifeSetting): RoomSetting;
  hasEndOfLife(): boolean;
  clearEndOfLife(): RoomSetting;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomSetting.AsObject;
  static toObject(includeInstance: boolean, msg: RoomSetting): RoomSetting.AsObject;
  static serializeBinaryToWriter(message: RoomSetting, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomSetting;
  static deserializeBinaryFromReader(message: RoomSetting, reader: jspb.BinaryReader): RoomSetting;
}

export namespace RoomSetting {
  export type AsObject = {
    tick?: TickSetting.AsObject,
    endOfLife?: EndOfLifeSetting.AsObject,
  }
}

export class TickSetting extends jspb.Message {
  getSize(): number;
  setSize(value: number): TickSetting;

  getFrequencyNanoseconds(): number;
  setFrequencyNanoseconds(value: number): TickSetting;

  getAlwaysActive(): boolean;
  setAlwaysActive(value: boolean): TickSetting;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TickSetting.AsObject;
  static toObject(includeInstance: boolean, msg: TickSetting): TickSetting.AsObject;
  static serializeBinaryToWriter(message: TickSetting, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TickSetting;
  static deserializeBinaryFromReader(message: TickSetting, reader: jspb.BinaryReader): TickSetting;
}

export namespace TickSetting {
  export type AsObject = {
    size: number,
    frequencyNanoseconds: number,
    alwaysActive: boolean,
  }
}

export class EndOfLifeSetting extends jspb.Message {
  getMaxDuration(): number;
  setMaxDuration(value: number): EndOfLifeSetting;

  getMaxDurationWhileIdle(): number;
  setMaxDurationWhileIdle(value: number): EndOfLifeSetting;

  getCloseWhenAllWriterDisconnected(): boolean;
  setCloseWhenAllWriterDisconnected(value: boolean): EndOfLifeSetting;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EndOfLifeSetting.AsObject;
  static toObject(includeInstance: boolean, msg: EndOfLifeSetting): EndOfLifeSetting.AsObject;
  static serializeBinaryToWriter(message: EndOfLifeSetting, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EndOfLifeSetting;
  static deserializeBinaryFromReader(message: EndOfLifeSetting, reader: jspb.BinaryReader): EndOfLifeSetting;
}

export namespace EndOfLifeSetting {
  export type AsObject = {
    maxDuration: number,
    maxDurationWhileIdle: number,
    closeWhenAllWriterDisconnected: boolean,
  }
}

