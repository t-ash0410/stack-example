import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { getHandler } from './get'

const sessionRoute = new Hono<AuthNEnv>().get('/', getHandler)

export { sessionRoute }
