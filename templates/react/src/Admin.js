import React from 'react'
import { Router, Route } from 'react-enroute'

import Forms from './views/admin/Forms'
import FormEdit from './views/admin/FormEdit'
import SessionLoader from './components/SessionLoader'
import LoadingIndicator from './components/LoadingIndicator'
import { ApolloProvider } from 'react-apollo'
import { router, apollo } from './utils'

export default function App({ location = '/' }) {
  if (navigator.userAgent === 'ReactSnap') {
    return <LoadingIndicator />
  }

  return (
    <ApolloProvider client={apollo}>
      <Router location={location}>
        <Route path="/admin" component={SessionLoader}>
          <Route path="forms" component={Forms} />
          <Route path="forms/:id" component={FormEdit} />
          <Route path="*" component={router.createRedirect('/admin/')} />
        </Route>
      </Router>
    </ApolloProvider>
  )
}
