import { getQueryClient } from '@frontend/lib'
import '@frontend/styles/globals.css'
import { QueryClientProvider } from '@tanstack/react-query'
import type { AppProps } from 'next/app'
import { Toaster } from 'sonner'

export default function App({ Component, pageProps }: AppProps) {
  const queryClient = getQueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <Toaster />
      <Component {...pageProps} />
    </QueryClientProvider>
  )
}
