import { useBff } from '@frontend/hooks/use-bff'
import { handleError } from '@frontend/util/handle-error'
import { useQuery } from '@tanstack/react-query'
import { sessionQueryKeys } from '../query-key-factory'

export const useGetSession = () => {
  const bff = useBff()
  return useQuery({
    queryKey: sessionQueryKeys.all,
    queryFn: async () => {
      const res = await bff.session.$get()
      if (!res.ok) {
        if (res.status === 401) {
          return {
            status: res.status,
          }
        }
        handleError(new Error(await res.text()))
        return
      }
      return {
        status: res.status,
        body: await res.json(),
      }
    },
  })
}
