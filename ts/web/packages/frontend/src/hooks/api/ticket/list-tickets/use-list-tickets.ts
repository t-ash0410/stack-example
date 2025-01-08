import { useBff } from '@frontend/hooks/use-bff'
import { useQuery } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'
import type { ListTicketsResponse } from './types'

export const useListTickets = () => {
  const bff = useBff()
  return useQuery<ListTicketsResponse, Error>({
    queryKey: ticketQueryKeys.all,
    queryFn: async () => {
      const res = await bff.tickets.$get()
      return await res.json()
    },
  })
}
