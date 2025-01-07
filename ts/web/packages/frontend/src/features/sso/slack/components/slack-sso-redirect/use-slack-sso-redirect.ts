import { localStorageKeys } from '@frontend/consts'
import { useGetSSOSlack } from '@frontend/hooks/api'
import { handleError } from '@frontend/util/handle-error'
import { useRouter, useSearchParams } from 'next/navigation'
import { useEffect } from 'react'

export function useSlackSSORedirect() {
  const router = useRouter()

  const params = useSearchParams()
  const code = params.get('code')
  const state = params.get('state')
  const { data, isLoading, isError } = useGetSSOSlack(
    {
      code: code || '',
      state: state || '',
    },
    !!code && !!state,
  )
  useEffect(() => {
    if (isLoading || !data) return
    if (isError) {
      handleError(new Error('サインインに失敗しました'))
      router.push('/')
      return
    }

    localStorage.setItem(localStorageKeys.SLACK_TEAM_ID, data.slackTeamId)
    router.push('/')
  }, [router, data, isLoading, isError])
}
