import type { ReactNode } from 'react'
import { CommonLayout } from '../common-layout'
import { Header } from './header'
import { useCheckSession } from './use-check-session'

type Props = {
  title: string
  children: ReactNode
}

const AuthZLayout = ({ title, children }: Props) => {
  useCheckSession()

  return (
    <CommonLayout title={title}>
      <Header />
      {children}
    </CommonLayout>
  )
}

export { AuthZLayout }
