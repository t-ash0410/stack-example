import getConfig from 'next/config'

const { publicRuntimeConfig } = getConfig()
const { API_BASE_URL, SLACK_CLIENT_ID } = publicRuntimeConfig

const apiBaseUrl = API_BASE_URL || 'http://localhost:8080'
const slackClientId = SLACK_CLIENT_ID || ''

export { apiBaseUrl as API_BASE_URL, slackClientId as SLACK_CLIENT_ID }
