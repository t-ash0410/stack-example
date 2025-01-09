import type { AuthNEnv } from '@bff/types'
import { Hono } from 'hono'
import { deleteHandler } from './delete'
import { getHandler } from './get'

const sessionRoute = new Hono<AuthNEnv>()
  .get('/', getHandler)
  .delete('/', deleteHandler)

export { sessionRoute }
