import { verifyJWTCookie } from '@bff/jwt'
import type { AuthNEnv } from '@bff/types'
import { createMiddleware } from 'hono/factory'

export const authN = createMiddleware<AuthNEnv>(async (c, next) => {
  const token = await verifyJWTCookie({ ctx: c })
  if (token.isErr()) {
    return c.json({}, 401)
  }
  c.set('activeUser', {
    userId: token.value.sub as string,
  })

  await next()
})
