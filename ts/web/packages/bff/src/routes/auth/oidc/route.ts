import { oidcApp } from './app'
import { sessionHandlers } from './session'
import { slackHandlers } from './slack'

const oidcRoute = oidcApp
  .get('/session', ...sessionHandlers)
  .get('/slack', ...slackHandlers)

export { oidcRoute }
