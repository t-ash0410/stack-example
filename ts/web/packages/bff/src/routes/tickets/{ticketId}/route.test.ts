import {
  afterEach,
  beforeEach,
  describe,
  expect,
  it,
  mock,
  spyOn,
} from 'bun:test'
import { initHonoApp, mockTicketQuerierServiceClient } from '@bff/testutil'
import { create } from '@bufbuild/protobuf'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import { GetTicketByIdResponseSchema } from '@stack-example/grpc'
import { ticketDetailRoute } from './route'

describe('GET /', async () => {
  beforeEach(() => {
    spyOn(mockTicketQuerierServiceClient, 'getTicketById').mockResolvedValue(
      create(GetTicketByIdResponseSchema, {
        ticket: {
          ticketId: 'ticket-001',
          createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
          updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
          createdBy: 'user-001',
          title: 'Some Ticket',
          description: 'Some ticket description.',
          deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
        },
      }),
    )
  })

  afterEach(() => {
    mock.restore()
  })

  it('returns tickets', async () => {
    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(200)
    expect(await res.json()).toMatchInlineSnapshot(`
{
  "createdAt": "2020-01-01T00:00:00.000Z",
  "deadline": "2020-01-10T00:00:00.000Z",
  "description": "Some ticket description.",
  "ticketId": "ticket-001",
  "title": "Some Ticket",
  "updatedAt": "2020-01-01T00:00:00.000Z",
}
`)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      2,
    )
    expect(
      mockTicketQuerierServiceClient.getTicketById,
    ).toHaveBeenNthCalledWith(2, {
      ticketId: 'some-ticket',
    })
  })

  it('returns 404 error if the ticket does not found', async () => {
    let called = false
    spyOn(mockTicketQuerierServiceClient, 'getTicketById').mockImplementation(
      async () => {
        if (called) {
          return create(GetTicketByIdResponseSchema, {}) // important
        }
        called = true

        return create(GetTicketByIdResponseSchema, {
          ticket: {
            ticketId: 'ticket-001',
            createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            createdBy: 'user-001',
            title: 'Some Ticket',
            description: 'Some ticket description.',
            deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
          },
        })
      },
    )

    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(404)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      2,
    )
  })

  it('returns 500 error if the ticket list fails to be retrieved', async () => {
    let called = false
    spyOn(mockTicketQuerierServiceClient, 'getTicketById').mockImplementation(
      async () => {
        if (called) {
          throw new Error('Some error') // important
        }
        called = true

        return create(GetTicketByIdResponseSchema, {
          ticket: {
            ticketId: 'ticket-001',
            createdAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            updatedAt: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
            createdBy: 'user-001',
            title: 'Some Ticket',
            description: 'Some ticket description.',
            deadline: timestampFromDate(new Date('2020-01-10T00:00:00.000Z')),
          },
        })
      },
    )

    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(500)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      2,
    )
  })
})
