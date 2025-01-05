import { cn } from '@frontend/lib/utils'
import { LoaderCircle } from 'lucide-react'

type Props = {
  size?: number
  className?: string
}

export function Spinner({ size = 40, className }: Props) {
  return (
    <LoaderCircle
      size={size}
      className={cn('mx-auto animate-spin', className)}
    />
  )
}
