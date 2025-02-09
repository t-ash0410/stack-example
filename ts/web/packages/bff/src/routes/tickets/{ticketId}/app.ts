import type { AuthNEnv } from '@bff/types'
import { zValidator } from '@hono/zod-validator'
import type { Ticket } from '@stack-example/grpc'
import { createFactory } from 'hono/factory'
import { HTTPException } from 'hono/http-exception'
import { ResultAsync, err, ok } from 'neverthrow'
import { z } from 'zod'

const ticketDetailFactory = createFactory<
  AuthNEnv & {
    Variables: {
      ticket: Ticket
    }
  }
>()

const ticketDetailApp = ticketDetailFactory.createApp()

const ticketDetailParamValidator = ticketDetailFactory.createMiddleware(
  zValidator(
    'param',
    z.object({
      ticketId: z.string(),
    }),
    async (r, c) => {
      const { ticketId } = r.data
      const res = await ResultAsync.fromThrowable(() =>
        c.var.ticketQuerierServiceClient.getTicketById({
          ticketId,
        }),
      )().andThen((res) =>
        res.ticket?.createdBy === c.var.activeUser.userId
          ? ok(res.ticket)
          : err(
              new HTTPException(403, {
                message: 'User does not match',
              }),
            ),
      )
      if (res.isErr()) {
        throw res.error
      }
      c.set('ticket', res.value)
    },
  ),
)

export { ticketDetailFactory, ticketDetailApp, ticketDetailParamValidator }
