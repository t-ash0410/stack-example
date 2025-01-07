import { zValidator } from '@hono/zod-validator'
import { z } from 'zod'

export const ticketDetailParamValidator = zValidator(
  'param',
  z.object({
    ticketId: z.string(),
  }),
)
