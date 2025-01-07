import {
  afterEach,
  beforeEach,
  describe,
  expect,
  it,
  mock,
  spyOn,
} from 'bun:test'
import {
  activeUser,
  initHonoApp,
  mockTicketQuerierServiceClient,
} from '@bff/testutil'
import { create } from '@bufbuild/protobuf'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { QueryTicketsResponseSchema } from '@stack-example/grpc'
import { ticketsRoute } from './route'

describe('GET /', async () => {
  beforeEach(() => {
    spyOn(mockTicketQuerierServiceClient, 'queryTickets').mockResolvedValue(
      create(QueryTicketsResponseSchema, {
        tickets: [
          {
            ticketId: 'ticket-001',
            createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            createdBy: 'user-001',
            title: 'Some Ticket',
            description: 'Some ticket description.',
            deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
          },
          {
            ticketId: 'ticket-002',
            createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            createdBy: 'user-001',
            title: 'Ticket 2',
            description: 'Foo baa',
            deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
          },
        ],
      }),
    )
  })

  afterEach(() => {
    mock.restore()
  })

  it('returns tickets', async () => {
    const app = initHonoApp().route('', ticketsRoute)

    const res = await app.request('/', {
      method: 'GET',
    })

    expect(res.status).toBe(200)
    expect(await res.json()).toMatchInlineSnapshot(`
[
  {
    "createdAt": "2020-01-01T00:00:00.000Z",
    "deadline": "2020-01-10T00:00:00.000Z",
    "description": "Some ticket description.",
    "ticketId": "ticket-001",
    "title": "Some Ticket",
    "updatedAt": "2020-01-01T00:00:00.000Z",
  },
  {
    "createdAt": "2020-01-01T00:00:00.000Z",
    "deadline": "2020-01-10T00:00:00.000Z",
    "description": "Foo baa",
    "ticketId": "ticket-002",
    "title": "Ticket 2",
    "updatedAt": "2020-01-01T00:00:00.000Z",
  },
]
`)

    expect(mockTicketQuerierServiceClient.queryTickets).toHaveBeenCalledTimes(1)
    expect(mockTicketQuerierServiceClient.queryTickets).toHaveBeenCalledWith({
      requestedBy: activeUser.userId,
    })
  })

  it('returns 500 error if the ticket list fails to be retrieved', async () => {
    spyOn(mockTicketQuerierServiceClient, 'queryTickets').mockRejectedValue(
      new Error('Some error'), // important
    )

    const app = initHonoApp().route('', ticketsRoute)

    const res = await app.request('/', {
      method: 'GET',
    })

    expect(res.status).toBe(500)

    expect(mockTicketQuerierServiceClient.queryTickets).toHaveBeenCalledTimes(1)
  })
})
