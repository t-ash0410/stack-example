import { localStorageKeys } from '@frontend/consts'
import { useGetSSOSlack } from '@frontend/hooks/api'
import { handleError } from '@frontend/util/handle-error'
import { useRouter, useSearchParams } from 'next/navigation'
import { useEffect } from 'react'

export function useSlackSSORedirect() {
  const router = useRouter()

  const params = useSearchParams()
  const code = params.get('code') || ''
  const state = params.get('state') || ''

  const { data, isLoading, isError } = useGetSSOSlack({ code, state })
  useEffect(() => {
    if (isLoading) return
    if (isError) {
      handleError(new Error('サインインに失敗しました'))
      router.push('/')
      return
    }

    const slackTeamId = data?.slackTeamId || ''
    localStorage.setItem(localStorageKeys.SLACK_TEAM_ID, slackTeamId)

    router.push('/')
  }, [router, data, isLoading, isError])
}
