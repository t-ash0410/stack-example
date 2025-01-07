import type { AuthNEnv, ValidatorInput } from '@bff/types'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import type { Context } from 'hono'
import { ResultAsync } from 'neverthrow'
import { z } from 'zod'
import type { ticketDetailParamValidator } from './middleware'

const validator = zValidator(
  'json',
  z.object({
    title: z.string(),
    description: z.string(),
    deadline: z.string().transform((s) => new Date(s)),
  }),
)

type UpdateTicketContext = Context<
  AuthNEnv,
  '',
  ValidatorInput<typeof ticketDetailParamValidator> &
    ValidatorInput<typeof validator>
>

const handler = async (c: UpdateTicketContext) => {
  const res = await updateTicket(c)
  if (res.isErr()) {
    throw res.error
  }
  return c.json({})
}

const updateTicket = (ctx: UpdateTicketContext) => {
  const { activeUser, ticketMgrServiceClient } = ctx.var
  const { ticketId } = ctx.req.valid('param')
  const { title, description, deadline } = ctx.req.valid('json')
  return ResultAsync.fromThrowable(() =>
    ticketMgrServiceClient.updateTicket({
      ticketId: ticketId,
      requestedBy: activeUser.userId,
      title,
      description,
      deadline: timestampFromDate(deadline),
    }),
  )()
}

export { handler as updateHandler, validator as updateValidator }
