import React from 'react'
import PropTypes from 'prop-types'
import { withApollo } from 'react-apollo'

import LoadingIndicator from './LoadingIndicator'
import auth from '../services/auth'
import { router } from '../utils'
import { queries } from '../constants'

class SessionLoader extends React.PureComponent {
  state = {
    loading: true,
    errorMsg: null,
  }

  componentDidMount() {
    this._auth()
  }

  render() {
    const { loading, errorMsg } = this.state

    if (loading !== false) {
      return <LoadingIndicator />
    } else if (errorMsg) {
      const Redirect = router.createRedirect('/login')
      return <Redirect />
    } else {
      return this.props.children || null
    }
  }

  _auth = async () => {
    const { client } = this.props
    try {
      this.setState({ loading: true })

      const { jwt } = await auth.getToken()

      if (!jwt) {
        // NOTE: if we don't logout, it'll cause a redirect loop
        await auth.logout()
        throw new Error('invalid token')
      }

      client.mutate({
        mutation: queries.LOGIN,
        variables: {
          accessToken: jwt,
        },
      })
      this.setState({ loading: false })
    } catch (err) {
      console.error(err)
      // TODO error handling separate exceptions from session expiry
      this.setState({ errorMsg: err.message })
      router.push('/login')
    }
  }
}

SessionLoader.propTypes = {
  client: PropTypes.object.isRequired,
}

export default withApollo(SessionLoader)
