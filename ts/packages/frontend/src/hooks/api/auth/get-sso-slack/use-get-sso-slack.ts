import { useBff } from '@frontend/hooks/use-bff'
import { handleError } from '@frontend/util/handle-error'
import { useQuery } from '@tanstack/react-query'
import { authQueryKeys } from '../query-key-factory'
import type { GetSSOSlackQueryParams } from './types'

export const useGetSSOSlack = (query: GetSSOSlackQueryParams) => {
  const bff = useBff()
  return useQuery({
    queryKey: [...authQueryKeys.ssoSlack(query)],
    queryFn: async () => {
      const res = await bff.auth.oidc.slack.$get({
        query,
      })
      if (!res.ok) handleError('システムエラー')
      return await res.json()
    },
  })
}