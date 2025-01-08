import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@frontend/components/ui/alert-dialog'
import { Button } from '@frontend/components/ui/button'
import { formatDate } from '@frontend/util/date'
import { Trash2 } from 'lucide-react'
import { useState } from 'react'
import { EditableDateTimeField, EditableTextField } from '../editable-field'

type Ticket = {
  id: string
  title: string
  description: string
  deadline: Date
  createdAt: Date
  updatedAt: Date
}

type TicketListProps = {
  tickets: Ticket[]
  onEdit: (updatedTicket: Ticket) => void
  onDelete: (id: string) => void
}

export function TicketList({ tickets, onEdit, onDelete }: TicketListProps) {
  const [ticketToDelete, setTicketToDelete] = useState<string | null>(null)

  const handleEditField = (
    ticket: Ticket,
    field: keyof Ticket,
    value: string,
  ) => {
    const updatedTicket = {
      ...ticket,
      [field]: value,
      updatedAt: new Date(),
    }
    onEdit(updatedTicket)
  }

  const handleEditDeadline = (ticket: Ticket, newDate: Date) => {
    const updatedTicket = {
      ...ticket,
      deadline: newDate,
      updatedAt: new Date(),
    }
    onEdit(updatedTicket)
  }

  return (
    <div className="space-y-4">
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
              <AlertDialog
                open={ticketToDelete === ticket.id}
                onOpenChange={(isOpen) => !isOpen && setTicketToDelete(null)}
              >
                <AlertDialogTrigger asChild>
                  <Button
                    variant="ghost"
                    size="icon"
                    aria-label="削除"
                    onClick={() => setTicketToDelete(ticket.id)}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>
                      本当にこのTODOを削除しますか？
                    </AlertDialogTitle>
                    <AlertDialogDescription>
                      この操作は取り消せません。
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>キャンセル</AlertDialogCancel>
                    <AlertDialogAction
                      onClick={() => {
                        if (ticketToDelete) {
                          onDelete(ticketToDelete)
                          setTicketToDelete(null)
                        }
                      }}
                    >
                      削除
                    </AlertDialogAction>
                  </AlertDialogFooter>
                </AlertDialogContent>
              </AlertDialog>
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
          <p className="text-xs text-gray-400" suppressHydrationWarning={true}>
            作成日時: {formatDate(ticket.createdAt)}
          </p>
          <p className="text-xs text-gray-400" suppressHydrationWarning={true}>
            更新日時: {formatDate(ticket.updatedAt)}
          </p>
        </div>
      ))}
    </div>
  )
}
