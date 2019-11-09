import React from 'react'
import Button from '@material-ui/core/Button'

import Head from '../components/Head'
import PageContainer from '../components/PageContainer'
import Centered from '../components/Centered'
import Browser from '../components/Browser'
import { H2, H4, Text } from '../components/Text'
import { to } from '../utils/router'

const Home = () => (
  <PageContainer>
    <Head title="Welcome" />

    <Centered container>
      <div className="content">
        <H2>phrase.home_title</H2>
        <H4>phrase.home_welcome</H4>
        <Button
          data-testid="start-btn"
          variant="contained"
          color="primary"
          onClick={to('/app')}
        >
          <Text>start</Text>
        </Button>
      </div>
      <Browser
        src="https://place-hold.it/800x450/eee"
        width={800}
        height={450}
      />
    </Centered>
  </PageContainer>
)

export default Home
