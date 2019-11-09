import React from 'react'
import Button from '@material-ui/core/Button'
import Typography from '@material-ui/core/Typography'

import { to } from '../utils/router'
import Centered from '../components/Centered'
import { H3, Text } from '../components/Text'

const View = () => {
  return (
    <Centered container>
      <Typography variant="h1">404</Typography>
      <H3>not_found</H3>
      <Button variant="outlined" onClick={to('/')}>
        <Text>home</Text>
      </Button>
    </Centered>
  )
}

export default View
