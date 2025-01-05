import { getLogger } from '@bff/log'
import { createMiddleware } from 'hono/factory'

export const informationLog = createMiddleware(async (c, next) => {
  const log = getLogger()
  log.info({
    message: `Request received ${c.req.url}`,
    body: await c.req.text(),
  })
  await next()
})
