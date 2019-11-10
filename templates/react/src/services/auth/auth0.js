import createAuth0Client from '@auth0/auth0-spa-js'

// ref: https://auth0.com/docs/libraries/auth0-spa-js
let auth0Client

async function initClient() {
  try {
    if (!auth0Client) {
      // WARNING: createAuth0Client seems to fail silently, doesn't throw error
      auth0Client = await createAuth0Client({
        domain: process.env.REACT_APP_AUTH_DOMAIN,
        client_id: process.env.REACT_APP_AUTH_CLIENT_ID,
        redirect_uri: process.env.REACT_APP_HOST + '/auth/callback',
        // onRedirectCallback: onRedirectCallback,
      })
    }

    return auth0Client
  } catch (e) {
    console.error(e)
    alert('Auth0 failed to initiate')
  }
}
initClient().then()

export async function login(params) {
  const auth0 = await initClient()

  await auth0.loginWithRedirect()
}

export async function handleCallback() {
  const auth0 = await initClient()
  return await auth0.handleRedirectCallback()
}

export async function getToken(params) {
  const auth0 = await initClient()
  const token = await auth0.getTokenSilently(params)
  return token
}

export async function getUser(params) {
  const auth0 = await initClient()
  const user = await auth0.getUser(params)
  return user
}

export async function logout() {
  const auth0 = await initClient()
  await auth0.logout({
    returnTo: process.env.REACT_APP_HOST,
  })
}

export default { getToken, logout, login, handleCallback, getUser }
