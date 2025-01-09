import { pagePaths } from '@frontend/consts'
import { SigninForm } from '@frontend/features/signin'
import { Layout } from '../_layout'

export default () => {
  return (
    <Layout title={pagePaths.signin.title}>
      <SigninForm />
    </Layout>
  )
}
