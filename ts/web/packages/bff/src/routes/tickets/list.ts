import { convertTicketToResponse } from '@bff/grpc/ticket'
import type { AuthNEnv } from '@bff/types'
import type { QueryTicketsResponse } from '@stack-example/grpc'
import type { Context } from 'hono'
import { createFactory } from 'hono/factory'
import { ResultAsync, ok } from 'neverthrow'

const handlers = createFactory<AuthNEnv>().createHandlers(async (c) => {
  const res = await listTickets(c).andThen(convertResponse)
  if (res.isErr()) {
    throw res.error
  }
  return c.json(res.value)
})

const listTickets = (ctx: Context<AuthNEnv>) => {
  const { activeUser, ticketQuerierServiceClient } = ctx.var
  return ResultAsync.fromThrowable(() =>
    ticketQuerierServiceClient.queryTickets({
      requestedBy: activeUser.userId,
    }),
  )()
}

const convertResponse = (res: QueryTicketsResponse) => {
  return ok({
    tickets: res.tickets.map(convertTicketToResponse),
  })
}

export { handlers as listHandlers }
