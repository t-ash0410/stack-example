import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'
import { deleteHandlers } from './delete'
import { getHandlers } from './get'

const sessionRoute = createFactory<AuthNEnv>()
  .createApp()
  .get('/', ...getHandlers)
  .delete('/', ...deleteHandlers)

export { sessionRoute }
