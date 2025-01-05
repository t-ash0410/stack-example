import type { DefaultEnv } from '@bff/types'
import { Hono } from 'hono'
import { oidcRoute } from './oidc'

const authRoute = new Hono<DefaultEnv>().route('/oidc', oidcRoute)

export { authRoute }
