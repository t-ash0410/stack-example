const { API_BASE_URL, SLACK_CLIENT_ID, NODE_ENV } = process.env

const apiBaseUrl = API_BASE_URL || 'http://0.0.0.0:8080'
const slackClientId = SLACK_CLIENT_ID || ''

export {
  NODE_ENV,
  apiBaseUrl as API_BASE_URL,
  slackClientId as SLACK_CLIENT_ID,
}
