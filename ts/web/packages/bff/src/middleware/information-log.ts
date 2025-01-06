import type { DefaultEnv } from '@bff/types/hono'
import { createMiddleware } from 'hono/factory'

export const informationLog = createMiddleware<DefaultEnv>(async (c, next) => {
  c.var.logger.info({
    message: `Request received ${c.req.url}`,
    body: await c.req.text(),
  })
  await next()
})
