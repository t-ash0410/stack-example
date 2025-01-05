import { type PropsWithChildren, createContext, useContext } from 'react'

type Env = {
  NODE_ENV: string
  API_BASE_URL: string
  SLACK_CLIENT_ID: string
}

const envContext = createContext<Env>({
  NODE_ENV: '',
  API_BASE_URL: '',
  SLACK_CLIENT_ID: '',
})

export const useEnv = () => {
  return useContext(envContext)
}

export const EnvProvider: React.FC<PropsWithChildren<{ env: Env }>> = ({
  children,
  env,
}) => {
  return <envContext.Provider value={env}>{children}</envContext.Provider>
}
