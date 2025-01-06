import type { BFF, BFFRequestQueryParams } from '@frontend/types'

type GetSSOSlackApi = BFF['auth']['oidc']['slack']['$get']

type GetSSOSlackQueryParams = BFFRequestQueryParams<GetSSOSlackApi>

export type { GetSSOSlackApi, GetSSOSlackQueryParams }
