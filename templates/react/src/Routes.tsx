import React, { useState, useEffect } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import { Router, Route } from 'react-enroute'
import queryString from 'qs'
import { SnackbarProvider } from 'notistack'

import appTheme from './utils/theme'
import AuthCallback from './views/AuthCallback'
import LoginView from './views/Login'
import SignupView from './views/Signup'
import PasswordResetView from './views/PasswordReset'
import Error404 from './views/Error404'
import LandingPage from './views/Landing'

import LoadingIndicator from './components/LoadingIndicator'
import history from './services/history'

const App = React.lazy(() => import('./App'))
const Admin = React.lazy(() => import('./Admin'))
const Terms = React.lazy(() => import('./views/Terms'))

export default function ConnectedRoutes() {
  const [state, setState] = useState({
    pathname: window.location.pathname,
    query: parseQuery(window.location.search),
  })

  useEffect(() => {
    // Returning the cleanup handler directly
    return history.listen(location => {
      const { pathname } = location
      setState({
        pathname,
        query: parseQuery(location.search),
      })
    })
  }, []) // disabling rerun on updates by passing a constant InputIdentityList

  return <Routes location={state.pathname} />
}

export function Routes({ location = '/' }) {
  return (
    <MuiThemeProvider theme={createMuiTheme(appTheme)}>
      <SnackbarProvider maxSnack={3}>
        <Router location={location}>
          <Route path="/" component={LandingPage} />
          <Route path="/auth/callback" component={AuthCallback} />
          <Route path="/login" component={LoginView} />
          <Route path="/signup" component={SignupView} />
          <Route path="/password-reset" component={PasswordResetView} />
          <Route path="/app(.*)" component={load(App, location)} />
          <Route path="/admin(.*)" component={load(Admin, location)} />
          <Route path="/terms-of-service" component={load(Terms, location)} />
          <Route path="(.*)" component={Error404} />
        </Router>
      </SnackbarProvider>
    </MuiThemeProvider>
  )
}

function parseQuery(query = '') {
  return queryString.parse(query, { ignoreQueryPrefix: true })
}

function load(Component: React.LazyExoticComponent<any>, location: string) {
  return (props: object) => (
    <React.Suspense fallback={<LoadingIndicator />}>
      <Component location={location} {...props} />
    </React.Suspense>
  )
}
