import type {} from '@bff/types'
import { ResultAsync } from 'neverthrow'
import { ticketDetailFactory, ticketDetailParamValidator } from './app'

const handlers = ticketDetailFactory.createHandlers(
  ticketDetailParamValidator,
  async (c) => {
    const { ticketMgrServiceClient } = c.var
    const { ticketId } = c.req.valid('param')
    const res = await ResultAsync.fromThrowable(() =>
      ticketMgrServiceClient.deleteTicket({
        ticketId,
      }),
    )()
    if (res.isErr()) {
      throw res.error
    }
    return c.json({})
  },
)

export { handlers as deleteHandlers }
