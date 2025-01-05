import { getLogger } from '@bff/log'
import type { Context, ErrorHandler } from 'hono'
import { HTTPException } from 'hono/http-exception'
import type { ContentfulStatusCode } from 'hono/utils/http-status'

export class ErrorWithResponseBody extends Error {
  body: Parameters<Context['json']>[0]
  status: ContentfulStatusCode

  constructor(
    message: string,
    body: Parameters<Context['json']>[0],
    status: ContentfulStatusCode = 500,
  ) {
    super(message)
    this.body = body
    this.status = status
  }
}

export const errorHandler: ErrorHandler = async (err, c) => {
  const logger = getLogger()

  const e = err as Error
  logger.error({
    message: `Server error occurred: ${c.req.url}`,
    error: {
      name: e.name,
      message: e.message,
      stack: e.stack,
    },
  })

  if (err instanceof ErrorWithResponseBody) {
    return c.json(err.body ?? {}, err.status)
  }
  if (err instanceof HTTPException) {
    return c.json({}, err.status)
  }
  return c.json({}, 500)
}
