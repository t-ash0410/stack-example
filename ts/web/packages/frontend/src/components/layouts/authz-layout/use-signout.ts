import { pagePaths } from '@frontend/consts'
import { getHC } from '@frontend/lib'
import { useRouter } from 'next/router'

const useSignout = () => {
  const router = useRouter()
  const handleSignout = async () => {
    const hc = getHC()
    await hc.session.$delete()
    router.push(pagePaths.signin.path)
  }
  return { handleSignout }
}

export { useSignout }
