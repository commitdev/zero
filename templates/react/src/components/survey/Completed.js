import React from 'react'
import SuccessIndicator from '../SuccessIndicator'
import Centered from '../../components/Centered'
import { H5, Text } from '../../components/Text'
import { Button, Box } from '@material-ui/core'
import { to } from '../../utils/router'

export default function Completed() {
  return (
    <Centered container>
      <SuccessIndicator />
      <Box p={2}>
        <H5>form completed</H5>
      </Box>
      <Button
        onClick={to('/app')}
        variant="contained"
        color="primary"
        size="large"
      >
        <Text>home</Text>
      </Button>
    </Centered>
  )
}
