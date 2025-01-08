import type { BFF, BFFResponse } from '@frontend/types'

type ListTicket = BFF['tickets']['$get']

type ListTicketResponse = BFFResponse<ListTicket>

export type { ListTicket, ListTicketResponse }
