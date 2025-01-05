import type { Env, MiddlewareHandler } from 'hono'
import type { Logger } from 'pino'
import type {
  AccountMgrServiceClient,
  TicketMgrServiceClient,
  TicketQuerierServiceClient,
} from './grpc'

export type DefaultEnv = {
  Variables: {
    logger: Logger
    accountMgrServiceClient: AccountMgrServiceClient
    ticketMgrServiceClient: TicketMgrServiceClient
    ticketQuerierServiceClient: TicketQuerierServiceClient
  }
}

export type ValidatorInput<T> = T extends MiddlewareHandler<
  Env,
  string,
  infer R
>
  ? R
  : never
