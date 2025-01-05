import {
  API_BASE_URL,
  EnvProvider,
  NODE_ENV,
  SLACK_CLIENT_ID,
} from '@frontend/env'
import { Head, Html, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body className="antialiased">
        <EnvProvider
          env={{
            NODE_ENV: NODE_ENV,
            API_BASE_URL: API_BASE_URL,
            SLACK_CLIENT_ID: SLACK_CLIENT_ID,
          }}
        >
          <Main />
          <NextScript />
        </EnvProvider>
      </body>
    </Html>
  )
}
