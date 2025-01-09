import { AuthZLayout } from '@frontend/components'
import { pagePaths } from '@frontend/consts'
import { TicketList } from '@frontend/features/tickets'

export default () => {
  return (
    <AuthZLayout title={pagePaths.tickets.title}>
      <main className="container mx-auto p-4">
        <h1 className="text-2xl font-bold mb-4">チケット一覧</h1>
        <TicketList />
      </main>
    </AuthZLayout>
  )
}
