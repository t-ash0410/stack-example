import { API_BASE_URL } from '@frontend/config'
import type { App } from '@stack-example/bff'
import { type ClientRequestOptions, hc } from 'hono/client'

export const getHC = (options?: ClientRequestOptions) => {
  return hc<App>(API_BASE_URL, {
    init: {
      credentials: 'include',
    },
    ...options,
  })
}
