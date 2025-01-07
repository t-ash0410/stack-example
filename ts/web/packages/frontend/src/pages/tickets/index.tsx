'use client'

import { Button } from '@frontend/components/ui/button'
import { TicketForm, TicketList } from '@frontend/features/tickets'
import { useState } from 'react'

type Ticket = {
  id: string
  title: string
  description: string
  deadline: Date
  createdAt: Date
  updatedAt: Date
}

const dummyTickets: Ticket[] = [
  {
    id: '1',
    title: 'プロジェクト計画書の作成',
    description: '来週の会議に向けて、新規プロジェクトの計画書を作成する',
    deadline: new Date('2023-06-30T17:00:00'),
    createdAt: new Date('2023-06-20T09:00:00'),
    updatedAt: new Date('2023-06-20T09:00:00'),
  },
  {
    id: '2',
    title: '週次レポートの提出',
    description: '先週の進捗状況をまとめて、上司に提出する',
    deadline: new Date('2023-06-23T12:00:00'),
    createdAt: new Date('2023-06-19T16:30:00'),
    updatedAt: new Date('2023-06-21T10:15:00'),
  },
  {
    id: '3',
    title: 'チームミーティングの準備',
    description: '明日のチームミーティングの議題と資料を準備する',
    deadline: new Date('2023-06-22T09:00:00'),
    createdAt: new Date('2023-06-21T14:00:00'),
    updatedAt: new Date('2023-06-21T14:00:00'),
  },
]

export default () => {
  const [tickets, setTickets] = useState<Ticket[]>(dummyTickets)
  const [isCreating, setIsCreating] = useState(false)

  const handleCreateTicket = (
    newTicket: Omit<Ticket, 'id' | 'createdAt' | 'updatedAt'>,
  ) => {
    const ticket: Ticket = {
      ...newTicket,
      id: Date.now().toString(),
      createdAt: new Date(),
      updatedAt: new Date(),
    }
    setTickets([...tickets, ticket])
    setIsCreating(false)
  }

  const handleUpdateTicket = (updatedTicket: Ticket) => {
    const updatedTickets = tickets.map((ticket) =>
      ticket.id === updatedTicket.id ? updatedTicket : ticket,
    )
    setTickets(updatedTickets)
  }

  const handleDeleteTicket = (id: string) => {
    setTickets(tickets.filter((ticket) => ticket.id !== id))
  }

  return (
    <main className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">TODOリスト管理</h1>
      {isCreating ? (
        <TicketForm
          onSubmit={handleCreateTicket}
          onCancel={() => setIsCreating(false)}
        />
      ) : (
        <Button onClick={() => setIsCreating(true)} className="mb-4">
          新しいTODOを追加
        </Button>
      )}
      <TicketList
        tickets={tickets}
        onEdit={handleUpdateTicket}
        onDelete={handleDeleteTicket}
      />
    </main>
  )
}
