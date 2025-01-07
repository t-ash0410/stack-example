import { useGetSession } from '@frontend/hooks/api'
import { handleError } from '@frontend/util/handle-error'
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'

export function useRootRedirect() {
  const router = useRouter()

  const { data, isLoading, isError } = useGetSession()
  useEffect(() => {
    if (isLoading || !data) return
    if (isError) {
      handleError(new Error('エラーが発生しました'))
      return
    }
    if (data.status === 401) {
      router.push('/signin')
      return
    }
    router.push('/tickets')
  }, [router, data, isLoading, isError])
}
