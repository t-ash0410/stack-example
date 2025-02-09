import type { DefaultEnv } from '@bff/types'
import { createFactory } from 'hono/factory'

const oidcFactory = createFactory<DefaultEnv>()

const oidcApp = oidcFactory.createApp()

export { oidcFactory, oidcApp }
