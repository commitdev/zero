import React from 'react'
import history from '../services/history'

export function push(uri) {
  history.push(uri)
}

export function to(uri) {
  return e => {
    e.preventDefault()
    history.push(uri)
  }
}

export const Link = (props = {}) => {
  return (
    <a {...props} onClick={to(props.to || props.href)}>
      {props.children}
    </a>
  )
}

/**
 * Redirect handling: https://github.com/tj/react-enroute/issues/4
 * - ex: <Route path="/unauthorized" component={replace('/login')} />
 */
export const createRedirect = redirectTo => {
  return function Redirect() {
    history.replace(redirectTo)
    return null
  }
}

export default {
  history,
  push,
  to,
  createRedirect,
  Link,
}
