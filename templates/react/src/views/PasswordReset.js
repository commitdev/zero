import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import { withSnackbar } from 'notistack'

import { router, lang } from '../utils'
import auth from '../services/auth'
import Head from '../components/Head'
import Centered from '../components/Centered'
import { SimpleForm, Field } from '../components/SimpleForm'
import { Text } from '../components/Text'

class PasswordReset extends React.PureComponent {
  state = {
    cognitoUser: null,
    errors: {},
    loading: false,
  }

  constructor(props) {
    super(props)
    this.simpleForm = React.createRef()
  }

  render() {
    const { cognitoUser } = this.state
    return (
      <React.Fragment>
        <Head title={lang.t('reset_password')} />
        <AppBar position="static" color="default">
          <Button onClick={router.to('/')}>
            <Text>home</Text>
          </Button>
        </AppBar>
        <Centered container>
          {cognitoUser
            ? this._renderVerificationForm()
            : this._renderForgotForm()}
        </Centered>
      </React.Fragment>
    )
  }

  _renderForgotForm() {
    const { errors } = this.state
    return (
      <SimpleForm onSubmit={this._submit} ref={this.simpleForm}>
        <Grid container spacing={2} style={{ maxWidth: 380 }}>
          <Grid item xs={12}>
            <Field
              name="email"
              type="email"
              error={Boolean(errors.email)}
              helperText={errors.email}
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              data-testid="reset-password-btn"
              variant="outlined"
              type="submit"
              disabled={this.state.loading}
              fullWidth
            >
              <Text>reset_password</Text>
            </Button>
          </Grid>
        </Grid>
      </SimpleForm>
    )
  }

  _renderVerificationForm() {
    const { errors } = this.state
    return (
      <SimpleForm onSubmit={this._resetPassword} ref={this.simpleForm}>
        <Grid container spacing={2} style={{ maxWidth: 380 }}>
          <Grid item xs={12}>
            <Field
              name="verificationCode"
              error={Boolean(errors.verificationCode)}
              helperText={errors.verificationCode}
            />
          </Grid>
          <Grid item xs={12}>
            <Field
              name="password"
              type="password"
              error={Boolean(errors.password)}
              helpText={errors.password}
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="outlined"
              type="submit"
              disabled={this.state.loading}
              fullWidth
            >
              <Text>reset_password</Text>
            </Button>
          </Grid>
        </Grid>
      </SimpleForm>
    )
  }

  _submit = async ({ email }) => {
    try {
      this._isLoading()
      const cognitoUser = await auth.forgotPassword({ username: email })
      this.setState({ loading: false, cognitoUser })
    } catch (err) {
      const errorMsg = err.message
      this.setState({ errors: { email: errorMsg } })
    } finally {
      this.setState({ loading: false })
    }
  }

  _resetPassword = ({ verificationCode, password }) => {
    this.state.cognitoUser.confirmPassword(verificationCode, password, {
      onSuccess() {
        this.props.enqueueSnackbar('Password successfully reset', {
          variant: 'success',
        })
        router.push('/login')
      },
      onFailure(err) {
        this.setState({ errors: { password: err.message } })
      },
    })
  }

  _isLoading = () => {
    this.setState({ errors: {}, loading: true })
  }
}

export default withSnackbar(PasswordReset)
