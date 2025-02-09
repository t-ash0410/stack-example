import { convertTicketToResponse } from '@bff/grpc/ticket'
import { ResultAsync, ok } from 'neverthrow'
import { ticketsFactory } from './app'

const handlers = ticketsFactory.createHandlers(async (c) => {
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
