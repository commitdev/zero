import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import { withSnackbar } from 'notistack'

import { router, lang } from '../utils'
import auth from '../services/auth'
import Head from '../components/Head'
import Centered from '../components/Centered'
import { SimpleForm, Field, getFormData } from '../components/SimpleForm'
import { Text } from '../components/Text'

class Signup extends React.PureComponent {
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
        <Head title={lang.t('signup')} />
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
                <Button
                  data-testid="signup-btn"
                  variant="outlined"
                  type="submit"
                  disabled={this.state.loading}
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

  _submit = async ({ email, password }) => {
    try {
      this._isLoading()
      const { email, password } = getFormData(
        this.simpleForm.current.form.current
      )
      await auth.signup({ email, password })
      this.props.enqueueSnackbar('Signup successful', { variant: 'success' })
      auth.push('/login')
    } catch (err) {
      if (err.code === 'UsernameExistsException') {
        const errorMsg = err.message
        this.props.enqueueSnackbar(errorMsg, { variant: 'error' })
      } else if (err.code === 'InvalidParameterException') {
        const errorMsg = err.message
        this.props.enqueueSnackbar(errorMsg, { variant: 'error' })
      }
    } finally {
      this.setState({ loading: false })
    }
  }

  _isLoading = () => {
    this.setState({ loading: true })
  }
}

export default withSnackbar(Signup)
