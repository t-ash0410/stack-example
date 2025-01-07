import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { listHandler } from './list'
import { ticketDetailRoute } from './{ticketId}'

const ticketsRoute = new Hono<AuthNEnv>()
  .get('/', listHandler)
  .route('/:ticketId', ticketDetailRoute)

export { ticketsRoute }
