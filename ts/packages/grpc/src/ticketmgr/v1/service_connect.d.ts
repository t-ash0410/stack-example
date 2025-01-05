// @generated by protoc-gen-connect-es v1.6.1
// @generated from file ticketmgr/v1/service.proto (package ticketmgr.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTicketRequest, CreateTicketResponse, DeleteTicketRequest, DeleteTicketResponse, UpdateTicketRequest, UpdateTicketResponse } from "./service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service ticketmgr.v1.TicketMgrService
 */
export declare const TicketMgrService: {
  readonly typeName: "ticketmgr.v1.TicketMgrService",
  readonly methods: {
    /**
     * @generated from rpc ticketmgr.v1.TicketMgrService.CreateTicket
     */
    readonly createTicket: {
      readonly name: "CreateTicket",
      readonly I: typeof CreateTicketRequest,
      readonly O: typeof CreateTicketResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ticketmgr.v1.TicketMgrService.UpdateTicket
     */
    readonly updateTicket: {
      readonly name: "UpdateTicket",
      readonly I: typeof UpdateTicketRequest,
      readonly O: typeof UpdateTicketResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ticketmgr.v1.TicketMgrService.DeleteTicket
     */
    readonly deleteTicket: {
      readonly name: "DeleteTicket",
      readonly I: typeof DeleteTicketRequest,
      readonly O: typeof DeleteTicketResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

