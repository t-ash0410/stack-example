import type { BFF, BFFResponse } from '@frontend/types'

type ListTickets = BFF['tickets']['$get']

type ListTicketsResponse = BFFResponse<ListTickets>

export type { ListTickets, ListTicketsResponse }
