import { authApp } from './app'
import { oidcRoute } from './oidc'

const authRoute = authApp.route('/oidc', oidcRoute)

export { authRoute }
