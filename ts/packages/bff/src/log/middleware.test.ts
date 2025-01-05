import { describe, expect, it } from 'bun:test'
import { Hono } from 'hono'
import { informationLog } from './middleware'

describe('informationLog', () => {
  it('logs the request when it arrives', async () => {
    // Setup Hono
    const app = new Hono().use(informationLog).get('/test', (c) => {
      return c.text('ok!')
    })

    // Run
    const res = await app.request('/test')
    expect(res.status).toBe(200)
    expect(await res.text()).toBe('ok!')
  })
})
