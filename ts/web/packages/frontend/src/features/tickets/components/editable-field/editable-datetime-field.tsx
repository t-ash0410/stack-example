import { DateTimePicker } from '@frontend/components/ui/date-time-picker'
import { formatDate } from '@frontend/util/date'
import { useEditableField } from './use-editable-field'

type Props = {
  value: Date
  onSave: (value: Date) => void
}

export const EditableDateTimeField = ({ value, onSave }: Props) => {
  const { isEditing, setIsEditing, editedValue, setEditedValue } =
    useEditableField(value)

  const handleSave = () => {
    onSave(editedValue)
    setIsEditing(false)
  }

  if (isEditing) {
    return (
      <DateTimePicker
        date={new Date(value)}
        setDate={(newDate) => setEditedValue(newDate)}
        // onBlur={handleSave}
      />
    )
  }
  return (
    <span
      onClick={() => setIsEditing(true)}
      onKeyUp={() => setIsEditing(true)}
      suppressHydrationWarning={true}
    >
      {formatDate(editedValue)}
    </span>
  )
}
