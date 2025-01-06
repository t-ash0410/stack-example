import type { DefaultEnv } from '@bff/types'
import { Hono } from 'hono'
import { sessionHandler } from './session'
import { slackHandler, slackValidator } from './slack'

const oidcRoute = new Hono<DefaultEnv>()
  .get('/session', sessionHandler)
  .get('/slack', slackValidator, slackHandler)

export { oidcRoute }
