import { useState } from 'react'

export const useEditableField = <S>(value: S) => {
  const [isEditing, setIsEditing] = useState(false)
  const [editedValue, setEditedValue] = useState(value)
  return {
    isEditing,
    setIsEditing,
    editedValue,
    setEditedValue,
  }
}
