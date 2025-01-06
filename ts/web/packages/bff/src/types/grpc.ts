import type { Client } from '@connectrpc/connect'
import type {
  AccountMgrService,
  TicketMgrService,
  TicketQuerierService,
} from '@stack-example/grpc'

type AccountMgrServiceClient = Client<typeof AccountMgrService>
type TicketMgrServiceClient = Client<typeof TicketMgrService>
type TicketQuerierServiceClient = Client<typeof TicketQuerierService>

export type {
  AccountMgrServiceClient,
  TicketMgrServiceClient,
  TicketQuerierServiceClient,
}
