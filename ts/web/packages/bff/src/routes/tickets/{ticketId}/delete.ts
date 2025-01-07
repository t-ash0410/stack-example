import type { AuthNEnv, ValidatorInput } from '@bff/types'
import type { Context } from 'hono'
import { ResultAsync } from 'neverthrow'
import type { ticketDetailParamValidator } from './middleware'

type DeleteTicketContext = Context<
  AuthNEnv,
  '',
  ValidatorInput<typeof ticketDetailParamValidator>
>

const handler = async (c: DeleteTicketContext) => {
  const res = await deleteTicket(c)
  if (res.isErr()) {
    throw res.error
  }
  return c.json({})
}

const deleteTicket = (ctx: DeleteTicketContext) => {
  const { ticketMgrServiceClient } = ctx.var
  const { ticketId } = ctx.req.valid('param')
  return ResultAsync.fromThrowable(() =>
    ticketMgrServiceClient.deleteTicket({
      ticketId,
    }),
  )()
}

export { handler as deleteHandler }
