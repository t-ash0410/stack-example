import {
  DOMAIN,
  SLACK_CLIENT_ID,
  SLACK_CLIENT_SECRET,
  SLACK_SSO_REDIRECT_URL,
} from '@bff/env'
import { createJWT, setJWTCookie } from '@bff/jwt'
import type { AccountMgrServiceClient } from '@bff/types'
import { zValidator } from '@hono/zod-validator'
import type { Context } from 'hono'
import { deleteCookie, getCookie } from 'hono/cookie'
import type { CookieOptions } from 'hono/utils/cookie'
import { type JWTPayload, createRemoteJWKSet, jwtVerify } from 'jose'
import { Result, ResultAsync, err, ok } from 'neverthrow'
import { z } from 'zod'
import { oidcFactory } from './app'

const SLACK_SSO_URL = 'https://slack.com/api/openid.connect.token'
const SLACK_JWK_URL = 'https://slack.com/openid/connect/keys'
const SLACK_JWT_ISSUER = 'https://slack.com'

const validator = oidcFactory.createMiddleware(
  zValidator(
    'query',
    z.object({
      code: z.string(),
      state: z.string(),
    }),
  ),
)

const handlers = oidcFactory.createHandlers(validator, async (c) => {
  deleteSession(c)

  const accountMgr = c.var.accountMgrServiceClient
  const { state, code } = c.req.valid('query')
  const res = await checkState(c, state)
    .asyncAndThen(() => getSlackJWT(code))
    .andThen(convertTokenResponse)
    .andThen((res) => verifyAndDecode(c, res.idToken))
    .andThen((res) => callAccountMgr(accountMgr, res.payload))
  if (res.isErr()) {
    throw res.error
  }

  const now = new Date()
  const jwt = await createJWT({
    userId: res.value.userId,
    now,
  })
  setJWTCookie({
    ctx: c,
    jwt,
    now,
  })

  return c.json({
    jwt,
    slackTeamId: res.value.slackTeamId,
  })
})

const deleteSession = (c: Context) => {
  const opts: CookieOptions = {
    path: '/',
    secure: true,
    domain: DOMAIN,
    httpOnly: true,
    sameSite: 'Strict',
  }
  deleteCookie(c, 'state', opts)
  deleteCookie(c, 'nonce', opts)
}

const checkState = (c: Context, state: string) => {
  const cookieState = getCookie(c, 'state')
  if (!cookieState) {
    return err(new Error('No state set in cookie'))
  }

  if (cookieState !== state) {
    return err(
      new Error(`State do not match: cookie=${cookieState}, request=${state}`),
    )
  }

  return ok(c)
}

const getSlackJWT = (code: string) =>
  ResultAsync.fromThrowable(
    () => {
      return global.fetch(SLACK_SSO_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
          client_id: SLACK_CLIENT_ID,
          client_secret: SLACK_CLIENT_SECRET,
          code,
          redirect_uri: SLACK_SSO_REDIRECT_URL,
        }),
      })
    },
    (e) => new Error(`Error returned from Slack API: ${e}`),
  )()

const convertTokenResponse = (res: Response) =>
  ResultAsync.fromThrowable(() => res.json())()
    .andThen((j) =>
      Result.fromThrowable(() => {
        const schema = z.object({
          ok: z.boolean(),
          id_token: z.string().optional(),
          error: z.string().optional(),
        })

        return schema.parse(j)
      })(),
    )
    .andThen((res) =>
      res.ok
        ? ok({
            idToken: res.id_token || '',
          })
        : err(new Error(`Error returned from Slack API: ${res.error}`)),
    )
    .mapErr((e) => new Error(`Failed to parse token response: ${e}`))

const verifyAndDecode = (c: Context, idToken: string) =>
  ResultAsync.fromThrowable(() => {
    return jwtVerify(idToken, createRemoteJWKSet(new URL(SLACK_JWK_URL)), {
      issuer: SLACK_JWT_ISSUER,
      audience: SLACK_CLIENT_ID,
    })
  })().andThen((jwt) => {
    const cookieNonce = getCookie(c, 'nonce')
    if (!cookieNonce) {
      return err(new Error('No nonce set in cookie'))
    }
    if (cookieNonce !== jwt.payload.nonce) {
      return err(
        new Error(
          `Nonce do not match: cookie=${cookieNonce}, response=${jwt.payload.nonce}`,
        ),
      )
    }
    return ok(jwt)
  })

const callAccountMgr = (
  accountMgr: AccountMgrServiceClient,
  payload: JWTPayload,
) => {
  const teamId = payload['https://slack.com/team_id'] as string
  return ResultAsync.fromThrowable(() => {
    return accountMgr.slackSSO({
      email: payload.email as string,
      name: payload.name as string,
      slackUserId: payload.sub as string,
      slackTeamId: teamId,
    })
  })().andThen((res) => {
    return ok({
      userId: res.userId,
      slackTeamId: teamId,
    })
  })
}

export { handlers as slackHandlers }
