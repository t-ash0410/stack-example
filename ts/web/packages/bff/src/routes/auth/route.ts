import type { DefaultEnv } from '@bff/types'
import { createFactory } from 'hono/factory'
import { oidcRoute } from './oidc'

const authRoute = createFactory<DefaultEnv>()
  .createApp()
  .route('/oidc', oidcRoute)

export { authRoute }
