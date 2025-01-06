const {
  BFF_PORT,
  CORS_ORIGIN,
  DOMAIN,
  JWT_SECRET,
  JWT_KEY,

  // Slack
  SLACK_CLIENT_ID,
  SLACK_CLIENT_SECRET,
  SLACK_SSO_REDIRECT_URL,

  // gRPC
  ACCOUNT_MGR_SERVICE_ENDPOINT,
  TICKET_MGR_SERVICE_ENDPOINT,
  TICKET_QUERIER_SERVICE_ENDPOINT,

  // develop
  USE_HTTPS,
  TLS_CERT_PATH,
  TLS_KEY_PATH,
} = process.env

const portAsNumber = BFF_PORT ? Number(BFF_PORT) : 8080
const corsOrigin = CORS_ORIGIN || 'https://localhost:3000'
const jwtKey = JWT_KEY || 'jwt'
const jwtSecret = JWT_SECRET || 'secret'

const slackClientId = SLACK_CLIENT_ID || ''
const slackClientSecret = SLACK_CLIENT_SECRET || ''
const slackSSORedirectUrl = SLACK_SSO_REDIRECT_URL || ''

const ticketMgrServiceEndpoint =
  TICKET_MGR_SERVICE_ENDPOINT || 'http://localhost:9080'

const ticketQuerierServiceEndpoint =
  TICKET_QUERIER_SERVICE_ENDPOINT || 'http://localhost:9081'

const accountMgrServiceEndpoint =
  ACCOUNT_MGR_SERVICE_ENDPOINT || 'http://localhost:9082'

const useHttps = USE_HTTPS === 'true'
const tlsCertPath = TLS_CERT_PATH || ''
const tlsKeyPath = TLS_KEY_PATH || ''

export {
  portAsNumber as BFF_PORT,
  corsOrigin as CORS_ORIGIN,
  DOMAIN,
  jwtSecret as JWT_SECRET,
  jwtKey as JWT_KEY,
  slackClientId as SLACK_CLIENT_ID,
  slackClientSecret as SLACK_CLIENT_SECRET,
  slackSSORedirectUrl as SLACK_SSO_REDIRECT_URL,
  accountMgrServiceEndpoint as ACCOUNT_MGR_SERVICE_ENDPOINT,
  ticketMgrServiceEndpoint as TICKET_MGR_SERVICE_ENDPOINT,
  ticketQuerierServiceEndpoint as TICKET_QUERIER_SERVICE_ENDPOINT,
  useHttps as USE_HTTPS,
  tlsCertPath as TLS_CERT_PATH,
  tlsKeyPath as TLS_KEY_PATH,
}
