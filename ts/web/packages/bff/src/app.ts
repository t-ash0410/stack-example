import {
  BFF_PORT,
  CORS_ORIGIN,
  TLS_CERT_PATH,
  TLS_KEY_PATH,
  USE_HTTPS,
} from '@bff/env'
import { errorHandler } from '@bff/error'
import { initLogger } from '@bff/log'
import { authN, informationLog, initContext } from '@bff/middleware'
import { authRoute, healthRoute, sessionRoute, ticketsRoute } from '@bff/routes'
import { Hono } from 'hono'
import { except } from 'hono/combine'
import { cors } from 'hono/cors'
import { secureHeaders } from 'hono/secure-headers'

initLogger('info')

const app = new Hono()
  .use(
    cors({
      origin: CORS_ORIGIN,
      credentials: true,
    }),
    initContext,
    secureHeaders(),
  )
  .use('*', except(['/health', '/session'], informationLog))
  .use('*', except(['/health', '/auth/*'], authN))
  .onError(errorHandler)
  .route('/auth', authRoute)
  .route('/health', healthRoute)
  .route('/session', sessionRoute)
  .route('/tickets', ticketsRoute)

// Run
export default {
  fetch: app.fetch,
  port: BFF_PORT,
  tls: USE_HTTPS
    ? {
        key: Bun.file(TLS_KEY_PATH),
        cert: Bun.file(TLS_CERT_PATH),
        serverName: 'localhost',
      }
    : undefined,
}

export type App = typeof app
