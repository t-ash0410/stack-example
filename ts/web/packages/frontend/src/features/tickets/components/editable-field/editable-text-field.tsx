import { Input } from '@frontend/components/ui/input'
import { Textarea } from '@frontend/components/ui/textarea'
import { useEditableField } from './use-editable-field'

type Props = {
  value: string
  onSave: (value: string) => void
  inputType?: 'input' | 'textarea'
}

export const EditableTextField = ({
  value,
  onSave,
  inputType = 'input',
}: Props) => {
  const { isEditing, setIsEditing, editedValue, setEditedValue } =
    useEditableField(value)

  const handleSave = () => {
    onSave(editedValue)
    setIsEditing(false)
  }

  if (isEditing) {
    const C = inputType === 'input' ? Input : Textarea
    return (
      <C
        value={editedValue}
        onChange={(e) => setEditedValue(e.target.value)}
        onBlur={handleSave}
        autoFocus
      />
    )
  }
  return (
    <span onClick={() => setIsEditing(true)} onKeyUp={() => setIsEditing(true)}>
      {value}
    </span>
  )
}
