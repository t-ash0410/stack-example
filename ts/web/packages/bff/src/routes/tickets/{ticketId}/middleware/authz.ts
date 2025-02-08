import type { AuthNEnv, ValidatorSchema } from '@bff/types'
import type { Ticket } from '@stack-example/grpc'
import type { Context } from 'hono'
import { createMiddleware } from 'hono/factory'
import { HTTPException } from 'hono/http-exception'
import { ResultAsync, err, ok } from 'neverthrow'
import type { ticketDetailParamValidator } from './param-validator'

type TicketDetailContext = Context<
  AuthNEnv,
  '',
  ValidatorSchema<typeof ticketDetailParamValidator>
>

const ticketDetailAuthZ = createMiddleware(async (c, next) => {
  const res = await getTicket(c).andThen((res) => verify(c, res.ticket))
  if (res.isErr()) {
    throw res.error
  }
  await next()
})

const getTicket = (c: TicketDetailContext) =>
  ResultAsync.fromThrowable(() => {
    const { ticketId } = c.req.valid('param')
    const { ticketQuerierServiceClient } = c.var
    return ticketQuerierServiceClient.getTicketById({
      ticketId,
    })
  })()

const verify = (c: TicketDetailContext, ticket?: Ticket) => {
  const { activeUser } = c.var
  return ticket?.createdBy === activeUser.userId
    ? ok({})
    : err(
        new HTTPException(403, {
          message: 'User does not match',
        }),
      )
}

export { ticketDetailAuthZ }
