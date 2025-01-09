import Head from 'next/head'
import type { ReactNode } from 'react'

const CommonLayout = (props: {
  title: string
  children: ReactNode
}) => {
  const t = `${props.title} | Stack Example`
  return (
    <>
      <Head>
        <title>{t}</title>
      </Head>
      {props.children}
      <footer>
        <div className="container mx-auto py-4 px-5 flex flex-wrap flex-col sm:flex-row">
          <p className="text-sm text-left">© 2025 Stack Example</p>
        </div>
      </footer>
    </>
  )
}

export { CommonLayout }
