import { ticketsApp } from './app'
import { createHandlers } from './create'
import { listHandlers } from './list'
import { ticketDetailRoute } from './{ticketId}'

const ticketsRoute = ticketsApp
  .get('/', ...listHandlers)
  .post('/', ...createHandlers)
  .route('/:ticketId', ticketDetailRoute)

export { ticketsRoute }
