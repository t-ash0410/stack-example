import { useBff } from '@frontend/hooks/use-bff'
import { useQuery } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'
import type { ListTicketResponse } from './types'

export const useInfiniteQueryJobs = () => {
  const bff = useBff()
  return useQuery<ListTicketResponse, Error>({
    queryKey: ticketQueryKeys.all,
    queryFn: async () => {
      const res = await bff.tickets.$get()
      return await res.json()
    },
  })
}
