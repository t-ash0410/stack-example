import { SLACK_CLIENT_ID } from '@frontend/config'
import { pagePaths } from '@frontend/consts'
import { useGetLocalStorage } from '@frontend/hooks'
import { getHC } from '@frontend/lib'
import { handleError } from '@frontend/util/handle-error'
import { useState } from 'react'

const SLACK_BASE_URL = 'https://slack.com/openid/connect/authorize'

export function useSlackSSOButton() {
  const teamId = useGetLocalStorage('SLACK_TEAM_ID')
  const bff = getHC()

  const [isPending, setIsPending] = useState<boolean>(false)

  async function onClick() {
    setIsPending(true)

    try {
      const res = await bff.auth.oidc.session.$get()
      if (!res.ok) {
        throw new Error('サインインに失敗しました')
      }
      const token = await res.json()

      const params = new URLSearchParams({
        scope: 'openid email profile',
        response_type: 'code',
        redirect_uri: `${location.origin}${pagePaths.sso.slack.path}`,
        client_id: SLACK_CLIENT_ID,
        team: teamId,
        state: token.state,
        nonce: token.nonce,
      })
      const url = `${SLACK_BASE_URL}?${params.toString()}`

      location.href = url
    } catch (err) {
      handleError(err)
      setIsPending(false)
    }
  }

  return { onClick, isPending }
}
