import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'

const sessionFactory = createFactory<AuthNEnv>()

const sessionApp = sessionFactory.createApp()

export { sessionFactory, sessionApp }
