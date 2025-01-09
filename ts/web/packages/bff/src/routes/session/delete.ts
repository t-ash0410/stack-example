import { DOMAIN, JWT_KEY } from '@bff/env'
import type { AuthNEnv } from '@bff/types'
import type { Context } from 'hono'
import { deleteCookie } from 'hono/cookie'
import type { CookieOptions } from 'hono/utils/cookie'

const handler = async (c: Context<AuthNEnv>) => {
  const opts: CookieOptions = {
    path: '/',
    secure: true,
    domain: DOMAIN,
    httpOnly: true,
    sameSite: 'Strict',
  }
  deleteCookie(c, JWT_KEY, opts)

  return c.json({})
}

export { handler as deleteHandler }
