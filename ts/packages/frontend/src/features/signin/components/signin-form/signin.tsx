import { SlackSSOButton } from '../slack-sso-button'

export function SigninForm() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="flex w-full max-w-md flex-col items-center gap-16">
        <SlackSSOButton />
      </div>
    </div>
  )
}
