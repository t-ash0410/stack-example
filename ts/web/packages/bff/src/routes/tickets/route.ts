import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'
import { createHandlers } from './create'
import { listHandlers } from './list'
import { ticketDetailRoute } from './{ticketId}'

const ticketsRoute = createFactory<AuthNEnv>()
  .createApp()
  .get('/', ...listHandlers)
  .post('/', ...createHandlers)
  .route('/:ticketId', ticketDetailRoute)

export { ticketsRoute }
