import React from 'react'
import Tabs from '@material-ui/core/Tabs/Tabs'
import Tab from '@material-ui/core/Tab/Tab'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'

import { Text } from './Text'
import Navbar from './Navbar'
import { router } from '../utils'

const AppNavbar = ({ location, classes }) => {
  const tab = (location.split('/')[2] || '').toLowerCase()
  const _handleTab = (e, tab) => router.push(`/app/${tab}`)

  return (
    <Navbar className={classes.navbar}>
      <Tabs
        variant="fullWidth"
        indicatorColor="primary"
        value={tab}
        onChange={_handleTab}
      >
        <Tab value="home" label={<Text>home</Text>} />
      </Tabs>
    </Navbar>
  )
}

AppNavbar.propTypes = {
  location: PropTypes.string,
}

const styles = theme => ({
  navbar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
})

export default withStyles(styles)(AppNavbar)
