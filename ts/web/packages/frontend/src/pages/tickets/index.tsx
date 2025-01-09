import { pagePaths } from '@frontend/consts'
import { TicketList } from '@frontend/features/tickets'
import { Layout } from '../_layout'

export default () => {
  return (
    <Layout title={pagePaths.tickets.title}>
      <main className="container mx-auto p-4">
        <h1 className="text-2xl font-bold mb-4">チケット一覧</h1>
        <TicketList />
      </main>
    </Layout>
  )
}
