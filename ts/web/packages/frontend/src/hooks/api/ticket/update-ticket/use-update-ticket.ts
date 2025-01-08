import { getHC } from '@frontend/lib'
import { type QueryClient, useMutation } from '@tanstack/react-query'
import { ticketQueryKeys } from '../query-key-factory'
import type { PutTicketRequestJsonParams } from './types'

type Param = {
  ticketId: string
  body: PutTicketRequestJsonParams
}

const useUpdateTicket = (queryClient: QueryClient) => {
  const bff = getHC()
  return useMutation({
    mutationFn: ({ ticketId, body }: Param) => {
      return bff.tickets[':ticketId'].$put({
        param: {
          ticketId,
        },
        json: body,
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ticketQueryKeys.all,
      })
    },
  })
}

export { useUpdateTicket }
