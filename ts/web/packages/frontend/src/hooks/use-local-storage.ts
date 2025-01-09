import { localStorageKeys } from '@frontend/consts'
import { useEffect, useState } from 'react'

export function useGetLocalStorage(key: keyof typeof localStorageKeys) {
  const [value, setValue] = useState<string>('')

  useEffect(() => {
    const v = localStorage.getItem(localStorageKeys[key])
    if (v) {
      setValue(v)
    }
  }, [key])

  return value
}

export function useSetLocalStorage(
  key: keyof typeof localStorageKeys,
  value: string,
) {
  useEffect(() => {
    localStorage.setItem(key, value)
  }, [key, value])
}
