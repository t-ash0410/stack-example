import { describe, expect, it } from 'bun:test'
import { JWT_KEY } from '@bff/env'
import { activeUser, initHonoApp } from '@bff/testutil'
import { sessionRoute } from './route'

describe('GET /', () => {
  it('returns user id', async () => {
    const app = initHonoApp().route('', sessionRoute)

    const res = await app.request('/')

    expect(res.status).toBe(200)
    expect(await res.json()).toEqual({
      userId: activeUser.userId,
    })
  })
})

describe('DELETE /', () => {
  it('returns 200 response', async () => {
    const app = initHonoApp().route('', sessionRoute)

    const res = await app.request('/', {
      method: 'DELETE',
    })

    expect(res.status).toBe(200)
    expect(res.headers.get('set-cookie')).toBe(
      `${JWT_KEY}=; Max-Age=0; Path=/; HttpOnly; Secure; SameSite=Strict`,
    )
  })
})
