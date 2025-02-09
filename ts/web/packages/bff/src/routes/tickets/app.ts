import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'

const ticketsFactory = createFactory<AuthNEnv>()

const ticketsApp = ticketsFactory.createApp()

export { ticketsFactory, ticketsApp }
