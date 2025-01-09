import type { ReactNode } from 'react'
import { CommonLayout } from '../common-layout'
import { Header } from './header'

type Props = {
  title: string
  children: ReactNode
}

const AuthZLayout = ({ title, children }: Props) => {
  return (
    <CommonLayout title={title}>
      <Header />
      {children}
    </CommonLayout>
  )
}

export { AuthZLayout }
