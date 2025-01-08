import { getHC } from '@frontend/lib'
import { type QueryClient, useMutation } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'
import type { PostTicketRequest } from './types'

const useCreateTicket = (queryClient: QueryClient) => {
  const bff = getHC()
  return useMutation({
    mutationFn: (req: PostTicketRequest) => {
      return bff.tickets.$post({
        json: req,
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ticketQueryKeys.all,
      })
    },
  })
}

export { useCreateTicket }
