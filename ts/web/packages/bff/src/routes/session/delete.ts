import { DOMAIN, JWT_KEY } from '@bff/env'
import type { AuthNEnv } from '@bff/types'
import { deleteCookie } from 'hono/cookie'
import { createFactory } from 'hono/factory'
import type { CookieOptions } from 'hono/utils/cookie'

const handlers = createFactory<AuthNEnv>().createHandlers(async (c) => {
  const opts: CookieOptions = {
    path: '/',
    secure: true,
    domain: DOMAIN,
    httpOnly: true,
    sameSite: 'Strict',
  }
  deleteCookie(c, JWT_KEY, opts)

  return c.json({})
})

export { handlers as deleteHandlers }
