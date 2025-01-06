// @generated by protoc-gen-es v2.2.3
// @generated from file ticketquerier/v1/service.proto (package ticketquerier.v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file ticketquerier/v1/service.proto.
 */
export const file_ticketquerier_v1_service = /*@__PURE__*/
  fileDesc("Ch50aWNrZXRxdWVyaWVyL3YxL3NlcnZpY2UucHJvdG8SEHRpY2tldHF1ZXJpZXIudjEi4QEKBlRpY2tldBIRCgl0aWNrZXRfaWQYASABKAkSLgoKY3JlYXRlZF9hdBgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASLgoKdXBkYXRlZF9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASEgoKY3JlYXRlZF9ieRgEIAEoCRINCgV0aXRsZRgFIAEoCRITCgtkZXNjcmlwdGlvbhgGIAEoCRIsCghkZWFkbGluZRgHIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXAiKwoTUXVlcnlUaWNrZXRzUmVxdWVzdBIUCgxyZXF1ZXN0ZWRfYnkYASABKAkiQQoUUXVlcnlUaWNrZXRzUmVzcG9uc2USKQoHdGlja2V0cxgBIAMoCzIYLnRpY2tldHF1ZXJpZXIudjEuVGlja2V0IikKFEdldFRpY2tldEJ5SWRSZXF1ZXN0EhEKCXRpY2tldF9pZBgBIAEoCSJBChVHZXRUaWNrZXRCeUlkUmVzcG9uc2USKAoGdGlja2V0GAEgASgLMhgudGlja2V0cXVlcmllci52MS5UaWNrZXQy2wEKFFRpY2tldFF1ZXJpZXJTZXJ2aWNlEl8KDFF1ZXJ5VGlja2V0cxIlLnRpY2tldHF1ZXJpZXIudjEuUXVlcnlUaWNrZXRzUmVxdWVzdBomLnRpY2tldHF1ZXJpZXIudjEuUXVlcnlUaWNrZXRzUmVzcG9uc2UiABJiCg1HZXRUaWNrZXRCeUlkEiYudGlja2V0cXVlcmllci52MS5HZXRUaWNrZXRCeUlkUmVxdWVzdBonLnRpY2tldHF1ZXJpZXIudjEuR2V0VGlja2V0QnlJZFJlc3BvbnNlIgBC0QEKFGNvbS50aWNrZXRxdWVyaWVyLnYxQgxTZXJ2aWNlUHJvdG9QAVpKZ2l0aHViLmNvbS90LWFzaDA0MTAvc3RhY2stZXhhbXBsZS9nby9hcGkvdGlja2V0cXVlcmllci92MTt0aWNrZXRxdWVyaWVydjGiAgNUWFiqAhBUaWNrZXRxdWVyaWVyLlYxygIQVGlja2V0cXVlcmllclxWMeICHFRpY2tldHF1ZXJpZXJcVjFcR1BCTWV0YWRhdGHqAhFUaWNrZXRxdWVyaWVyOjpWMWIGcHJvdG8z", [file_google_protobuf_timestamp]);

/**
 * Describes the message ticketquerier.v1.Ticket.
 * Use `create(TicketSchema)` to create a new message.
 */
export const TicketSchema = /*@__PURE__*/
  messageDesc(file_ticketquerier_v1_service, 0);

/**
 * Describes the message ticketquerier.v1.QueryTicketsRequest.
 * Use `create(QueryTicketsRequestSchema)` to create a new message.
 */
export const QueryTicketsRequestSchema = /*@__PURE__*/
  messageDesc(file_ticketquerier_v1_service, 1);

/**
 * Describes the message ticketquerier.v1.QueryTicketsResponse.
 * Use `create(QueryTicketsResponseSchema)` to create a new message.
 */
export const QueryTicketsResponseSchema = /*@__PURE__*/
  messageDesc(file_ticketquerier_v1_service, 2);

/**
 * Describes the message ticketquerier.v1.GetTicketByIdRequest.
 * Use `create(GetTicketByIdRequestSchema)` to create a new message.
 */
export const GetTicketByIdRequestSchema = /*@__PURE__*/
  messageDesc(file_ticketquerier_v1_service, 3);

/**
 * Describes the message ticketquerier.v1.GetTicketByIdResponse.
 * Use `create(GetTicketByIdResponseSchema)` to create a new message.
 */
export const GetTicketByIdResponseSchema = /*@__PURE__*/
  messageDesc(file_ticketquerier_v1_service, 4);

/**
 * @generated from service ticketquerier.v1.TicketQuerierService
 */
export const TicketQuerierService = /*@__PURE__*/
  serviceDesc(file_ticketquerier_v1_service, 0);
