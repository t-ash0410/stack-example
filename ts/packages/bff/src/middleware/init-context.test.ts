import { describe, expect, it } from 'bun:test'
import type { DefaultEnv } from '@bff/types/hono'
import { Hono } from 'hono'
import { initContext } from './init-context'

describe('initContext', () => {
  it('sets the component in context', async () => {
    const app = new Hono<DefaultEnv>().use(initContext).get('/', (c) => {
      expect(c.var.logger).toBeTruthy()
      expect(c.var.accountMgrServiceClient).toBeTruthy()
      expect(c.var.ticketMgrServiceClient).toBeTruthy()
      expect(c.var.ticketQuerierServiceClient).toBeTruthy()
      return c.text('ok!')
    })
    await app.request('/')
  })
})
