import { createMiddleware } from 'hono/factory'
import { getLogger } from './logger'

export const informationLog = createMiddleware(async (c, next) => {
  const log = getLogger()
  log.info({
    message: `Request received ${c.req.url}`,
    body: await c.req.text(),
  })
  await next()
})
