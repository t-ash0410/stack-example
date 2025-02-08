import type { AuthNEnv, ValidatorSchema } from '@bff/types'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import type { Context } from 'hono'
import { ResultAsync } from 'neverthrow'
import { z } from 'zod'
import type { ticketDetailParamValidator } from './middleware'

const validator = zValidator(
  'json',
  z.object({
    title: z.string().optional(),
    description: z.string().optional(),
    deadline: z
      .string()
      .optional()
      .transform((s) => (s ? new Date(s) : undefined)),
  }),
  (r, c) => {
    if (r.data.deadline && Number.isNaN(r.data.deadline.getTime())) {
      return c.json(
        {
          message: 'invalid date',
        },
        400,
      )
    }
  },
)

type UpdateTicketContext = Context<
  AuthNEnv,
  '',
  ValidatorSchema<typeof ticketDetailParamValidator> &
    ValidatorSchema<typeof validator>
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
      deadline: deadline ? timestampFromDate(deadline) : undefined,
    }),
  )()
}

export { handler as updateHandler, validator as updateValidator }
