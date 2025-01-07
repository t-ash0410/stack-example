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
  mockTicketMgrServiceClient,
  mockTicketQuerierServiceClient,
} from '@bff/testutil'
import { create } from '@bufbuild/protobuf'
import { timestampFromDate } from '@bufbuild/protobuf/wkt'
import {
  DeleteTicketResponseSchema,
  GetTicketByIdResponseSchema,
  UpdateTicketResponseSchema,
} from '@stack-example/grpc'
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

describe('PUT /', async () => {
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
    spyOn(mockTicketMgrServiceClient, 'updateTicket').mockResolvedValue(
      create(UpdateTicketResponseSchema, {}),
    )
  })

  afterEach(() => {
    mock.restore()
  })

  it('returns 200 response', async () => {
    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title: 'Some Ticket',
        description: 'Some ticket description.',
        deadline: new Date('2020-01-01T00:00:00.000Z'),
      }),
    })

    expect(res.status).toBe(200)
    expect(await res.json()).toStrictEqual({})

    expect(mockTicketMgrServiceClient.updateTicket).toHaveBeenCalledTimes(1)
    expect(mockTicketMgrServiceClient.updateTicket).toHaveBeenCalledWith({
      ticketId: 'some-ticket',
      requestedBy: activeUser.userId,
      title: 'Some Ticket',
      description: 'Some ticket description.',
      deadline: timestampFromDate(new Date('2020-01-01T00:00:00.000Z')),
    })
  })

  it('returns 400 error if request validation fails', async () => {
    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        // important
      }),
    })

    expect(res.status).toBe(400)

    expect(mockTicketMgrServiceClient.updateTicket).not.toHaveBeenCalled()
  })

  it('returns 500 error if ticket update fails', async () => {
    spyOn(mockTicketMgrServiceClient, 'updateTicket').mockRejectedValue(
      new Error('Some error'), // important
    )

    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title: 'Some Ticket',
        description: 'Some ticket description.',
        deadline: new Date('2020-01-01T00:00:00.000Z'),
      }),
    })

    expect(res.status).toBe(500)

    expect(mockTicketMgrServiceClient.updateTicket).toHaveBeenCalledTimes(1)
  })
})

describe('DELETE /', async () => {
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
    spyOn(mockTicketMgrServiceClient, 'deleteTicket').mockResolvedValue(
      create(DeleteTicketResponseSchema, {}),
    )
  })

  afterEach(() => {
    mock.restore()
  })

  it('returns 200 response', async () => {
    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'DELETE',
    })

    expect(res.status).toBe(200)
    expect(await res.json()).toStrictEqual({})

    expect(mockTicketMgrServiceClient.deleteTicket).toHaveBeenCalledTimes(1)
    expect(mockTicketMgrServiceClient.deleteTicket).toHaveBeenNthCalledWith(1, {
      ticketId: 'some-ticket',
    })
  })

  it('returns 500 error if ticket deletion fails', async () => {
    const called = false
    spyOn(mockTicketMgrServiceClient, 'deleteTicket').mockRejectedValue(
      new Error('Some error'),
    )

    const app = initHonoApp().route(':ticketId', ticketDetailRoute)

    const res = await app.request('/some-ticket', {
      method: 'DELETE',
    })

    expect(res.status).toBe(500)

    expect(mockTicketMgrServiceClient.deleteTicket).toHaveBeenCalledTimes(1)
  })
})
