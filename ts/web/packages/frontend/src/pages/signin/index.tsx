import { CommonLayout } from '@frontend/components'
import { pagePaths } from '@frontend/consts'
import { SigninForm } from '@frontend/features/signin'

export default () => {
  return (
    <CommonLayout title={pagePaths.signin.title}>
      <SigninForm />
    </CommonLayout>
  )
}
