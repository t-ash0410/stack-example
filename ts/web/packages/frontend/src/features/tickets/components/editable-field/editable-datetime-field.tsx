import { DateTimePicker } from '@frontend/components/ui/date-time-picker'
import { formatDate } from '@frontend/util/date'
import { useState } from 'react'

type Props = {
  value: Date
  onSave: (value: Date) => void
}

export const EditableDateTimeField = ({ value, onSave }: Props) => {
  const [isEditing, setIsEditing] = useState(false)

  const handleSave = (newDate: Date) => {
    onSave(newDate)
    setIsEditing(false)
  }

  if (isEditing) {
    return <DateTimePicker date={new Date(value)} setDate={handleSave} />
  }
  return (
    <span
      onClick={() => setIsEditing(true)}
      onKeyUp={() => setIsEditing(true)}
      suppressHydrationWarning={true}
    >
      {formatDate(value)}
    </span>
  )
}
