import { Button } from '@frontend/components/ui/button'
import { LogOut } from 'lucide-react'

const HamburgerMenu = () => {
  const handleLogout = () => {
    console.log('ログアウトが実行されました')
  }
  return (
    <nav className="py-2">
      <ul>
        <li>
          <Button
            variant="ghost"
            className="w-full justify-start"
            onClick={handleLogout}
          >
            <LogOut className="mr-2 h-4 w-4" />
            ログアウト
          </Button>
        </li>
      </ul>
    </nav>
  )
}

export { HamburgerMenu }
