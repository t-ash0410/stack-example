import { errorHandler } from '@bff/error'
import { getLogger } from '@bff/log'
import type { DefaultEnv } from '@bff/types/hono'
import { createClient, createRouterTransport } from '@connectrpc/connect'
import {
  AccountMgrService,
  TicketMgrService,
  TicketQuerierService,
} from '@stack-example/grpc'
import { Hono } from 'hono'

const mockAccountMgrServiceClient = createClient(
  AccountMgrService,
  createRouterTransport(({ service }) => {
    service(AccountMgrService, {})
  }),
)

const mockTicketMgrServiceClient = createClient(
  TicketMgrService,
  createRouterTransport(({ service }) => {
    service(TicketMgrService, {})
  }),
)

const mockTicketQuerierServiceClient = createClient(
  TicketQuerierService,
  createRouterTransport(({ service }) => {
    service(TicketQuerierService, {})
  }),
)

const initHonoApp = () =>
  new Hono<DefaultEnv>()
    .use(async (c, next) => {
      c.set('logger', getLogger())
      c.set('accountMgrServiceClient', mockAccountMgrServiceClient)
      c.set('ticketMgrServiceClient', mockTicketMgrServiceClient)
      c.set('ticketQuerierServiceClient', mockTicketQuerierServiceClient)
      await next()
    })
    .onError(errorHandler)

export {
  initHonoApp,
  mockAccountMgrServiceClient,
  mockTicketMgrServiceClient,
  mockTicketQuerierServiceClient,
}
