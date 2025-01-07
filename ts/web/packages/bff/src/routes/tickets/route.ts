import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { listHandler } from './list'

const ticketsRoute = new Hono<AuthNEnv>().get('/', listHandler)

export { ticketsRoute }
