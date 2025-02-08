import type { AuthNEnv } from '@bff/types'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { zValidator } from '@hono/zod-validator'
import { createFactory } from 'hono/factory'
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

const handlers = createFactory<AuthNEnv>().createHandlers(
  validator,
  async (c) => {
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
  },
)

export { handlers as createHandlers }
