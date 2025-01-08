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
import { Trash2 } from 'lucide-react'
import { useState } from 'react'

type Props = {
  onCancel: () => void
}

export const DeleteButton = ({ onCancel }: Props) => {
  const [open, setOpen] = useState<boolean>(false)
  const onClick = () => {
    setOpen(false)
    onCancel()
  }
  return (
    <AlertDialog open={open} onOpenChange={(o) => setOpen(o)}>
      <AlertDialogTrigger asChild>
        <Button
          variant="ghost"
          size="icon"
          aria-label="削除"
          onClick={() => setOpen(true)}
        >
          <Trash2 className="h-4 w-4" />
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>
            チケットを削除します。よろしいですか？
          </AlertDialogTitle>
          <AlertDialogDescription>
            この操作は取り消せません。
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>キャンセル</AlertDialogCancel>
          <AlertDialogAction onClick={onClick}>OK</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
