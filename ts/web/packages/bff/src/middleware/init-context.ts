import {
  ACCOUNT_MGR_SERVICE_ENDPOINT,
  TICKET_MGR_SERVICE_ENDPOINT,
  TICKET_QUERIER_SERVICE_ENDPOINT,
} from '@bff/env'
import { getLogger } from '@bff/log'
import type { DefaultEnv } from '@bff/types/hono'
import { createClient } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'
import {
  AccountMgrService,
  TicketMgrService,
  TicketQuerierService,
} from '@stack-example/grpc'
import { createMiddleware } from 'hono/factory'

const accountMgrServiceClient = createClient(
  AccountMgrService,
  createGrpcTransport({
    baseUrl: ACCOUNT_MGR_SERVICE_ENDPOINT,
  }),
)

const ticketMgrServiceClient = createClient(
  TicketMgrService,
  createGrpcTransport({
    baseUrl: TICKET_MGR_SERVICE_ENDPOINT,
  }),
)

const ticketQuerierServiceClient = createClient(
  TicketQuerierService,
  createGrpcTransport({
    baseUrl: TICKET_QUERIER_SERVICE_ENDPOINT,
  }),
)

export const initContext = createMiddleware<DefaultEnv>(async (c, next) => {
  c.set('logger', getLogger())
  c.set('accountMgrServiceClient', accountMgrServiceClient)
  c.set('ticketMgrServiceClient', ticketMgrServiceClient)
  c.set('ticketQuerierServiceClient', ticketQuerierServiceClient)
  await next()
})
