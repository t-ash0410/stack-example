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
import { ticketDetailAuthZ } from './authz'
import { ticketDetailParamValidator } from './param-validator'

describe('ticketDetailAuthZ', () => {
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

  it('succeeds verification', async () => {
    const app = initHonoApp().get(
      '/:ticketId',
      ticketDetailParamValidator,
      ticketDetailAuthZ,
      async (c) => {
        return c.json({})
      },
    )

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(200)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      1,
    )
    expect(
      mockTicketQuerierServiceClient.getTicketById,
    ).toHaveBeenNthCalledWith(1, {
      ticketId: 'some-ticket',
    })
  })

  it('returns 403 error if the request is not made by the ticket creator him/herself', async () => {
    spyOn(mockTicketQuerierServiceClient, 'getTicketById').mockResolvedValue(
      create(GetTicketByIdResponseSchema, {
        ticket: {
          ticketId: 'ticket-001',
          createdBy: 'other-user', // important
        },
      }),
    )

    const app = initHonoApp().get(
      '/:ticketId',
      ticketDetailParamValidator,
      ticketDetailAuthZ,
      async (c) => {
        return c.json({})
      },
    )

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(403)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      1,
    )
  })

  it('returns 500 error if the ticket fails to be retrieved', async () => {
    spyOn(mockTicketQuerierServiceClient, 'getTicketById').mockRejectedValue(
      new Error('Some error'), // important
    )

    const app = initHonoApp().get(
      '/:ticketId',
      ticketDetailParamValidator,
      ticketDetailAuthZ,
      async (c) => {
        return c.json({})
      },
    )

    const res = await app.request('/some-ticket', {
      method: 'GET',
    })

    expect(res.status).toBe(500)

    expect(mockTicketQuerierServiceClient.getTicketById).toHaveBeenCalledTimes(
      1,
    )
  })
})
