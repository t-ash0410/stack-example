import { Button } from '@frontend/components'
import { formatDate } from '@frontend/util/date'
import { useState } from 'react'
import { DeleteButton } from '../delete-button'
import { EditableDateTimeField, EditableTextField } from '../editable-field'
import { TicketForm } from '../ticket-form'

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

export const TicketList = () => {
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

  const handleEditField = (
    ticket: Ticket,
    field: keyof Ticket,
    value: string,
  ) => {
    const updatedTicket = {
      ...ticket,
      [field]: value,
    }
    handleUpdateTicket(updatedTicket)
  }

  const handleEditDeadline = (ticket: Ticket, newDate: Date) => {
    const updatedTicket = {
      ...ticket,
      deadline: newDate,
    }
    handleUpdateTicket(updatedTicket)
  }

  return (
    <div className="space-y-4">
      <div className="mb-4">
        {isCreating ? (
          <TicketForm
            onSubmit={handleCreateTicket}
            onCancel={() => setIsCreating(false)}
          />
        ) : (
          <Button onClick={() => setIsCreating(true)}>新しいTODOを追加</Button>
        )}
      </div>
      {tickets.map((ticket) => (
        <div key={ticket.id} className="bg-white p-4 rounded-lg shadow">
          <div className="flex items-center justify-between mb-2">
            <h3 className="text-lg font-semibold w-full">
              <EditableTextField
                value={ticket.title}
                onSave={(value) => handleEditField(ticket, 'title', value)}
              />
            </h3>
            <div className="flex space-x-2">
              <DeleteButton onDelete={() => handleDeleteTicket(ticket.id)} />
            </div>
          </div>
          <div className="mb-2">
            <EditableTextField
              value={ticket.description}
              onSave={(value) => handleEditField(ticket, 'description', value)}
              inputType="textarea"
            />
          </div>
          <p className="text-sm text-gray-500">
            期限:&nbsp;
            <EditableDateTimeField
              value={ticket.deadline}
              onSave={(value) => handleEditDeadline(ticket, value)}
            />
          </p>
          <p
            className="pt-1 text-xs text-gray-400"
            suppressHydrationWarning={true}
          >
            作成日時:&nbsp;{formatDate(ticket.createdAt)}
          </p>
          <p className="text-xs text-gray-400" suppressHydrationWarning={true}>
            更新日時:&nbsp;{formatDate(ticket.updatedAt)}
          </p>
        </div>
      ))}
    </div>
  )
}
