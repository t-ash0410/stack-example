import { BFF_PORT, CORS_ORIGIN } from '@bff/env'
import { errorHandler } from '@bff/error'
import { informationLog, initLogger } from '@bff/log'
import { healthRoute } from '@bff/routes'
import { Hono } from 'hono'
import { except } from 'hono/combine'
import { cors } from 'hono/cors'
import { secureHeaders } from 'hono/secure-headers'

initLogger('info')

const app = new Hono()
app.use(
  cors({
    origin: CORS_ORIGIN,
    credentials: true,
  }),
  // setupDefaultContext,
  secureHeaders(),
)
app.use('*', except(['/health'], informationLog))
// app.use('*', except(['/health', '/auth/*'], authN))
app.onError(errorHandler)

const routes = app.route('/health', healthRoute)

// Run
export default {
  fetch: app.fetch,
  port: BFF_PORT,
}

export type App = typeof routes
