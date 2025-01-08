import type { BFF, BFFRequestJsonParams } from '@frontend/types'

type PutTicketApi = BFF['tickets'][':ticketId']['$put']

type PutTicketRequestJsonParams = BFFRequestJsonParams<PutTicketApi>

export type { PutTicketApi, PutTicketRequestJsonParams }
