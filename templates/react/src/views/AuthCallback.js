import { useEffect } from 'react'
import auth from '../services/auth'
import router from '../utils/router'

export default function AuthCallbackHandler() {
  useEffect(() => {
    auth
      .handleCallback()
      .then(resp => {
        router.push('/app')
      })
      .catch(err => {
        router.push('/')
      })
  })

  return null
}
