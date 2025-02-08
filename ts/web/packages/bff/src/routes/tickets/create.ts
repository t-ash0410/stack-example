import type { AuthNEnv, ValidatorSchema } from '@bff/types'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import type { Context } from 'hono'
import { ResultAsync } from 'neverthrow'
import { z } from 'zod'

const validator = zValidator(
  'json',
  z.object({
    title: z.string(),
    description: z.string(),
    deadline: z.string().transform((s) => new Date(s)),
  }),
)

type CreateTicketContext = Context<
  AuthNEnv,
  '',
  ValidatorSchema<typeof validator>
>

const handler = async (c: CreateTicketContext) => {
  const res = await createTicket(c)
  if (res.isErr()) {
    throw res.error
  }
  return c.json({
    ticketId: res.value.ticketId,
  })
}

const createTicket = (ctx: CreateTicketContext) => {
  const { activeUser, ticketMgrServiceClient } = ctx.var
  const { title, description, deadline } = ctx.req.valid('json')
  return ResultAsync.fromThrowable(() =>
    ticketMgrServiceClient.createTicket({
      requestedBy: activeUser.userId,
      title,
      description,
      deadline: timestampFromDate(deadline),
    }),
  )()
}

export { handler as createHandler, validator as createValidator }
