import { convertTicketToResponse } from '@bff/grpc/ticket'
import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'
import { ResultAsync, ok } from 'neverthrow'

const handlers = createFactory<AuthNEnv>().createHandlers(async (c) => {
  const { activeUser, ticketQuerierServiceClient } = c.var
  const res = await ResultAsync.fromThrowable(() =>
    ticketQuerierServiceClient.queryTickets({
      requestedBy: activeUser.userId,
    }),
  )().andThen((res) =>
    ok({
      tickets: res.tickets.map(convertTicketToResponse),
    }),
  )
  if (res.isErr()) {
    throw res.error
  }
  return c.json(res.value)
})

export { handlers as listHandlers }
