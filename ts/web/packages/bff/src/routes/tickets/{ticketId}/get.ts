import { HTTPException } from 'hono/http-exception'
import { ResultAsync, err, ok } from 'neverthrow'
import { ticketDetailFactory, ticketDetailParamValidator } from './app'

const handlers = ticketDetailFactory.createHandlers(
  ticketDetailParamValidator,
  async (c) => {
    const { ticketQuerierServiceClient } = c.var
    const { ticketId } = c.req.valid('param')

    const res = await ResultAsync.fromThrowable(() =>
      ticketQuerierServiceClient.getTicketById({
        ticketId: ticketId,
      }),
    )().andThen((res) => (res.ticket ? ok(res) : err(new HTTPException(404))))
    if (res.isErr()) {
      throw res.error
    }
    return c.json(res.value)
  },
)

export { handlers as getHandlers }
