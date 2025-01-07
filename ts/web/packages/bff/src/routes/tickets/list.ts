import type { AuthNEnv } from '@bff/types'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import type { QueryTicketsResponse } from '@stack-example/grpc'
import type { Context } from 'hono'
import { ResultAsync, ok } from 'neverthrow'

const handler = async (c: Context<AuthNEnv>) => {
  const res = await listTickets(c).andThen(convertResponse)
  if (res.isErr()) {
    throw res.error
  }
  return c.json(res.value)
}

const listTickets = (ctx: Context<AuthNEnv>) => {
  const { activeUser, ticketQuerierServiceClient } = ctx.var
  return ResultAsync.fromThrowable(() =>
    ticketQuerierServiceClient.queryTickets({
      requestedBy: activeUser.userId,
    }),
  )()
}

const convertResponse = (res: QueryTicketsResponse) => {
  return ok(
    res.tickets.map((t) => {
      return {
        ticketId: t.ticketId,
        createdAt: t.createdAt ? timestampDate(t.createdAt) : undefined,
        updatedAt: t.updatedAt ? timestampDate(t.updatedAt) : undefined,
        title: t.title,
        description: t.description,
        deadline: t.deadline ? timestampDate(t.deadline) : undefined,
      }
    }),
  )
}

export { handler as listHandler }
