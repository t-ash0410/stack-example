import { useEnv } from '@frontend/env'
import type { App } from '@stack-example/bff'
import { type ClientRequestOptions, hc } from 'hono/client'

export const useBff = (options?: ClientRequestOptions) => {
  const { API_BASE_URL } = useEnv()
  return hc<App>(API_BASE_URL, {
    init: {
      credentials: 'include',
    },
    ...options,
  })
}
