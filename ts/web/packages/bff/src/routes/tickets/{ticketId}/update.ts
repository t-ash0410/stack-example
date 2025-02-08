import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import { ResultAsync } from 'neverthrow'
import { z } from 'zod'
import { ticketDetailFactory, ticketDetailParamValidator } from './app'

const validator = ticketDetailFactory.createMiddleware(
  zValidator(
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
  ),
)

const handlers = ticketDetailFactory.createHandlers(
  ticketDetailParamValidator,
  validator,
  async (c) => {
    const { activeUser, ticketMgrServiceClient } = c.var
    const { ticketId } = c.req.valid('param')
    const { title, description, deadline } = c.req.valid('json')

    const res = await ResultAsync.fromThrowable(() =>
      ticketMgrServiceClient.updateTicket({
        ticketId: ticketId,
        requestedBy: activeUser.userId,
        title,
        description,
        deadline: deadline ? timestampFromDate(deadline) : undefined,
      }),
    )()
    if (res.isErr()) {
      throw res.error
    }
    return c.json({})
  },
)

export { handlers as updateHandlers }
