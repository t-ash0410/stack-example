// @generated by protoc-gen-es v2.2.3
// @generated from file ticketquerier/v1/service.proto (package ticketquerier.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { Timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file ticketquerier/v1/service.proto.
 */
export declare const file_ticketquerier_v1_service: GenFile;

/**
 * @generated from message ticketquerier.v1.Ticket
 */
export declare type Ticket = Message<"ticketquerier.v1.Ticket"> & {
  /**
   * @generated from field: string ticket_id = 1;
   */
  ticketId: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 2;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 3;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: string created_by = 4;
   */
  createdBy: string;

  /**
   * @generated from field: string title = 5;
   */
  title: string;

  /**
   * @generated from field: string description = 6;
   */
  description: string;

  /**
   * @generated from field: google.protobuf.Timestamp deadline = 7;
   */
  deadline?: Timestamp;
};

/**
 * Describes the message ticketquerier.v1.Ticket.
 * Use `create(TicketSchema)` to create a new message.
 */
export declare const TicketSchema: GenMessage<Ticket>;

/**
 * @generated from message ticketquerier.v1.QueryTicketsRequest
 */
export declare type QueryTicketsRequest = Message<"ticketquerier.v1.QueryTicketsRequest"> & {
  /**
   * @generated from field: string requested_by = 1;
   */
  requestedBy: string;
};

/**
 * Describes the message ticketquerier.v1.QueryTicketsRequest.
 * Use `create(QueryTicketsRequestSchema)` to create a new message.
 */
export declare const QueryTicketsRequestSchema: GenMessage<QueryTicketsRequest>;

/**
 * @generated from message ticketquerier.v1.QueryTicketsResponse
 */
export declare type QueryTicketsResponse = Message<"ticketquerier.v1.QueryTicketsResponse"> & {
  /**
   * @generated from field: repeated ticketquerier.v1.Ticket tickets = 1;
   */
  tickets: Ticket[];
};

/**
 * Describes the message ticketquerier.v1.QueryTicketsResponse.
 * Use `create(QueryTicketsResponseSchema)` to create a new message.
 */
export declare const QueryTicketsResponseSchema: GenMessage<QueryTicketsResponse>;

/**
 * @generated from message ticketquerier.v1.GetTicketByIdRequest
 */
export declare type GetTicketByIdRequest = Message<"ticketquerier.v1.GetTicketByIdRequest"> & {
  /**
   * @generated from field: string ticket_id = 1;
   */
  ticketId: string;
};

/**
 * Describes the message ticketquerier.v1.GetTicketByIdRequest.
 * Use `create(GetTicketByIdRequestSchema)` to create a new message.
 */
export declare const GetTicketByIdRequestSchema: GenMessage<GetTicketByIdRequest>;

/**
 * @generated from message ticketquerier.v1.GetTicketByIdResponse
 */
export declare type GetTicketByIdResponse = Message<"ticketquerier.v1.GetTicketByIdResponse"> & {
  /**
   * @generated from field: ticketquerier.v1.Ticket ticket = 1;
   */
  ticket?: Ticket;
};

/**
 * Describes the message ticketquerier.v1.GetTicketByIdResponse.
 * Use `create(GetTicketByIdResponseSchema)` to create a new message.
 */
export declare const GetTicketByIdResponseSchema: GenMessage<GetTicketByIdResponse>;

/**
 * @generated from service ticketquerier.v1.TicketQuerierService
 */
export declare const TicketQuerierService: GenService<{
  /**
   * @generated from rpc ticketquerier.v1.TicketQuerierService.QueryTickets
   */
  queryTickets: {
    methodKind: "unary";
    input: typeof QueryTicketsRequestSchema;
    output: typeof QueryTicketsResponseSchema;
  },
  /**
   * @generated from rpc ticketquerier.v1.TicketQuerierService.GetTicketById
   */
  getTicketById: {
    methodKind: "unary";
    input: typeof GetTicketByIdRequestSchema;
    output: typeof GetTicketByIdResponseSchema;
  },
}>;

