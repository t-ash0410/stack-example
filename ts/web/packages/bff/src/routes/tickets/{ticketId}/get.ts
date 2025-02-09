import { convertTicketToResponse } from '@bff/grpc/ticket'
import { ticketDetailFactory, ticketDetailParamValidator } from './app'

const handlers = ticketDetailFactory.createHandlers(
  ticketDetailParamValidator,
  async (c) => {
    return c.json(convertTicketToResponse(c.var.ticket))
  },
)

export { handlers as getHandlers }
