import { TicketList } from '@frontend/features/tickets'

export default () => {
  return (
    <main className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">TODOリスト管理</h1>
      <TicketList />
    </main>
  )
}
