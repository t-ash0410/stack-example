import { pagePaths } from '@frontend/consts'
import { useGetSession } from '@frontend/hooks/api'
import { handleError } from '@frontend/util/handle-error'
import { useRouter } from 'next/router'
import { useEffect } from 'react'

const useCheckSession = () => {
  const router = useRouter()

  const { data, isLoading, isError } = useGetSession()
  useEffect(() => {
    if (isLoading || !data) return
    if (isError) {
      handleError(new Error('エラーが発生しました'))
      return
    }
    if (data.status === 401) {
      router.push(pagePaths.signin.path)
      return
    }
  }, [router, data, isLoading, isError])
}

export { useCheckSession }
