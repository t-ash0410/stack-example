import { DOMAIN } from '@bff/env'
import type { DefaultEnv } from '@bff/types'
import type { Context } from 'hono'
import { setCookie } from 'hono/cookie'
import { createFactory } from 'hono/factory'
import type { CookieOptions } from 'hono/utils/cookie'

const factory = createFactory<DefaultEnv>()

const handlers = factory.createHandlers(async (c: Context<DefaultEnv>) => {
  const expires = new Date()
  expires.setMinutes(expires.getMinutes() + 10)

  const opts: CookieOptions = {
    path: '/',
    secure: true,
    domain: DOMAIN,
    httpOnly: true,
    expires,
    sameSite: 'Strict',
  }

  const state = crypto.randomUUID()
  setCookie(c, 'state', state, opts)

  const nonce = crypto.randomUUID()
  setCookie(c, 'nonce', nonce, opts)

  return c.json({
    state,
    nonce,
  })
})

export { handlers as sessionHandlers }
