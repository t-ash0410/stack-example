import type { AuthNEnv } from '@bff/types'
import { createFactory } from 'hono/factory'

const handlers = createFactory<AuthNEnv>().createHandlers(async (c) => {
  const { userId } = c.var.activeUser
  return c.json({
    userId,
  })
})

export { handlers as getHandlers }
