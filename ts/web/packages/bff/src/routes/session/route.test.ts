import { describe, expect, it } from 'bun:test'
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
