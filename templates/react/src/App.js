import React from 'react'
import { Router, Route } from 'react-enroute'
import { ApolloProvider } from 'react-apollo'

import { router, apollo } from './utils'

import Dashboard from './views/Dashboard'
import Form from './views/Form'

import SessionLoader from './components/SessionLoader'
import LoadingIndicator from './components/LoadingIndicator'

export default function App(props) {
  if (navigator.userAgent === 'ReactSnap') {
    return <LoadingIndicator />
  }
  const { location } = props
  // TODO withParams for sessionLoader isn't necessary here, just there for demo
  return (
    <ApolloProvider client={apollo}>
      <Router location={location}>
        <Route path="/app" component={withParams(SessionLoader, { location })}>
          <Route path="" component={Dashboard} />
          <Route path="funds" component={Dashboard} />
          <Route path="inbox" component={Dashboard} />
          <Route path="forms/:id/step/:step" component={Form} />
          <Route path="(.*)" component={router.createRedirect('/app/')} />
        </Route>
      </Router>
    </ApolloProvider>
  )
}

function withParams(Component, params) {
  return props => <Component {...props} {...params} />
}
