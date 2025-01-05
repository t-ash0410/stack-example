import { describe, expect, it } from 'bun:test'
import { Hono } from 'hono'
import { HTTPException } from 'hono/http-exception'
import { ErrorWithResponseBody, errorHandler } from './handler'

describe('errorHandler', () => {
  it('returns the error in the response when an ErrorWithResponseBody is thrown', async () => {
    // Setup Hono
    const app = new Hono().onError(errorHandler).get('/test', () => {
      throw new ErrorWithResponseBody(
        'some error',
        {
          foo: 'baa',
        },
        401,
      )
    })

    // Run
    const res = await app.request('/test')
    expect(res.status).toBe(401)
    expect(await res.json()).toStrictEqual({
      foo: 'baa',
    })
  })

  it('returns the error in the response when an HTTP exception is thrown', async () => {
    // Setup Hono
    const app = new Hono().onError(errorHandler).get('/test', () => {
      throw new HTTPException(400)
    })

    // Run
    const res = await app.request('/test')
    expect(res.status).toBe(400)
  })

  it('returns the 500 response when an unexpected exception is thrown', async () => {
    // Setup Hono
    const app = new Hono().onError(errorHandler).get('/test', () => {
      throw new Error('some error')
    })

    // Run
    const res = await app.request('/test')
    expect(res.status).toBe(500)
  })
})
