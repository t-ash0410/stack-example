import { Button } from '@frontend/components'
import { useListTickets } from '@frontend/hooks/api/ticket'
import {} from '@frontend/types'
import { formatDate } from '@frontend/util/date'
import { handleError } from '@frontend/util/handle-error'
import { useState } from 'react'
import { DeleteButton } from '../delete-button'
import { EditableDateTimeField, EditableTextField } from '../editable-field'
import { TicketForm } from '../ticket-form'

export const TicketList = () => {
  const { data, isLoading, error, isError } = useListTickets()

  const [isCreating, setIsCreating] = useState(false)

  const handleCreateTicket = (newTicket: {
    title: string
    description: string
    deadline: Date
  }) => {
    setIsCreating(false)
  }

  const handleEditTitle = (id: string, value: string) => {}
  const handleEditDescription = (id: string, value: string) => {}
  const handleEditDeadline = (id: string, newDate: Date) => {}

  const handleDeleteTicket = (id: string) => {}

  if (isError) {
    handleError(error)
    return
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
      {isLoading ? <div>Loading...</div> : <></>}
      {data?.tickets.map((ticket) => (
        <div key={ticket.ticketId} className="bg-white p-4 rounded-lg shadow">
          <div className="flex items-center justify-between mb-2">
            <h3 className="text-lg font-semibold w-full">
              <EditableTextField
                value={ticket.title}
                onSave={(value) => handleEditTitle(ticket.ticketId, value)}
              />
            </h3>
            <div className="flex space-x-2">
              <DeleteButton
                onDelete={() => handleDeleteTicket(ticket.ticketId)}
              />
            </div>
          </div>
          <div className="mb-2">
            <EditableTextField
              value={ticket.description}
              onSave={(value) => handleEditDescription(ticket.ticketId, value)}
              inputType="textarea"
            />
          </div>
          <p className="text-sm text-gray-500">
            期限:&nbsp;
            <EditableDateTimeField
              value={ticket.deadline ? new Date(ticket.deadline) : undefined}
              onSave={(value) => handleEditDeadline(ticket.ticketId, value)}
            />
          </p>
          <p
            className="pt-1 text-xs text-gray-400"
            suppressHydrationWarning={true}
          >
            作成日時:&nbsp;{formatDate(new Date(ticket.createdAt))}
          </p>
          <p className="text-xs text-gray-400" suppressHydrationWarning={true}>
            更新日時:&nbsp;{formatDate(new Date(ticket.updatedAt))}
          </p>
        </div>
      ))}
    </div>
  )
}
