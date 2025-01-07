import { Input } from '@frontend/components/ui/input'
import { Textarea } from '@frontend/components/ui/textarea'
import type React from 'react'
import { useState } from 'react'

type EditableFieldProps = {
  value: string
  onSave: (value: string) => void
  inputType?: 'input' | 'textarea'
}

export const EditableField: React.FC<EditableFieldProps> = ({
  value,
  onSave,
  inputType = 'input',
}) => {
  const [isEditing, setIsEditing] = useState(false)
  const [editedValue, setEditedValue] = useState(value)

  const handleSave = () => {
    onSave(editedValue)
    setIsEditing(false)
  }

  if (isEditing) {
    const InputComponent = inputType === 'input' ? Input : Textarea
    return (
      <InputComponent
        value={editedValue}
        onChange={(e) => setEditedValue(e.target.value)}
        onBlur={handleSave}
        autoFocus
      />
    )
  }

  return <span onClick={() => setIsEditing(true)}>{value}</span>
}
