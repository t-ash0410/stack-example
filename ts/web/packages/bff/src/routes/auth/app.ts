import type { DefaultEnv } from '@bff/types'
import { createFactory } from 'hono/factory'

const authFactory = createFactory<DefaultEnv>()

const authApp = authFactory.createApp()

export { authFactory, authApp }
