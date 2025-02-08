import { convertTicketToResponse } from '@bff/grpc/ticket'
import type { AuthNEnv, ValidatorSchema } from '@bff/types'
import type { GetTicketByIdResponse } from '@stack-example/grpc'
import type { Context } from 'hono'
import { HTTPException } from 'hono/http-exception'
import { ResultAsync, err, ok } from 'neverthrow'
import type { ticketDetailParamValidator } from './middleware'

type GetTicketContext = Context<
  AuthNEnv,
  '',
  ValidatorSchema<typeof ticketDetailParamValidator>
>

const handler = async (c: GetTicketContext) => {
  const res = await getTicket(c).andThen(convertResponse)
  if (res.isErr()) {
    throw res.error
  }
  return c.json(res.value)
}

const getTicket = (ctx: GetTicketContext) => {
  const { ticketQuerierServiceClient } = ctx.var
  const { ticketId } = ctx.req.valid('param')
  return ResultAsync.fromThrowable(() =>
    ticketQuerierServiceClient.getTicketById({
      ticketId: ticketId,
    }),
  )()
}

const convertResponse = (res: GetTicketByIdResponse) => {
  const t = res.ticket
  if (!t) {
    return err(new HTTPException(404))
  }
  return ok(convertTicketToResponse(t))
}

export { handler as getHandler }
