import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { getHandler } from './get'
import { ticketDetailAuthZ, ticketDetailParamValidator } from './middleware'

const ticketDetailRoute = new Hono<AuthNEnv>()
  .use(ticketDetailParamValidator)
  .use(ticketDetailAuthZ)
  .get('/', getHandler)

export { ticketDetailRoute }
