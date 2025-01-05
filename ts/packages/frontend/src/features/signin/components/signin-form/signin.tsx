import Image from 'next/image'
import { SlackSSOButton } from '../slack-sso-button'

export function SigninForm() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="flex w-full max-w-md flex-col items-center gap-16">
        <Image
          src="/images/logo.jpg"
          alt="Stack Example"
          width={120}
          height={120}
        />
        <SlackSSOButton />
      </div>
    </div>
  )
}
