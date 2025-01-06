import type { App } from '@stack-example/bff'
import type { ClientRequestOptions, ClientResponse, hc } from 'hono/client'

type BFF = ReturnType<typeof hc<App>>

// biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
type BFFResponse<T extends (...args: any[]) => Promise<ClientResponse<any>>> =
  Awaited<ReturnType<T>> extends ClientResponse<infer Out> ? Out : never

type BFFRequestQueryParams<
  T extends (
    args: {
      // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
      query: any
    },
    options?: ClientRequestOptions,
    // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  ) => any,
> = Parameters<T>['0']['query']

export type { BFF, BFFRequestQueryParams, BFFResponse }
