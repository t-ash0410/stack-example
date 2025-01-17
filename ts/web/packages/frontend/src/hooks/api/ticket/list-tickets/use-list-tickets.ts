import { getHC } from '@frontend/lib'
import { useQuery } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'

export const useListTickets = () => {
  const bff = getHC()
  return useQuery({
    queryKey: ticketQueryKeys.all,
    queryFn: async () => {
      const res = await bff.tickets.$get()
      if (!res.ok) {
        throw new Error(await res.text())
      }
      return await res.json()
    },
  })
}
