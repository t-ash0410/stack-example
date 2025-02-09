import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import { ResultAsync } from 'neverthrow'
import { z } from 'zod'
import { ticketsFactory } from './app'

const validator = ticketsFactory.createMiddleware(
  zValidator(
    'json',
    z.object({
      title: z.string(),
      description: z.string(),
      deadline: z.string().transform((s) => new Date(s)),
    }),
  ),
)

const handlers = ticketsFactory.createHandlers(validator, async (c) => {
  const { activeUser, ticketMgrServiceClient } = c.var
  const { title, description, deadline } = c.req.valid('json')
  const res = await ResultAsync.fromThrowable(() =>
    ticketMgrServiceClient.createTicket({
      requestedBy: activeUser.userId,
      title,
      description,
      deadline: timestampFromDate(deadline),
    }),
  )()
  if (res.isErr()) {
    throw res.error
  }
  return c.json({
    ticketId: res.value.ticketId,
  })
})

export { handlers as createHandlers }
