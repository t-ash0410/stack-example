import { describe, expect, it } from 'bun:test'
import { JWT_KEY } from '@bff/env'
import { initHonoApp } from '@bff/testutil'
import { createJWT, setJWTCookie } from './jwt'

describe('createJWT', async () => {
  it('returns jwt', async () => {
    const res = await createJWT({
      userId: 'user-001',
      now: new Date('2020-01-01T00:00:00.000Z'),
    })

    expect(res).toBe(
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLTAwMSIsImlhdCI6MTU3NzgzNjgwMCwiZXhwIjoxNTc3ODQ3NjAwfQ.1xemy5NAn5H2OlW9Eiuqipu3aIqW2XSKj6Kf3PnP4TU',
    )
  })
})

describe('setJWTCookie', async () => {
  it('sets jwt in the cookie', async () => {
    const app = initHonoApp().get('/', (c) => {
      setJWTCookie({
        ctx: c,
        jwt: 'jwt',
        now: new Date('2020-01-01T00:00:00.000Z'),
      })
      return c.text('ok!')
    })
    const res = await app.request('/')

    expect(res.headers.get('set-cookie')).toBe(
      `${JWT_KEY}=jwt; Path=/; Expires=Wed, 01 Jan 2020 03:00:00 GMT; HttpOnly; Secure; SameSite=Strict`,
    )
  })
})
