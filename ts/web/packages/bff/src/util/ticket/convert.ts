import { timestampDate } from '@bufbuild/protobuf/wkt'
import type { Ticket } from '@stack-example/grpc'

type TicketResponse = {
  ticketId: string
  createdAt: Date
  updatedAt: Date
  title: string
  description: string
  deadline?: Date
}

const convertTicketToResponse: (t: Ticket) => TicketResponse = (t: Ticket) => {
  return {
    ticketId: t.ticketId,
    createdAt: t.createdAt ? timestampDate(t.createdAt) : new Date(0),
    updatedAt: t.updatedAt ? timestampDate(t.updatedAt) : new Date(0),
    title: t.title,
    description: t.description,
    deadline: t.deadline ? timestampDate(t.deadline) : undefined,
  }
}

export { convertTicketToResponse }
