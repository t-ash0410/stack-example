import type { DefaultEnv } from '@bff/types'
import { createFactory } from 'hono/factory'
import { sessionHandlers } from './session'
import { slackHandlers } from './slack'

const oidcRoute = createFactory<DefaultEnv>()
  .createApp()
  .get('/session', ...sessionHandlers)
  .get('/slack', ...slackHandlers)

export { oidcRoute }
