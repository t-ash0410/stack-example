// @generated by protoc-gen-es v1.10.0
// @generated from file todomgr/v1/service.proto (package todomgr.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message todomgr.v1.CreateTicketRequest
 */
export declare class CreateTicketRequest extends Message<CreateTicketRequest> {
  /**
   * @generated from field: string requested_by = 1;
   */
  requestedBy: string;

  /**
   * @generated from field: string title = 2;
   */
  title: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: google.protobuf.Timestamp deadline = 4;
   */
  deadline?: Timestamp;

  constructor(data?: PartialMessage<CreateTicketRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.CreateTicketRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTicketRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTicketRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTicketRequest;

  static equals(a: CreateTicketRequest | PlainMessage<CreateTicketRequest> | undefined, b: CreateTicketRequest | PlainMessage<CreateTicketRequest> | undefined): boolean;
}

/**
 * @generated from message todomgr.v1.CreateTicketResponse
 */
export declare class CreateTicketResponse extends Message<CreateTicketResponse> {
  /**
   * @generated from field: string ticket_id = 1;
   */
  ticketId: string;

  constructor(data?: PartialMessage<CreateTicketResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.CreateTicketResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTicketResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTicketResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTicketResponse;

  static equals(a: CreateTicketResponse | PlainMessage<CreateTicketResponse> | undefined, b: CreateTicketResponse | PlainMessage<CreateTicketResponse> | undefined): boolean;
}

/**
 * @generated from message todomgr.v1.UpdateTicketRequest
 */
export declare class UpdateTicketRequest extends Message<UpdateTicketRequest> {
  /**
   * @generated from field: string ticket_id = 1;
   */
  ticketId: string;

  /**
   * @generated from field: string requested_by = 2;
   */
  requestedBy: string;

  /**
   * @generated from field: optional string title = 3;
   */
  title?: string;

  /**
   * @generated from field: optional string description = 4;
   */
  description?: string;

  /**
   * @generated from field: google.protobuf.Timestamp deadline = 5;
   */
  deadline?: Timestamp;

  constructor(data?: PartialMessage<UpdateTicketRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.UpdateTicketRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTicketRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTicketRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTicketRequest;

  static equals(a: UpdateTicketRequest | PlainMessage<UpdateTicketRequest> | undefined, b: UpdateTicketRequest | PlainMessage<UpdateTicketRequest> | undefined): boolean;
}

/**
 * @generated from message todomgr.v1.UpdateTicketResponse
 */
export declare class UpdateTicketResponse extends Message<UpdateTicketResponse> {
  constructor(data?: PartialMessage<UpdateTicketResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.UpdateTicketResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTicketResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTicketResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTicketResponse;

  static equals(a: UpdateTicketResponse | PlainMessage<UpdateTicketResponse> | undefined, b: UpdateTicketResponse | PlainMessage<UpdateTicketResponse> | undefined): boolean;
}

/**
 * @generated from message todomgr.v1.DeleteTicketRequest
 */
export declare class DeleteTicketRequest extends Message<DeleteTicketRequest> {
  /**
   * @generated from field: string ticket_id = 1;
   */
  ticketId: string;

  constructor(data?: PartialMessage<DeleteTicketRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.DeleteTicketRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTicketRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTicketRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTicketRequest;

  static equals(a: DeleteTicketRequest | PlainMessage<DeleteTicketRequest> | undefined, b: DeleteTicketRequest | PlainMessage<DeleteTicketRequest> | undefined): boolean;
}

/**
 * @generated from message todomgr.v1.DeleteTicketResponse
 */
export declare class DeleteTicketResponse extends Message<DeleteTicketResponse> {
  constructor(data?: PartialMessage<DeleteTicketResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "todomgr.v1.DeleteTicketResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTicketResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTicketResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTicketResponse;

  static equals(a: DeleteTicketResponse | PlainMessage<DeleteTicketResponse> | undefined, b: DeleteTicketResponse | PlainMessage<DeleteTicketResponse> | undefined): boolean;
}

