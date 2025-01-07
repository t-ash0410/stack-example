import { describe, expect, it } from 'bun:test'
import { create } from '@bufbuild/protobuf'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { TicketSchema } from '@stack-example/grpc'
import { convertTicketToResponse } from './convert'

describe('convertTicketToResponse', () => {
  it('returns ticket response', async () => {
    const ticket = create(TicketSchema, {
      ticketId: 'ticket-001',
      createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
      updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
      createdBy: 'user-001',
      title: 'Some Ticket',
      description: 'Some ticket description.',
      deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
    })
    const res = convertTicketToResponse(ticket)
    expect(res).toStrictEqual({
      ticketId: 'ticket-001',
      createdAt: new Date('2020-01-01T00:00:00.000Z'),
      updatedAt: new Date('2020-01-01T00:00:00.000Z'),
      title: 'Some Ticket',
      description: 'Some ticket description.',
      deadline: new Date('2020-01-10T00:00:00.000Z'),
    })
  })
})
