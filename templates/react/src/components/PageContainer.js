import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Button from '@material-ui/core/Button'

import { Text } from './Text'
import { Link } from '../utils/router'
import Footer from './Footer'

export default ({ children }) => (
  <React.Fragment>
    <AppBar position="static" color="default">
      <Toolbar>
        <Link href="/">
          <Button>
            <Text>home</Text>
          </Button>
        </Link>
        <Link href="/app">
          <Button>
            <Text>login</Text>
          </Button>
        </Link>
      </Toolbar>
    </AppBar>
    {children}
    <Footer />
  </React.Fragment>
)
