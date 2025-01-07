import { describe, expect, it } from 'bun:test'
import { initHonoApp } from '@bff/testutil'
import { ticketDetailParamValidator } from './param-validator'

describe('ticketDetailParamValidator', () => {
  it('succeeds validation', async () => {
    const app = initHonoApp().get(
      '/:ticketId',
      ticketDetailParamValidator,
      async (c) => {
        return c.json({})
      },
    )

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(200)
  })

  it('returns 400 error if the path does not contain an ticketId', async () => {
    const app = initHonoApp().get(
      '/:invalid', // important
      ticketDetailParamValidator,
      async (c) => {
        return c.json({})
      },
    )

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(400)
  })
})
