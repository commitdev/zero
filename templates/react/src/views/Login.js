import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import Divider from '@material-ui/core/Divider'

import { router, lang } from '../utils'
import auth from '../services/auth'
import Head from '../components/Head'
import Centered from '../components/Centered'
import { SimpleForm, Field } from '../components/SimpleForm'
import { withSnackbar } from 'notistack'
import { Text, P } from '../components/Text'

class Login extends React.PureComponent {
  state = {
    loading: false,
  }

  constructor(props) {
    super(props)
    this.simpleForm = React.createRef()
  }

  render() {
    return (
      <React.Fragment>
        <Head title={lang.t('login')} />
        <AppBar position="static" color="default">
          <Button onClick={router.to('/')}>
            <Text>home</Text>
          </Button>
        </AppBar>
        <Centered container>
          <SimpleForm onSubmit={this._submit} ref={this.simpleForm}>
            <Grid container spacing={2} style={{ maxWidth: 380 }}>
              <Grid item xs={12}>
                <Field name="email" type="email" />
              </Grid>
              <Grid item xs={12}>
                <Field name="password" type="password" />
              </Grid>
              <Grid item xs={12}>
                <router.Link href="/password-reset">
                  <P>phrase.forgot_password</P>
                </router.Link>
              </Grid>
              <Grid item xs={12}>
                <Button
                  data-testid="login-btn"
                  variant="outlined"
                  type="submit"
                  disabled={this.state.loading}
                  fullWidth
                >
                  <Text>login</Text>
                </Button>
              </Grid>
              <Grid item xs={12}>
                <Divider />
              </Grid>
              <Grid item xs={12}>
                <Button
                  variant="outlined"
                  onClick={router.to('/signup')}
                  fullWidth
                >
                  <Text>signup</Text>
                </Button>
              </Grid>
            </Grid>
          </SimpleForm>
        </Centered>
      </React.Fragment>
    )
  }

  _closeNotification = () => {
    this.setState({ notifOpen: false })
  }

  _openNotification = (message, variant) => {
    this.setState({
      notifOpen: true,
      notifMessage: message,
      notifVariant: variant,
    })
  }

  _submit = async ({ email, password }) => {
    try {
      this._isLoading()
      await auth.login({email, password})
      window.location.assign('/app/')
    } catch (err) {
      const errorMsg = err.message
      this.props.enqueueSnackbar(errorMsg, { variant: 'error' })
    } finally {
      this.setState({ loading: false })
    }
  }

  _isLoading = () => {
    this.setState({ loading: true })
  }
}

export default withSnackbar(Login)
