import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { createHandler, createValidator } from './create'
import { listHandler } from './list'
import { ticketDetailRoute } from './{ticketId}'

const ticketsRoute = new Hono<AuthNEnv>()
  .get('/', listHandler)
  .post('/', createValidator, createHandler)
  .route('/:ticketId', ticketDetailRoute)

export { ticketsRoute }
