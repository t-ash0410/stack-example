import { Button } from '@frontend/components/ui/button'
import { Calendar } from '@frontend/components/ui/calendar'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@frontend/components/ui/popover'
import { cn } from '@frontend/lib/utils'
import { format } from 'date-fns'
import { ja } from 'date-fns/locale'
import { CalendarIcon } from 'lucide-react'

export function DateTimePicker({
  date,
  setDate,
}: { date: Date; setDate: (date: Date) => void }) {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={'outline'}
          className={cn(
            'w-full justify-start text-left font-normal',
            !date && 'text-muted-foreground',
          )}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          {date ? (
            format(date, 'PPP HH:mm', { locale: ja })
          ) : (
            <span>日時を選択</span>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0">
        <Calendar
          mode="single"
          selected={date}
          onSelect={(newDate) => newDate && setDate(newDate)}
          initialFocus
        />
        <div className="p-3 border-t border-border">
          <input
            type="time"
            value={format(date, 'HH:mm')}
            onChange={(e) => {
              const [hours, minutes] = e.target.value.split(':').map(Number)
              const newDate = new Date(date)
              newDate.setHours(hours)
              newDate.setMinutes(minutes)
              setDate(newDate)
            }}
            className="w-full p-2 border rounded"
          />
        </div>
      </PopoverContent>
    </Popover>
  )
}
