// @generated by protoc-gen-connect-es v1.5.0
// @generated from file accountmgr/v1/service.proto (package accountmgr.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SlackSSORequest, SlackSSOResponse } from "./service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service accountmgr.v1.AccountMgrService
 */
export declare const AccountMgrService: {
  readonly typeName: "accountmgr.v1.AccountMgrService",
  readonly methods: {
    /**
     * @generated from rpc accountmgr.v1.AccountMgrService.SlackSSO
     */
    readonly slackSSO: {
      readonly name: "SlackSSO",
      readonly I: typeof SlackSSORequest,
      readonly O: typeof SlackSSOResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

