import React from 'react'
import Grid from '@material-ui/core/Grid'

import Centered from '../components/Centered'
import { Text } from '../components/Text'
import SampleRestQuery from '../components/graphql/SampleRestQuery'
import AppNavbar from '../components/AppNavbar'

function Dashboard({ location }) {
  return (
    <React.Fragment>
      <AppNavbar location={location} />
      <Centered container>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <Text>phrase.home_welcome</Text>:
            <SampleRestQuery />
          </Grid>
        </Grid>
      </Centered>
    </React.Fragment>
  )
}

export default Dashboard
