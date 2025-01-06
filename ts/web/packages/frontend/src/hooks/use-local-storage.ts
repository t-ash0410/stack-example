import * as React from 'react'

export function useGetLocalStorage(key: string) {
  const [value, setValue] = React.useState<string>('')

  React.useEffect(() => {
    const v = localStorage.getItem(key)
    if (v) {
      setValue(v)
    }
  }, [key])

  return value
}

export function useSetLocalStorage(key: string, value: string) {
  React.useEffect(() => {
    localStorage.setItem(key, value)
  }, [key, value])
}
