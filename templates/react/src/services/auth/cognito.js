import {
  CognitoUserPool,
  CognitoUserAttribute,
  CognitoUser,
  AuthenticationDetails,
} from 'amazon-cognito-identity-js'
// ref: https://github.com/aws-amplify/amplify-js/tree/master/packages/amazon-cognito-identity-js

const poolData = {
  UserPoolId: process.env.REACT_APP_COGNITO_POOL_ID,
  ClientId: process.env.REACT_APP_COGNITO_CLIENT_ID,
}
const userPool = new CognitoUserPool(poolData)

export function login({ email, password }) {
  const authDetails = new AuthenticationDetails({
    Username: email,
    Password: password,
  })
  const user = new CognitoUser({
    Username: email,
    Pool: userPool,
  })
  return new Promise((resolve, reject) => {
    user.authenticateUser(authDetails, {
      onSuccess: function(session) {
        var accessToken = session.getAccessToken().getJwtToken()
        resolve(accessToken)
      },
      onFailure: function(err) {
        reject(err)
      },
    })
  })
}

export function signup({ email, password }) {
  const dataEmail = {
    Name: 'email',
    Value: email,
  }
  const attributes = [new CognitoUserAttribute(dataEmail)]
  return new Promise((resolve, reject) => {
    userPool.signUp(email, password, attributes, null, (err, result) => {
      if (err) {
        reject(err)
      } else {
        resolve(result.user)
      }
    })
  })
}

export function forgotPassword({ username }) {
  const userData = {
    Username: username,
    Pool: userPool,
  }

  const cognitoUser = new CognitoUser(userData)
  return new Promise((resolve, reject) => {
    cognitoUser.forgotPassword({
      onSuccess: function(data) {
        resolve(cognitoUser)
      },
      onFailure: function(err) {
        reject(err)
      },
    })
  })
}

export function getToken() {
  const cognitoUser = userPool.getCurrentUser()
  if (cognitoUser != null) {
    return new Promise((resolve, reject) => {
      // NOTE: this get session function will automatically refresh access token
      cognitoUser.getSession(function(err, session) {
        if (err) {
          reject(err)
        } else {
          // note: both access and id tokens can be used to authenticate
          // https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html
          // ex. session.getAccessToken().getJwtToken()
          resolve({
            jwt: session.getIdToken().getJwtToken(),
            payload: session.getIdToken().payload,
          })
        }
      })
    })
  } else {
    throw new Error('Session not available')
  }
}

export function logout() {
  const cognitoUser = userPool.getCurrentUser()
  if (cognitoUser) {
    cognitoUser.signOut()
  }
}

export default { getToken, logout, login, signup }
