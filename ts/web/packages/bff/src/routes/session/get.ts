import type { AuthNEnv } from '@bff/types'
import type { Context } from 'hono'

const handler = async (c: Context<AuthNEnv>) => {
  const { userId } = c.var.activeUser
  return c.json({
    userId,
  })
}

export { handler as getHandler }
