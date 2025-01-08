import type { App } from '@stack-example/bff'
import type { ClientResponse, hc } from 'hono/client'

type BFF = ReturnType<typeof hc<App>>

// biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
type BFFResponse<T extends (...args: any[]) => Promise<ClientResponse<any>>> =
  Awaited<ReturnType<T>> extends ClientResponse<infer Out> ? Out : never

type BFFRequestArgs = {
  // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  query: any
  // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  param: any
  // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  json: any
}

type BFFRequestQueryParams<
  T extends (
    args: BFFRequestArgs,
    // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  ) => any,
> = Parameters<T>['0']['query']

type BFFRequestJsonParams<
  T extends (
    args: BFFRequestArgs,
    // biome-ignore lint/suspicious/noExplicitAny: Defined as a utility type, it should be widely available.
  ) => any,
> = Parameters<T>['0']['json']

export type { BFF, BFFRequestQueryParams, BFFRequestJsonParams, BFFResponse }
