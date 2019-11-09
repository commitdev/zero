import React from 'react'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'

import Centered from '../components/Centered'
import { P, Text } from '../components/Text'
import SampleRestQuery from '../components/graphql/SampleRestQuery'
import AppNavbar from '../components/AppNavbar'
import { to } from '../utils/router'
import Query from '../components/graphql/Query'
import Queries from '../constants/queries'

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
          <Query query={Queries.GET_ALL_FORM_TEMPLATES}>{Forms}</Query>
        </Grid>
      </Centered>
    </React.Fragment>
  )
}

function Forms(data = {}) {
  const { formTemplates = [] } = data
  return (
    <Grid item xs={12}>
      {formTemplates.map(form => (
        <Card>
          <CardActions>
            <Button
              onClick={to(`/app/forms/${form.id}/step/0`)}
              size="small"
              color="primary"
            >
              <Text>{form.title}</Text>
            </Button>
          </CardActions>
        </Card>
      ))}
    </Grid>
  )
}

export default Dashboard
