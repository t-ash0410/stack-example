import { getHC } from '@frontend/lib'
import { type QueryClient, useMutation } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'

const useDeleteTicket = (queryClient: QueryClient) => {
  const bff = getHC()
  return useMutation({
    mutationFn: (ticketId: string) => {
      return bff.tickets[':ticketId'].$delete({
        param: {
          ticketId,
        },
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ticketQueryKeys.all,
      })
    },
  })
}

export { useDeleteTicket }
