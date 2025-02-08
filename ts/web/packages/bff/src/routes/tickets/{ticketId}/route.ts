import { ticketDetailApp } from './app'
import { deleteHandlers } from './delete'
import { getHandlers } from './get'
import { updateHandlers } from './update'

const ticketDetailRoute = ticketDetailApp
  .get('/', ...getHandlers)
  .put('/', ...updateHandlers)
  .delete('/', ...deleteHandlers)

export { ticketDetailRoute }
