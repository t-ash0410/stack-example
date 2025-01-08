import type { BFF, BFFRequestJsonParams } from '@frontend/types'

type PostTicketApi = BFF['tickets']['$post']

type PostTicketRequest = BFFRequestJsonParams<PostTicketApi>

export type { PostTicketApi, PostTicketRequest }
